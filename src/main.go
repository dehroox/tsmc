package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var location string

	fmt.Print("What file(s) should be compiled? ")
	fmt.Scanln(&location)
	fmt.Println("Got it! Compiling now... ")

	file, err := os.Open(location)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
