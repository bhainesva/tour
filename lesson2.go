package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strings"
	"tour/fractal"
)

type helloHandler struct {}
func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, " + r.FormValue("name")))
}

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s))
}

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.Greeting + s.Punct + s.Who))
}

type Image [][]uint8

func (i Image) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, i)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(i), len(i[0]))
}

func (i Image) At(x,y int) color.Color {
	val := i[y][x]
	return color.RGBA{val, 0,val, 255}
}


func Pic2(dx, dy int) Image {
	out := make([][]uint8, dy)
	for i:=0;i<dy;i++ {
		out[i] = make([]uint8, dx)

		for j:=0;j<dx;j++ {
			//out[i][j] = uint8((i + j) / 2)
			//out[i][j] = uint8(math.Pow(float64(i), float64(j)))
			out[i][j] = uint8(i * j)
		}
	}

	return Image(out)
}

type Colorer func (i, n int) color.Color

type FractalHandler struct {
	parser func(string, Colorer) image.Image
	colorer Colorer
}

func (h FractalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("p")

	colorer := fractal.Ramp
	if h.colorer != nil {
		colorer = h.colorer
	}

	img := h.parser(p, colorer)
	w.Header().Set("Content-Type", "image/png")

	png.Encode(w, img)
}

func in(c complex128, z complex128, limit int) int {
	for i:=0;i<limit;i++ {
		z = z * z + c
		if cmplx.Abs(z) > 2 {
			return i
		}
	}

	return limit
}

func mandlebrotParser(config string, colorer Colorer) image.Image {
	var (
		x, y int
		c0, dc complex128
		n int
	)

	fmt.Fscanf(strings.NewReader(config), "%d %d %g %g %d", &x, &y, &c0, &dc, &n)

	return FractalImage{
		x:  x,
		y:  y,
		at: func (x2, y2 int) color.Color {
			fx := float64(x2) / float64(x)
			fy := float64(y2) / float64(y)
			c := c0 + complex(real(dc)*fx, imag(dc) * fy)
			i := in(c, complex(0.0, 0), n)
			return colorer(i, n)
		},
	}
}

func juliaParser(config string, colorer Colorer) image.Image {
	var (
		x, y int
		z0, dz, c complex128
		n int
	)

	fmt.Fscanf(strings.NewReader(config), "%d %d %g %g %g %d", &x, &y, &z0, &dz, &c, &n)

	return FractalImage{
		x:  x,
		y:  y,
		at: func (x2, y2 int) color.Color {
			fx := float64(x2) / float64(x)
			fy := float64(y2) / float64(y)
			z := z0 + complex(real(dz)*fx, imag(dz) * fy)
			i := in(c, z, n)
			return colorer(i, n)
		},
	}
}

type FractalImage struct {
	x, y int
	at func(x, y int) color.Color
}

func (m FractalImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (m FractalImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.x, m.y)
}

func (m FractalImage) At(x,y int) color.Color {
	return m.at(x, y)
}


