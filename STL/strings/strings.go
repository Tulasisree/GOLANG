package main

import (
	"fmt"
	"strings"
)

func main() {
	// str1 := "tulasi"
	// str2 := "TUlasi"

	// fmt.Println(CompareCaseIns(str1, str2))

	// str := "This|is|Tulasi"
	// strcoll := strings.Split(str, "|")
	// for _, s := range strcoll {
	// 	println(s)
	// }
	// fmt.Println(strings.SplitAfter(str, "|"))     //includes |
	// fmt.Println(strings.SplitAfterN(str, "|", 2)) // number of splits

	// str := "there was a tree and the was on the..."
	// if len(os.Args) > 1 {
	// 	result := strings.Contains(str, os.Args[1])
	// 	if result {
	// 		fmt.Println("contains the string")
	// 	} else {
	// 		fmt.Println("doesnt contain the string")
	// 	}
	// } else {
	// 	fmt.Println("enter a search term")
	// }

	// str := "this language is called go but never mentioned where to go (sreeg) streegs"
	// r, _ := regexp.Compile(`s([a-z]+)g`) // starts with a s ends with a g  //gives both sreeg from (sreeg) sreeg from sreegs
	// //to get only the words which start with s and end with g
	// r1, _ := regexp.Compile(`s(\w[a-z]+)g\b`) // \w word \b word boundary
	// fmt.Println(r.MatchString(str))
	// fmt.Println(r.FindAllString(str, -1)) // -1 return all the results
	// fmt.Println(r.FindAllString(str, -1))
	// fmt.Println(r1.FindAllString(str, -1))
	// fmt.Println(r1.FindStringIndex(str)) //start at index i and ends at index j

	// newText := r.ReplaceAllString(str, "sree")
	// fmt.Println(newText)

	// str := "       tulasi hi       "
	// fmt.Println(str, "\"")
	// str2 := strings.TrimSpace(str)
	// str3 := strings.TrimLeft(str, " ")
	// fmt.Println(str2, "\"")
	// fmt.Println(str3, "\"")
	// url := "https://www.pluralsight.com"
	// domain := strings.TrimPrefix(url, "https://") //TrimSuffix
	// fmt.Println(domain)

	str := "welcome to the dollhouse"
	fmt.Println(strings.ToLower(str), strings.ToUpper(str), "\n", strings.Title(str), "\n", ProperTitle(str))
}

// func CompareCaseIns(a, b string) bool {
// 	if len(a) == len(b) {
// 		if strings.ToLower(a) == strings.ToLower(b) {
// 			return true
// 		}
// 	}
// 	return false
// }

func ProperTitle(str string) string {
	words := strings.Fields(str)
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
