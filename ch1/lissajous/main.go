// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette variable
	blackIndex = 1 // second color in palette variable
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
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
