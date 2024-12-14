package utils

func FindMaxima(firstSlice []int, secondSlice []int) (firstMaxVal int, secondMaxVal int, firstIndex int, secondIndex int) {

	firstMaxVal = firstSlice[0]
	secondMaxVal = firstSlice[0]

	firstIndex = 0
	secondIndex = 0

	for i := range firstSlice {

		firstValue := firstSlice[i]
		secondValue := secondSlice[i]

		if firstValue > firstMaxVal {
			firstMaxVal = firstValue
			firstIndex = i
		}

		if secondValue > secondMaxVal {
			secondMaxVal = secondValue
			secondIndex = i
		}
	}

	if firstMaxVal == 0 && secondMaxVal == 0 {
		return -1, -1, -1, -1
	}

	return firstMaxVal, secondMaxVal, firstIndex, secondIndex
}

func SimilarityScore(toCheck int, slice []int) int {
	noOccurrences := 0

	for i := 0; i < len(slice); i++ {
		if toCheck == slice[i] {
			noOccurrences += 1
		}
	}
	return toCheck * noOccurrences
}

