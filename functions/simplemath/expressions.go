package simplemath

import (
	"errors"
)

//Pulblic and private
//Capitalized name- Public

//returning error
//naming return values
//
func Divide(p1, p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}
	answer = p1 / p2
	return
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

//variadic functions
//it should be the final parameter
//get number a values and makes it into slice
func Sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}
