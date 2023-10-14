package basics

import (
	"encoding/hex"
	// "fmt"
)

func HexXor(hex1 string, hex2 string) string {
	data1, err1 := hex.DecodeString(hex1)
	data2, err2 := hex.DecodeString(hex2)

	if err1 != nil || err2 != nil {
		panic("There is an issue decoding hex strings.")
	}

	xor := []byte{}

	for i := 0; i < len(data1); i++ {
		xor = append(xor, data1[i]^data2[i])
	}
	result := hex.EncodeToString(xor)
	// fmt.Println(result)

	return result

}
