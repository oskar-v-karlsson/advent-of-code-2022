package utils

import (
	"bufio"
	"log"
	"os"
)

var file *os.File

func GetScanner() *bufio.Scanner {
	file := openOS()
	return bufio.NewScanner(file)
}

func openOS() *os.File {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func CloseOS() {
	file.Close()
}
