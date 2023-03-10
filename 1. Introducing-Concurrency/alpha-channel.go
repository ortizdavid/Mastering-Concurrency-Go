package main

import (
	"fmt"
	"sync"
	"runtime"
	"strings"
)

var initialString string
var finalString string

var stringLenght int


func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup)  {
	letter := <- letterChannel
	finalString += letter
	wg.Done()
}

func capitalize(letterChannel chan string, currentLetter string, wg *sync.WaitGroup)  {
	thisLetter := strings.ToUpper(currentLetter)
	wg.Done()
	letterChannel <- thisLetter
}

func sampleString() string {
return `Four score and seven years ago our fathers
brought forth on this continent, a new nation, conceived in
Liberty, and dedicated to the proposition that all men are
created equal.`
}


func main() {
	
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	
	initialString = sampleString()
	initialBytes := []byte(initialString)

	var letterChannel chan string = make(chan string)
	stringLenght = len(initialBytes)

	for i := 0; i < stringLenght; i++ {
		wg.Add(2)

		go capitalize(letterChannel, string(initialBytes[i]), &wg)
		go addToFinalStack(letterChannel, &wg)

		wg.Wait()
	}

	fmt.Println(finalString)
}