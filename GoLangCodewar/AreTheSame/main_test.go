// main.py
package main

import (
	"testing"
)

func TestComp(t *testing.T) {
	// Calculate Angle  for right  testcase
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}

	res := Comp(a, b)
	if res == false {
		t.Error("Result shoudl eb true")
	} else {
		t.Log("SUCCESS1:")

	}

}
