package main

import (
	"adventsofcode/task1"
	"adventsofcode/task2"
	"fmt"
	"os"
)

func main() {
	task := "1a"
	result := 0

	if len(os.Args) > 1 {
		task = os.Args[1]
	}

	// Run task
	fmt.Println()
	fmt.Printf("##### Task %s #####\n", task)
	switch task {
	case "1a":
		result = task1.Task1a()

	case "1b":
		result = task1.Task1b()

  case "2a":
    result = task2.Task2a()

	}

	fmt.Printf("%d\n", result)
	fmt.Println()
}
