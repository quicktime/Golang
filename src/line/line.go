package main

import (
	"bufio"
	"fmt"
	"os"
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
  for i, line := range lines {
	fmt.Println(i, line)
  }
}