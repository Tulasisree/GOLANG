package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2,2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	t := time.Now()
	fmt.Print(t.Year(), "\n", t.Month(), "\n", t.Day(), "\n")
	//time formats
	fmt.Println(t.Format(time.ANSIC)) //time.RFC3339 , time.UnixDate, time.Kitchen

	//creating own time formats
	//Mon Jan 2 15:04:05 MST 2006
	fmt.Println(t.Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Println(t.Format("Monday, January 2 in the year 2006"))
	startDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC) //year month day hour min sec nanosec timeformat
	bedtime := time.Date(2023, 02, 20, 23, 00, 00, 00, time.Local)
	fmt.Println(startDate.Format("Monday, January 2 in the year 2006"))
	fmt.Println(startDate.Format(layoutISO))
	fmt.Println(startDate.Format(layoutEU))
	fmt.Println(startDate.Format(layoutUS))

	elapsed := time.Since(startDate)
	fmt.Printf("%v has passed since %v", elapsed, startDate) //elapsed.Hours(),minutes(),seconds() ,

	future := t.AddDate(0, 6, 0) //year months days
	past := t.AddDate(0, -6, 0)
	//t.Add(6*time.Hour)
	fmt.Printf("In six months it will be %v\n", future.Format("Monday, January 2"))
	fmt.Printf("before six months it was %v\n", past.Format("Monday, January 2"))

	fmt.Printf("there is %.0f hours until bed time", time.Until(bedtime).Hours())
}
