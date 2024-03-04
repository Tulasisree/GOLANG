package main

import (
	"fmt"
	organization "github/interfaces/person"
)

func main() {
	p := organization.NewPerson(
		"James",
		"Wilson",
		organization.NewEuropeanUnionIdentifier(123, "Germany"),
	)
	err := p.SetTwitterHandler("@jam_wills")
	fmt.Printf("%T\n", p.TwitterHandler())

	if err != nil {
		fmt.Printf("An error occured setting twitter handler: %s\n", err.Error())
	}
	println(p.ID())
	println(p.Country())
}

//ssn:=people.NewSocialSecurityNumber("123-456")
//eu := people.NewEuropeanUnionIdentifier("123","France")
//eu2 := peoplr.NewEuropeanUnionIdentifier("123","France")

//if ssn==eu allowed
//if eu ==eu2 allowed (lets say be include a slice in eu struct)
//then ssn==eu doesnt throw any error
//but eu==eu2 throws runtime error ..compilor sees that both are of type eu and goes for comparing
//during run time it sees that they are uncomparable types

//ZERO VALUE COMPARISON
//name1:=Name{First:"",Last:""}
//if name1 == (Name{})
//if name1 == nil (name1 := &Name{First:"",Last:""}) returning a pointer
//but value comparision is much better since it is cheaper memory wise as it stores on stack but pointer go on heap
