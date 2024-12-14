package task1

import (
	"adventsofcode/utils"
	"bufio"
	"fmt"
	"os"
    "path/filepath"
    "runtime"
	"strconv"
	"strings"
)

func Task1b() int {
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

	// To contain number series
	var firstDistanceList [1005]int
	var secondDistanceList [1005]int

	// Read contents from file
	r := bufio.NewReader(f)
	for i := 0; i < len(firstDistanceList); i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

        numbers := strings.Fields(line)
		if len(numbers) < 2 {
			break
		}

		// Parse numbers
		firstNum, err1 := strconv.Atoi(numbers[0])
		secondNum, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
            errorMessage := fmt.Sprintf("Could not parse line %s\n", line)
            panic(errorMessage)
		}

		firstDistanceList[i] = firstNum
		secondDistanceList[i] = secondNum
	}

	// Collect sum
	totalSum := 0
	for i := 0; i < len(firstDistanceList); i++ {
		totalSum += utils.SimilarityScore(firstDistanceList[i], secondDistanceList[:])
	}
	f.Close()

	return totalSum
}

