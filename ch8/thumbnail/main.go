package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func thumbnail(image string) (string, error) {
	time.Sleep(5 * time.Millisecond)
	return "", nil
}

func createThumbnails(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			if fnew, err := thumbnail(f); err != nil {
				log.Printf("error creating thumbnail: %s", err)
			} else {
				log.Printf("thumbnail created for %s as %s\n", f, fnew)
			}
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

func thumbnail2(image string) (string, error) {
	time.Sleep(5 * time.Millisecond)
	if rand.Intn(700) > 600 {
		return "", fmt.Errorf("thumbnail creation failed for %s", image)
	}
	return "", nil
}

func createThumbnails2(filenames []string) error {
	errors := make(chan error, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			if fnew, err := thumbnail2(f); err != nil {
				log.Printf("error creating thumbnail: %s", err)
				errors <- err
			} else {
				log.Printf("thumbnail created for %s as %s\n", f, fnew)
				errors <- nil
			}
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

func createThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			if _, err := thumbnail2(f); err != nil {
				log.Printf("error creating thumbnail: %s", err)
				return
			}
			sizes <- 100
		}(f)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for s := range sizes {
		total += s
	}
	return total
}

func main() {
	var files []string
	root := "/Users/bharath/Pictures/chromecast1200"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatalf("error listing folder: %s", err)
	}
	log.Printf("NUMBER OF FILES: %d\n", len(files))
	//err = createThumbnails2(files)
	//log.Println(err)
	fmt.Println("SIZE:", createThumbnails6(files))
}
