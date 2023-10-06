package basics

import (
	"fmt"
	"testing"
)

// test base64 challenge
func TestHexToBase64Challenge(t *testing.T) {
	result := HexToBase64Challenge("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	fmt.Print("result = " + result)

	if result != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatalf("Base 64 conversion does not match.")
	}
}
