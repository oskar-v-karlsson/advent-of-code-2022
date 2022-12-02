package main

import "fmt"

func main() {
	fmt.Printf("Elf with largest amount of calories is carrying: %v calories!\n", CalorieCount())
	fmt.Printf("The three elves with largest amount of calories are carrying: %v calories!\n", CalorieCountThree())
}
