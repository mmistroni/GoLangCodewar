// Codewars1 project main.go
package main

import (
	"fmt"
	"math"
)

func calculate(d, vt float64) float64 {
	vtInCm := vt * 1000
	radius := d / 2
	rds := (math.Pi * math.Pow(radius, 2))
	return vtInCm / rds
}

// https://www.mathsisfun.com/geometry/cylinder-horizontal-volume.html

func calculateAngle(radius float64, height float64) (float64, float64) {
	fmt.Printf("radius:%v", radius)
	fmt.Printf("Height:%v", height)
	cos_theta := (radius - height) / radius
	// theta is half of the angle we need.
	res := math.Acos(cos_theta)
	fmt.Printf("Cos:%v, res:%v\n", cos_theta, res)
	return res, cos_theta
}

func sectorArea(theta float64, radius float64) float64 {
	return radius * radius * theta
}

func calculateTriangleArea(theta float64, radius float64) float64 {
	return 1 / 2 * radius * radius * math.Sin(2*theta) // Need  to multiply theta* 2
}

func calculateTriangleArea2(height float64, radius float64) float64 {
	return (radius - height) * math.Sqrt((2*radius*height)-(height*height))
}

func calculateTriangleArea3(height float64, diameter float64, volume float64) float64 {
	radius := diameter / 2.0
	tankWidth := calculate(diameter, volume)

	fmt.Printf("====Tank width:%v\n", tankWidth)
	tArea := calculateTriangleArea2(height, radius)

	sectorArea := math.Acos((radius-height)/radius) * radius * radius

	fluidLevel := (sectorArea - tArea) / 1000.0

	total := fluidLevel * tankWidth

	fmt.Printf("====Radius:%v, Height:%v, TArea:%v, Sarea:%v\n", radius, height, tArea, sectorArea)

	if height < radius {
		return total
	} else {
		halfCircle := volume / 2

		// so its halfCircle + (halfCircle = segment
		// cal works  here https://www.mathsisfun.com/geometry/cylinder-horizontal-volume.html
		// new sArea = 907.. chord = 3848/2 - 907 = 1017

		newHeight := diameter - height
		tArea2 := calculateTriangleArea2(newHeight, radius)
		sectorArea2 := math.Acos((radius-newHeight)/radius) * radius * radius

		fluidLevel2 := (sectorArea2 - tArea2) / 1000.0

		total2 := halfCircle - (fluidLevel2 * tankWidth)

		return halfCircle + total2

	}

}

// https://www.codewars.com/kata/55f3da49e83ca1ddae0000ad/train/go

func TankVol(h, d, vt int) float64 {
	height := float64(h)
	tankWidth := calculate(float64(d), float64(vt))
	theta, _ := calculateAngle(float64(d/2), height)
	fmt.Printf("Height:%v, radius:%v, angle:%v", height, d/2, theta)
	sarea := sectorArea(theta, float64(d/2))
	triangleArea := calculateTriangleArea(theta, float64(d/2))
	fluidLevel := (sarea - triangleArea) / 1000
	return fluidLevel * tankWidth
}

func main() {
	fmt.Println("Hello World!")
	res := TankVol(60, 120, 3500)
	fmt.Printf("RES1:%v\n", res)

	res2 := TankVol(6, 12, 3500)
	fmt.Printf("RES2:%v\n", res2)
}
