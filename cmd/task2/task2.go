package task2

import (
	"os"
	"path/filepath"
	"runtime"
)

func Task2a() int {
	// Load dataset
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: Unable to get file information")
	}

	dir := filepath.Dir(filename)
	filepath := filepath.Join(dir, "task1.txt")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	f.Close()
	safeLevelCount := 0
	return 0
}

func processLine(line string) []int {
	panic("Not implemented")
}

func checkLevel(level [5]int) {
	panic("Not implemented")
}
