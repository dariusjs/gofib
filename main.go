package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many Fibonacci interations to run")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if integer, err := strconv.Atoi(text); err == nil {
			fmt.Printf("integer=%d, type: %T\n", integer, integer)
			return integer
		} else {
			panic("Input is not an integer.")
		}
	}
}

func main() {
	iterations := readInput()
	fmt.Println(iterations)

	counter := 0
	startingInteger := 0
	secondInteger := 1

	for counter < iterations {
		fmt.Println("fibonacci number:", startingInteger)

		sequencer := startingInteger + secondInteger
		startingInteger = secondInteger
		secondInteger = sequencer
		counter += 1
	}
}
