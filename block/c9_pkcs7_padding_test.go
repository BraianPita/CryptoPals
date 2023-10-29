package block

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddPkcs7Padding(t *testing.T) {
	input := "YELLOW SUBMARINE"
	blockSize := uint8(20)

	result := AddPkcs7Padding([]byte(input), blockSize)

	fmt.Printf("Input: %v\nPadded : %v", []byte(input), result)
	if !reflect.DeepEqual(result, append([]byte(input), 0x04, 0x04, 0x04, 0x04)) {
		t.Error("AddPkcs7Padding - Expected value does not match the result.")
	}

}
