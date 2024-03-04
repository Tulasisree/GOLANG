package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	/*file, err := os.Open("family.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var age int
		var name string

		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)

		if err != nil {
			panic(err)
		}

		if n == 2 {
			fmt.Println(age, name)
		}
	}*/

	/*fmt.Println("name please? ")
	var fname, lname string
	fmt.Scanln(&fname, &lname) //space separated inp
	fmt.Printf("%s%s", fname, lname)*/

	/*output := fmt.Sprintf("hello this is tulasi")
	fmt.Printf("\033[1;34n%s\033[0n", output)*/

	/*file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	// log.Println("This is a log message!")

	InfoLogger := log.New(file, "info:", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger2 := log.New(file, "info:", log.LUTC|log.Lmicroseconds|log.Llongfile)
	log.Println(InfoLogger)
	log.Println(InfoLogger2)*/

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("We did not create a trace file %v\n", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("We did not close a trace file %v\n", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("trace failed to start %d\n", err)
	}

	defer trace.Stop()

	AddRandomNumber()

}

func AddRandomNumber() {
	fn := rand.Intn(100)
	sn := rand.Intn(100)

	time.Sleep(2 * time.Second)

	var result = fn + sn
	fmt.Println("result :", result)
}
