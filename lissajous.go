// lissajous_multi.go
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"
)

// Define a palette with multiple colors
var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // Red
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // Green
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // Blue
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // Yellow
	color.RGBA{0xff, 0x00, 0xff, 0xff}, // Magenta
	color.RGBA{0x00, 0xff, 0xff, 0xff}, // Cyan
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Handle HTTP requests and generate Lissajous GIFs
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Number of complete oscillations
		res     = 0.001 // Angular resolution
		size    = 100   // Image canvas covers [-size..+size]
		nframes = 64    // Number of animation frames
		delay   = 8     // Delay between frames (in 10ms units)
	)

	freq := rand.Float64() * 3.0 // Random frequency for y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Phase difference

	// Ensure a different color for each frame
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// Cycle through the colors in the palette (excluding the background)
		colorIndex := uint8((i % (len(palette) - 1)) + 1)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // Encode the animation and send to the output
}
