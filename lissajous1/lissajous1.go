package main

import (
	"math"
	"math/rand"
	"os"
	"fmt"
	"image/gif"
	"image/color"
	"image"
	"io"
)

const (
	cycles = 5
	size=100
	rez = 0.001
	delay = 8
	nframes = 64
	backgroundColorID = 0
	graphColorID = 1
)

func main() {
	
	freq:=rand.Float64() * 3

	fileName := fmt.Sprintf("./lissajous1/llissajous_%.3f.gif", freq)

	f, err := os.Create(fileName)
	if err!=nil {
		panic(err)
	}
	defer f.Close()


	lissajous(f,freq)

}

func randomRGBAvalue() uint8 {
	return uint8(rand.Int31n(256))
}

func randomRGBAcolor() color.Color {
	return color.RGBA{
		R: randomRGBAvalue(),
		G: randomRGBAvalue(),
		B: randomRGBAvalue(),
		A: 255,
	}
}

func lissajous (out io.Writer, freq float64) {
	anim := gif.GIF{LoopCount:nframes}
	rect := image.Rect(0,0,2*size+1,2*size+1)
	phase := 0.0
	for frameNumber:=0; frameNumber<nframes; frameNumber++{
		palette := []color.Color{color.Black, randomRGBAcolor()}
		frame := image.NewPaletted(rect, palette)
		for angle:=0.0; angle<=2*math.Pi*cycles; angle+=rez {
			x:=math.Sin(angle)
			y:=math.Sin(angle*freq + phase)
			frame.SetColorIndex(size + int(size*x+0.5),size + int(size*y+0.5),graphColorID )
		}
		phase+=0.1
		anim.Delay = append(anim.Delay,delay)
		anim.Image = append(anim.Image,frame)
	}
	
	gif.EncodeAll(out, &anim)
}