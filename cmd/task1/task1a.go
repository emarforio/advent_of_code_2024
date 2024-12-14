package task1

import (
	"adventsofcode/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Task1a() int {
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

	// Arrays containing coordinate data
	var firstDistanceList [1005]int
	var secondDistanceList [1005]int

	// Read contents from file
	r := bufio.NewReader(f)
	for i := 0; i < len(firstDistanceList); i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
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
		firstVal, secondVal, firstIndex, secondIndex := utils.FindMaxima(firstDistanceList[:], secondDistanceList[:])

		if firstVal == -1 {
			break
		}

		totalSum += int(math.Abs(float64(firstVal - secondVal)))

		firstDistanceList[firstIndex] = 0
		secondDistanceList[secondIndex] = 0
	}
	f.Close()

	return totalSum
}
