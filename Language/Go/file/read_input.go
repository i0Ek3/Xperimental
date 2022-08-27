package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	str1, str2, str, inputStr string
	i                         int
	f                         float32
	input                     = "56 / 56.1 / Go"
	format                    = "%d / %f / %s"
	inputReader               *bufio.Reader
	err                       error
)

func read_input1() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&str1, &str2)
	fmt.Printf("Hi %s %s!\n", str1, str2)
	fmt.Sscanf(input, format, &i, &f, &str)
	fmt.Println("From the string we read: ", i, f, str)
}

func read_input2() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some words: ")
	inputStr, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s.\n", inputStr)
	}
}

func main() {
	read_input1()
	read_input2()
}
