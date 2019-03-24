// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	width, height = 900, 700            // canvas size in pixels
	cells         = 140                 // number of grid cells
	xyrange       = 25.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	//f, err := os.Create("SVG.xml")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()

	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/s", svg)

	log.Fatalln(http.ListenAndServe("localhost:8000", nil))

	//svg(f)
}

func svg(out http.ResponseWriter, r *http.Request /*, f *os.File*/) {
	out.Header().Set("Content-Type", "image/svg+xml")
	out.WriteHeader(200)
	_, err := out.Write([]byte(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:black;fill:white;fill-opacity:0.9;stroke-width: 0.5' "+
		"width='%d' height='%d'>\n", width, height)))
	if err != nil {
		panic(err)
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)

			dx, dy, col := corner(i+1, j+1)
			col = strings.Replace(fmt.Sprint("#0000ff"), " ", "", -1)

			_, err = out.Write([]byte(
				fmt.Sprintf(
					"<polygon style='fill:%s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					col, ax, ay, bx, by, cx, cy, dx, dy)))
			if err != nil {
				panic(err)
			}
		}
	}
	//if ; err!= nil{
	//	out.WriteString(buf)
	//}
	_, err = out.Write([]byte("</svg>"))
	if err != nil {
		panic(err)
	}
}

func corner(i, j int) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	var col = "#"
	//fmt.Print("_ ", z, x, y, " _")
	if x <= 0.0 {
		//col += strings.Replace(fmt.Sprint("ffffff"), " ", "", -1)
		col += "ff0000"
	} else {
		col += "0000ff"
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, col
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
