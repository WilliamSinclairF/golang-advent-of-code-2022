package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInputFromFile() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func main() {
	input := readInputFromFile()
	lines := strings.Split(string(input), "\n")

	count := 0
	highest := 0

	for _, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			// if I can't parse, this is a new line, meaning new elf lol
			fmt.Println("next elf")
			count = 0
			continue
		}

		count += num

		if count > highest {
			highest = count
		}

		fmt.Println("Current count: ", count)
		fmt.Println("Current highest: ", highest)
	}
}
