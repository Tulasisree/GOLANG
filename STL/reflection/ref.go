package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

//"listings/movies/media"

func main() {

	// fmt.Println(" my fav movie")
	// m1 := media.Movie{}
	// m1.SetTitle("TULASI")
	// m1.SetRating(media.NC17)
	// m1.SetBoxOffice(6.23)
	// fmt.Println(m1.GetBoxOffice(), m1.GetRating(), m1.GetTitle())

	/*type fruit struct {
		name  string
		color string
	}

	type vegetable struct {
		name  string
		color string
	}

	newfruit := fruit{"apple", "red"}
	newveg := vegetable{"potato", "brown"}

	addvf(newfruit)
	addvf(newveg)*/

	timed := maketimedfunc(sayhello).(func(string) string) //singnature
	myname := timed("Tulasi")
	fmt.Println(myname)

}

//making func
func maketimedfunc(f interface{}) interface{} {
	rf := reflect.TypeOf(f)

	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}

	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in) //calling the original function we passed in
		end := time.Now()
		fmt.Printf("calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func sayhello(name string) string {
	fmt.Println("hello", name)
	return "say hello called"
}

func addvf(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		switch reflect.TypeOf(p).Name() {
		case "fruit":
			fmt.Println("Added fruit")
		case "vegetable":
			fmt.Println("Added vegetable")
		}
		return true
	} else {
		return false
	}
}
