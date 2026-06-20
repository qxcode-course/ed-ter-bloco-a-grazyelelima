package main

import (
	"fmt"
	"math/rand"
)

func rand_int(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func pentatree(pen *Pen, dist float64) {
	if dist < 1 {
		return
	}
	angulo := 72.0
	fator := 0.3

	for range 5 {
		pen.SetLineWidth(0.5)
		pen.Right(angulo)

		pen.Walk(dist)

		pentatree(pen, dist*fator)
		pen.Walk(-dist)
	}
}

func main() {
	pen := NewPen(800, 800)
	pen.SetHeading(90)
	pen.SetPosition(0, 0)
	pen.FillSquare(800, 800)

	pen.SetPosition(400, 420)

	pen.SetRGB(255, 255, 255)
	pentatree(pen, 280)

	pen.SavePNG("pentatree.png")
	fmt.Println("PNG file created successfully.")
}