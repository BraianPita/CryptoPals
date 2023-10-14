package basics

import "testing"

func TestHammingDistance(t *testing.T) {
	m1 := "this is a test"
	m2 := "wokka wokka!!!"

	if HammingDistance(m1, m2) != 37 {
		t.Error("C6 Error - Hamming distance function returns wrong value")
	}
}
