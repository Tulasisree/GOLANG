package main

import (
	"fmt"
	"github/gofuncs/simplemath"
	"net/http"
	"strings"
)

func main() {
	//ignoring results from function
	//ans, _ := divide(2, 3)
	ans, err := simplemath.Divide(2, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ans)
	}
	// fmt.Printf("%f\n", simplemath.Add(1, 2))

	fmt.Printf("%f\n", simplemath.Sum(1.1, 2.2, 3.3, 4.4))

	numbers := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	//this will split numbers
	//taking a slice and spreading them out
	fmt.Printf("%f\n", simplemath.Sum(numbers...))

	sv := simplemath.NewSemanticVersion(1, 2, 3)
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	fmt.Printf("%s", sv.String())

	//for interfaces you need to have &
	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}
