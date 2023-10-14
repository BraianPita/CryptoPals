package basics

import (
	"fmt"
	"testing"
)

func TestSingleXorCypher(t *testing.T) {
	secretMessage := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decodedString := XorCypherBestGuess(secretMessage)

	fmt.Printf("SECRET MESSAGE : %v\n", decodedString)

}
