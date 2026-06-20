package main

import (
	"fmt"
	"math/rand"
)

func rand_int(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func circles(pen *Pen, dist float64) {
	if dist < 6 {
		return
	}
	angulo := 60.0
	fator := 0.339

	pen.DrawCircle(dist)
	for range 6 {
		pen.SetLineWidth(1.0)
		pen.Right(angulo)

		pen.Up()
		pen.Walk(dist)
		pen.Down()

		pen.DrawCircle(dist * fator)
		circles(pen, dist*fator)

		pen.Up()
		pen.Walk(-dist)
		pen.Down()
	}
}

func main() {
	pen := NewPen(800, 800)
	pen.SetHeading(80)
	pen.SetPosition(0, 0)
	pen.FillSquare(800, 800)

	pen.SetPosition(400, 420)
	pen.SetRGB(250, 250, 250)
	circles(pen, 250)

	pen.SavePNG("circles.png")
	fmt.Println("PNG file created successfully.")
}