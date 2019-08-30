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

// go语言中的复合声明
var palette = []color.Color{color.White, color.Black, color.RGBA{
	R: 6,
	G: 123,
	B: 33,
	A: 0,
}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func main() {
	// 当前时间伪随机
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

// 丽萨如图形
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	// 	两层for循环
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
