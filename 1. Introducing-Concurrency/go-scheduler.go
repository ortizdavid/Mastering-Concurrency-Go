package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"strconv"
	"runtime"
)

type Job struct {
	i int
	max int
	text string
}

func outputText(j *Job)  {
	fileName := j.text + ".txt"
	fileContents := ""
	for j.i < j.max {
		time.Sleep((1000 * time.Millisecond))
		fileContents += j.text
		fmt.Println(j.text)
		j.i++
	}
	err := ioutil.WriteFile(fileName, []byte (fileContents), 0644)
	if err != nil {
		panic("Something went awry")
	}
}

func outputTextExample () {
	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world := new(Job)
	world.text = "wello"
	world.i = 0
	hello.max = 5

	go outputText(hello)
	go outputText(world)
}

func showNumber(num int)  {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
	time.Sleep(time.Millisecond * 10)
}

func showNumberExamples() {
	runtime.GOMAXPROCS(0)
	iterations := 10

	for i := 0; i <= iterations; i++ {
		go showNumber(i)
	}
	fmt.Println("GoodBye!")
}

func listThreads() int  {
	threads := runtime.GOMAXPROCS(0)
	return threads
}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Printf("%d threads(s) available to Go.", listThreads())
}