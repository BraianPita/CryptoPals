package basics

import (
	"bufio"
	"os"
)

// read lines of a plain file into array of strings
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Similar to XorCypherBestGuess but for already decoded strings
func rankStringList(input []string) string {
	letterFreqList := make(map[string]float64)

	for _, message := range input {
		letterFreqList[message] = CalculateEnglishScore(message)
	}

	return RankByLetterFreq(letterFreqList)
}

// take list of encoded hexes and returns the best guess english message of them
// by first decoding and finding the best guess at decoded message from each
func DecodeHexLetterFrequencyGuess(input []string) string {
	// list to get the best guest decoding for each string
	bestList := make([]string, len(input))

	for i, message := range input {
		bestList[i] = XorCypherBestGuess(message)
	}

	return rankStringList(bestList)
}
