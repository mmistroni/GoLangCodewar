// main.py
package main

// https://www.codewars.com/kata/550498447451fbbd7600041c/train/go
/**

Given two arrays a and b write a function comp(a, b)
(orcompSame(a, b))
that checks whether the two arrays have the "same" elements,
with the same multiplicities
(the multiplicity of a member is the number
of times it appears).
"Same" means, here, that the elements in b
are the elements in a squared, regardless of the order."
**/
import (
	"fmt"
)

func Comp(array1 []int, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	}
	// Not good. need to make a set of the input
	results := []bool{}
	for _, source := range array1 {
		for _, target := range array2 {
			if source*source == target {
				results = append(results, true)
				fmt.Printf("%v, %v \n", source, target)
			}
		}
	}
	if len(results) == len(array1) {
		return true
	}
	return false
}

func main() {
	fmt.Println("Hello World!")
	a := []int{1, 2}
	b := []int{3, 4}
	fmt.Println(Comp(a, b))
}
