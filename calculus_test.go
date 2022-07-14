package main

import (
	"fmt"
	"math"
	"testing"
)

type Quadruple struct {
	height   float64
	diameter float64
	volume   float64
	result   float64
}

func TestCalculateTriangleArea(t *testing.T) {
	// Calculate Angle  for right  testcase
	radius := 3.5
	height := 5.0
	tankWidth := calculate(radius*2, 3848.0)
	theta, _ := calculateAngle(radius, height)
	sarea := sectorArea(theta, radius)
	triangleArea := calculateTriangleArea(theta, radius)
	fluidLevel := (sarea - triangleArea) / 1000
	t.Log("SUCCESS1:", fluidLevel*tankWidth)
}

func TestCalculateAnotherAngle(t *testing.T) {
	// Calculate Angle  for right  testcase
	diameter := 120.0
	radius := diameter / 2
	height := 60.0
	vt := 3500.0
	tankWidth := calculate(diameter, vt)
	tArea := calculateTriangleArea2(height, radius)

	sectorArea := math.Acos((radius-height)/height) * radius * radius

	fluidLevel := (sectorArea - tArea) / 1000
	t.Log("SUCCESS1:", fluidLevel*tankWidth)

}

func TestCalculateTriangleArea3(t *testing.T) {
	// Calculate Angle  for right  testcase
	//dotest(5, 7, 3848, 2940)
	//dotest(2, 7, 3848, 907)
	q1 := Quadruple{height: 3.0, diameter: 6.0, volume: 3500.0, result: 1750.0}
	q2 := Quadruple{height: 5.0, diameter: 7.0, volume: 3848.0, result: 2940.0}
	q3 := Quadruple{height: 2.0, diameter: 7.0, volume: 3848.0, result: 2940.0}
	holder := []Quadruple{q1, q2, q3}
	// https://www.mathsisfun.com/geometry/cylinder-horizontal-volume.html
	//https://www.cuemath.com/trigonometry/inverse-cosine/
	// https://math.stackexchange.com/questions/2806550/area-between-two-parallel-chords
	for _, tpl := range holder {
		fmt.Printf("----------- %v --------\n", tpl)
		result := calculateTriangleArea3(tpl.height, tpl.diameter, tpl.volume)
		t.Log("SUCCESS3:", result, tpl.result)
	}

	tester := math.Acos((3.5-2.0)/3.5) * 3.5 * 3.5
	t.Log("FINAL TESTER:", tester)

}
