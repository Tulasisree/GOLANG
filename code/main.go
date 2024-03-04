//module -> entire project space while source code
//package - within module we have packages discrete units of code related to some concept within module
package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

/*const (
	first  = iota
	second = iota
	third  = iota + 6
	fourth = 2 << iota
	//if we dont sepcify value to fifth it reused the above const exp
	fifth
	sixth
)

//iota gets reset with new constant block
const (
	seventh = iota
)*/

func main() {

	/*Variable declaration
	once declared must be used otherwise compilor throws error
	var i int //initialised to 0
	i = 42

	should specify size of float in go
	init with declaration
	var i float32 = 3.14

	implict initialisation - go implies the datatype based on what data is assigned
	firstName := "Arthur"
	b := true //boolean
	complex data type
	c := complex(3,4)
	r, im := real(c), imag(c)
	*/

	/*Pointers
	no pointer arithemetic in go
	var firstName *string
	*firstName = "Arthur" - error assigning aruthur to uninitialized pointer

	var firstName *string = new(string)
	*firstName = "Arthur"
	ptr := &firstName
	fmt.Println(*firstName)
	fmt.Println(ptr, **ptr)*/

	/*constants
	initialize during declaration
	value of constant has to be setermined at compile time
	function return values are noy determined till run time so cant assign them to const
	//const pi = 3.14 //impilicity typed const (no datatype)
	const pi int = 3
	fmt.Println(pi + 1)
	fmt.Println(float32(pi) + 3.1)*/

	/*iota
	each time we use iota it increments the value
	fmt.Println(first, second, third, fourth, fifth, sixth, seventh)*/

	/*arrays
	//var arr [3]int
	arr := [3]int{1, 2, 3}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr, arr[1])*/

	/*slices
	arr := [3]int{1, 2, 3}
	slice := arr[:] //beginning to end
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice) // same data in both -> slice pointing to data array is keeping

	slice := []int{1, 2, 3} //no size specified
	slice = append(slice, 4)
	fmt.Println(slice)
	s2 := slice[1:2]//upto not including 2
	fmt.Println(s2)*/

	/*maps
	m := map[string]int{"foo": 42} //key -> string ; value -> int
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m["foo"])
	//deleting elements from map
	delete(m, "foo")
	fmt.Println(m)*/

	/*struct
	allows discrete datatypes together
	type user struct { // we are defining a type user which is going to be struct
		ID        int
		firstName string
		lastName  string
	}
	var u user //each value in u is initialized to 0 (int) null(string)
	u.ID, u.firstName, u.lastName = 1, "Tulasi", "Sree"
	fmt.Println(u)

	u2 := user{ID: 1, firstName: "Arthur",
		lastName: "Dent",
	}
	fmt.Println(u2)*/

	/*startWebServer as an entity or object to be passed
	// _ write only variable which just dumps value into it (when u dont need the return value)
	port := 3000
	_, err := startWebServer(port, 2) //call
	fmt.Println(err)*/

	/*u := models.User{
		Id:        2,
		LastName:  "Srer",
		FirstName: "Tulasi",
	}
	fmt.Println(u)*/

	/* LOOPS
	only one loop -> for
	//var i int
	for i < 5 {
		i++
		if i == 3 {
			//break
			continue
		}
		println("continuing")
	}
	//for with post clause implements always we fininsh one iteration
	// for ; i < 5; i++ {
	for i := 0; i < 5; i++ {
		println(i)
	}
	//infinite loop
	for i < 5 {}
	//looping through collections
	slice := []int{1, 2, 3}
	ports := map[string]int{"http": 80, "https": 443}
	// for i := 0; i < len(slice); i++ {
	// 	println(slice[i])
	// }
	for i, v := range slice {
		println(i, v)
	}
	for i, v := range ports {
		println(i, v)
	}
	for i := range ports { //printing only keys
		println(i)
	}
	for _, v := range ports { //printing only values
		println(v)
	}*/

	/*Panic Function
	panic("Something bad just happened") //stops execution of current routine
	println("after panic")*/

	/*if stmt
	if a==b {

	} else if a>b {
	}else {
		}*/

	/* SWITCH
	//no need of break implicitly breaks
	//if we want to fall throught use fallthrough keyword
	switch x {
		case "a":
			//code
			fallthorough
		case "b":
			//code
		default:
			//code
	}*/

	controllers.RegisterControllers()
	//1 para. ip address we are listening on
	//2 para. servermux or servemultiplexer handles all request that are coming in and also the high level routing
	http.ListenAndServe(":3000", nil) //nil -> default servermux

}

/*func startWebServer(port int, numberOfRetries int) {
error return type specifies something bad happened
error is a pointer data type very least we can return nil
func startWebServer(port, numberOfRetries int) error { //both parameters are of same type so you can have one int
	fmt.Println("Starting server...")
	fmt.Println("Server Started on port", port, "number of retries", numberOfRetries)
	//return nil
	return errors.New("Something went wrong")
}

//more than one return values
func startWebServer(port, numberOfRetries int) (int, error) { //both parameters are of same type so you can have one int
	fmt.Println("Starting server...")
	fmt.Println("Server Started on port", port, "number of retries", numberOfRetries)
	//return nil
	return port, nil
}*/
