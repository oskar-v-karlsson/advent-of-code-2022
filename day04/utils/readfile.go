package utils

import (
	"bufio"
	"log"
	"os"
)

var file *os.File

func getScanner(filename string) *bufio.Scanner {
	file := openOS(filename)
	return bufio.NewScanner(file)
}

func openOS(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func closeOS() {
	file.Close()
}
