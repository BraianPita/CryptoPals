package basics

import (
	"fmt"
	"math"
	"testing"
)

func TestSingleXorCypher(t *testing.T) {
	secretMessage := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	for x := 0; x < int(math.Pow(2, 32)); x++ {
		fmt.Println(SingleXor(secretMessage, byte(x)))
	}
}
