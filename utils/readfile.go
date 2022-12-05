package utils

import (
	"bufio"
	"log"
	"os"
)

var file *os.File

func GetScanner(filepath string) *bufio.Scanner {
	file := openOS(filepath)
	return bufio.NewScanner(file)
}

func openOS(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func CloseOS() {
	file.Close()
}
