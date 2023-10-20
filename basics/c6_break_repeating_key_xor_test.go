package basics

import (
	"encoding/base64"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	m1 := "this is a test"
	m2 := "wokka wokka!!!"

	if HammingDistance(m1, m2) != 37 {
		t.Error("C6 Error - Hamming distance function returns wrong value")
	}
}

func TestBreakRepeatingKeyXor(t *testing.T) {

	lines, err := ReadLines("data/c6_encrypted_data.txt")
	data := make([]byte, 0)

	if err != nil {
		t.Fatalf("Error reading data from file: %v", err)
	}

	for _, line := range lines {
		curr, err := base64.StdEncoding.DecodeString(line)
		if err != nil {
			t.Fatalf("Error decoding base64 string: %v", err)
		}
		data = append(data, curr[:]...)
	}

	BreakRepeatingKeyXor(data)
}
