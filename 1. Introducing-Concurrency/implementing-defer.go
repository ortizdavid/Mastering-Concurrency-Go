package main

import (
	"os"
	//"fmt"
)

func main() {
	file, _ := os.Create("defer.txt")

	defer file.Close()

	for {
		break
	}

	/*aValue := new(int)

	defer fmt.Println(*aValue)
	for i := 0; i < 100; i++ {
		*aValue++
	}*/
}