package main

import (
	"fmt"
	"strings"
)


func shortenString(message string) func() string  {
	return func() string {
		messageSlice := strings.Split(message, " ")
		wordLenght := len(messageSlice)
		if wordLenght < 1 {
			return "Nothin Left"
		} else {
			messageSlice = messageSlice[:(wordLenght-1)]
			message = strings.Join(messageSlice, " ")
			return message
		}
	}
}

func main() {
	myString := shortenString("Welcome to concurrency in Go! ...")
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
	fmt.Println(myString())
}