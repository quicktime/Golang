package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("dat")
	if err != nil {
		fmt.Println("butts")
	}
	sort.Strings(lines)

	size := len(lines)
	
	s1 := []string{"The longest line in the file is:", lines[size-1]}
	
	fmt.Println(lines)
	fmt.Print("The number of lines in the file is:   ")
	fmt.Println(len(lines))
	fmt.Println(strings.Join(s1, "   "))
}
