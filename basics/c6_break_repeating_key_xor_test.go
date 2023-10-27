package basics

import (
	"reflect"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	m1 := "this is a test"
	m2 := "wokka wokka!!!"

	if HammingDistance(m1, m2) != 37 {
		t.Error("C6 Error - Hamming distance function returns wrong value")
	}
}

func TestTransposeChunks(t *testing.T) {
	input := [][]byte{{2, 4, 6}, {4, 6, 8}, {9, 0}}
	output := [][]byte{{2, 4, 9}, {4, 6, 0}, {6, 8}}

	result := transposeChunks(input)

	if !reflect.DeepEqual(result, output) {
		t.Error("Result is not right for transpose.")
	}
}

func TestBreakRepeatingKeyXor(t *testing.T) {

	lines, err := ReadLines("data/c6_encrypted_data.txt")

	if err != nil {
		t.Fatalf("Error reading data from file: %v", err)
	}

	data := Base64ToBytes(lines)

	BreakRepeatingKeyXor(data)
}
