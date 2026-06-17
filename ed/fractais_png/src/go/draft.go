package main

import (
	"fmt"
	"math/rand"
)

func rand_int(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}


func main() {
	pen := NewPen(500, 500)
	pen.SetRGB(0, 0, 200)
	pen.SetPosition(250, 250)
	pen.SetHeading(0)
	pen.Walk(100)
	pen.Right(90)
	pen.SetRGB(200, 0, 0)
	pen.Walk(200)
	pen.Right(30)
	pen.Walk(250)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}