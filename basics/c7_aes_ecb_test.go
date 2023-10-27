package basics

import (
	"fmt"
	"testing"
)

func TestAesEcb(t *testing.T) {
	dataLines, err := ReadLines("data/c7_data.txt")
	key := "YELLOW SUBMARINE"

	if err != nil {
		t.Error("Error reading file for AES ECB.")
	}

	data := Base64ToBytes(dataLines)
	decrypted := DecryptAesEcb(data, key)

	fmt.Printf("DECODED:\n%v", decrypted)
}
