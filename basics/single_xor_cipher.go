package basics

import "encoding/hex"

func SingleXor(hexString string, char byte) string {
	data, err := hex.DecodeString(hexString)
	secret := byte(char)

	result := ""

	if err != nil {
		panic("Error decoding hex string")
	}

	for _, val := range data {
		result += string(val ^ secret)
	}

	return result

}
