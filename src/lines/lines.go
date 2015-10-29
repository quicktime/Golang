package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("dat")
	check(err)
	
	fileScanner := bufio.NewScanner(file)
	check(err)
	
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("Number of lines in the file: ", lineCount)

	file.Close()
}