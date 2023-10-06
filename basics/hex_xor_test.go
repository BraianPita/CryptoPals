package basics

import "testing"

func TestHexXor(t *testing.T) {
	h1 := "1c0111001f010100061a024b53535009181c"
	h2 := "686974207468652062756c6c277320657965"

	result := HexXor(h1, h2)

	if result != "746865206b696420646f6e277420706c6179" {
		t.Fatal("Xor result does not match.")
	}
}
