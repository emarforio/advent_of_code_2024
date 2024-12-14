package main

import (
	"fmt"
	"os"
)

import (
	"adventsofcode/task1"
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

	}

	fmt.Printf("%d\n", result)
	fmt.Println()
}
