package task2

import (
	"bufio"
	"fmt"
  "math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Task2a() int {
	// Load dataset
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: Unable to get file information")
	}

	dir := filepath.Dir(filename)
	filepath := filepath.Join(dir, "task2.txt")

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	safeLevelCount := 0

  for scanner.Scan() {
    reportLine := scanner.Text()
    if (reportIsValid(reportLine, true)) {
      safeLevelCount += 1
    }
  }

	f.Close()

	return safeLevelCount
}


func reportIsValid(reportLine string, checkAgain bool) bool {
  numbers_str := strings.Fields(reportLine)
  
  // Used to check consistent direction
  isDescending := false
  isAscending := false

  for i := 1; i < len(numbers_str); i++ {
    // Parse
    num1, err1 := strconv.Atoi(numbers_str[i-1])
    num2, err2 := strconv.Atoi(numbers_str[i])

		if err1 != nil || err2 != nil {
			errorMessage := fmt.Sprintf("Could not parse line %s\n", reportLine)
      fmt.Println(errorMessage)
		}
    
    // Collect difference and validate
    numberDiff := num2 - num1

    if numberDiff == 0 || int(math.Abs(float64(numberDiff))) >= 4 {
      if checkAgain {
        return false
      }
      return false
    } else if numberDiff < 0 {
      isDescending = true
    } else if numberDiff >= 0 {
      isAscending = true
    }
  }

  return isAscending != isDescending
}


