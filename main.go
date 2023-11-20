package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	if *lines {
		fmt.Println(countLines(os.Stdin))
	} else if *bytes {
		fmt.Println(countBytes(os.Stdin))
	} else {
		fmt.Println(countWords(os.Stdin))
	}
}

func countLines(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func countWords(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}

func countBytes(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanBytes)

	byteCount := 0

	for scanner.Scan() {
		byteCount++
	}

	return byteCount
}
