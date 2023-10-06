package main

import (
	"fmt"
	"localhost/crypto/basics"
)

func main() {
	result := basics.HexToBase64Challenge("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f")

	// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
	fmt.Print("RESULT " + result)
}
