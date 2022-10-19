package main

import (
	"fmt"
	twelweDays "github.com/kramar/go-lessons/week_1/twelve-days"
	"github.com/kramar/go-lessons/week_1/leap"
	"github.com/kramar/go-lessons/week_1/pangram"
)

func main() {
	fmt.Println(twelweDays.Song())
	fmt.Println(leap.IsLeapYear(2000))
	fmt.Println(pangram.IsPangram("This is test string"))
}