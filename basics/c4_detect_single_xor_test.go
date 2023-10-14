package basics

import (
	"fmt"
	"testing"
)

func TestRankStringList(t *testing.T) {

	encodedMessages, err := ReadLines("data/cryptopals-challenge4-input.txt")

	if err != nil {
		t.Errorf("Could not read lines: %v", err)
	}

	result := DecodeHexLetterFrequencyGuess(encodedMessages)

	fmt.Printf("BEST GUESS : %v\n", result)
}
