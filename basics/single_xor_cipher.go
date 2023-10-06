package basics

import "encoding/hex"

func SingleXor(hexString string, char byte) string {
	data, err := hex.DecodeString(hexString)
	secret := byte(char)

	if err != nil {
		panic("Error decoding hex string")
	}

	for idx, val := range data {
		data[idx] = val ^ secret
	}

	return hex.EncodeToString(data)

}
