package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"fmt"
)

var black = color.Black

const (
	graphColorIndex = 1 //следующий цвет палитры
)

func main() {

	// Генерируем freq
    freq := rand.Float64() * 3.0

    // Формируем имя файла с округленным значением freq
    filename := fmt.Sprintf("lissajous_%.3f.gif", freq)

    // Создаем файл
    f, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Передаем freq в функцию lissajous
    lissajous(f, freq)
}

func lissajous (out io.Writer, freq float64) {
	const (
		cycles = 5 // количество полных колебаний Х
		res = 0.001 // угловое разрешение
		size = 100 // Канва изображения охватывает [size ... +size]
		nframes = 64 // количество кадров анимации
		delay = 8 // задержка между кадрами (единица - 10 мс)
	)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i:=0; i<nframes; i++{
		rect:=image.Rect(0,0,2*size+1, 2*size+1)
		rValue := uint8(rand.Float64()* 255) 
		gValue := uint8(rand.Float64()* 255) 
		bValue := uint8(rand.Float64()* 255) 
		aValue := uint8(255)
		customColor := color.RGBA{
			R: rValue,
			G: gValue,
			B: bValue,
			A: aValue,
		}
		palette := []color.Color{black, customColor }
		img:= image.NewPaletted(rect, palette)
		for t:=0.0 ; t<cycles*2*math.Pi; t+= res {
			x:=math.Sin(t)
			y:=math.Sin(t*freq + phase)
			img.SetColorIndex(size +int(x*size+0.5), size + int(y*size +0.5), graphColorIndex)
		}
		
		
		anim.Image = append(anim.Image, img)
		anim.Delay = append(anim.Delay, delay)
		phase +=0.1
	}
	gif.EncodeAll(out, &anim)
}
