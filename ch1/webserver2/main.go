// Webserver to count requests and serve Lissajous figures
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Mutex to count requests correctly
var mu sync.Mutex

// The number of requests received by the server
var requestCount int

// The list of all colors that will be in the image
var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette variable
	blackIndex = 1 // second color in palette variable
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	lissajousHandler := func(w http.ResponseWriter, r *http.Request) {
		Lissajous(w)
	}
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	requestCount++
	mu.Unlock()

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	localRequestCount := requestCount
	mu.Unlock()
	fmt.Fprintf(w, "Request count = %d\n", localRequestCount)
}

// This method generates an animated gif of a Lissajous figure and writes it out to STDOUT
func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // x goes from 0 to 2pi 5 times
		res     = 0.001 // angular resolution
		size    = 100   // [-size..+size] is the canvas
		nframes = 64    // 64 frames in gif
		delay   = 8     // units of 10 ms
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		limit := cycles * 2 * math.Pi
		for t := 0.0; t < limit; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
