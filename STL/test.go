package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// "bufio"
// "os"
// "strconv"
// "runtime"

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("what is your name?")
	// text, _ := reader.ReadString('\n')   // the newline here waits for use to enter which then populates the buffer with our input
	// fmt.Printf("my name is %v !!", text) // notice that there is newline after name ie even taking it as delimiter during input it is still keeping it as text while printing the name
	// fmt.Println("current version is :", runtime.Version())

	//arguments
	/*args := os.Args //args is a slice with collection of all parameters that are sent to the application
	args1 := os.Args[1:]

	if len(args1) == 1 && args1[0] == "/help" {
		fmt.Printf("usage: <mealtotal> <tip>\n example: 20 10")
	} else {
		if len(args1) == 0 {
			fmt.Println("You must enter arguments! type /help for help")
		} else {
			mealTotal, _ := strconv.ParseFloat(args1[0], 32) //arguments return as string
			tipAmount, _ := strconv.ParseFloat(args1[1], 32)
			fmt.Printf("total amount %.2f", (mealTotal + tipAmount))
		}
	}
	fmt.Println(args)*/

	/*archPtr := flag.String("arch", "x86", "cpu type") //datatype and name of the flag , value(default value) , text string to describe what the flag is (-h shows this description ex: ./test -h)
	flag.Parse()                                      //to create the flag
	//./test gives 32bit.... as x86 is default
	//./test -arch amb64
	switch *archPtr {
	case "x86":
		fmt.Println("running in 32 bit mode")
	case "amb64":
		fmt.Println("running in 64 bit mode")
	case "imb64":
		fmt.Println("remember imb64")
	}*/

	/*var name string
	fmt.Println("ur name?")
	// inputs, _ := fmt.Scanf("%s", &name)
	inputs, _ := fmt.Scanf("%q", &name)

	switch inputs {
	case 0:
		fmt.Println("u must enter a name")
	case 1:
		fmt.Println("hello ", name)
	}*/

	/*scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "/quit" {
			fmt.Println("quitting ....")
			os.Exit(1)
		} else {
			fmt.Println("you have typed ", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil { // check the statement in if
		fmt.Println(err)
	}*/

	/*file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}*/

	args := os.Args[1:]
	if len(args) == 0 || args[0] == "/help" {
		fmt.Println("usage: test <input file>")
	} else {
		fmt.Println("how would you like to see the text: 1.caps , 2.title , 3.lower case")
		var i int
		fmt.Println("enter your choice:")
		_, err := fmt.Scanf("%v", &i)

		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			switch i {
			case 1:
				fmt.Println(strings.ToUpper(scanner.Text()))
			case 2:
				fmt.Println(strings.ToTitle(scanner.Text()))
			case 3:
				fmt.Println(strings.ToLower(scanner.Text()))
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

	}

}
