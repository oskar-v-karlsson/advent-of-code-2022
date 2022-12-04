package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var file *os.File

func GetScanner(filepath string) *bufio.Scanner {
	fmt.Println(os.Getwd())
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
