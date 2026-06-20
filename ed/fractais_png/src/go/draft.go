package main

import (
	"fmt"
	"math/rand"
)

func rand_int(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func embua(pen *Pen, dist float64) {
	if dist < 1{
		return
	}

	pen.SetLineWidth(dist / 30)
	pen.Walk(dist)
	pen.Right(90)
	dist *= 0.97
	embua(pen, dist)
}

func main() {
	pen := NewPen(500, 500)
	pen.SetRGB(0, 0, 200)
	pen.SetHeading(0)
	pen.SetPosition(100, 100)
	embua(pen, 300)
	pen.SavePNG("embua.png")
	fmt.Println("PNG file created successfully.")
}