package basics

import (
	"encoding/hex"
)

func RepeatingKeyXorHex(key string, data string) string {
	dataBytes := []byte(data)
	keyBytes := []byte(key)

	numKeyBytes := len(keyBytes)

	index := 0

	for idx, dataByte := range dataBytes {
		dataBytes[idx] = dataByte ^ keyBytes[index]
		index = (index + 1) % numKeyBytes
	}

	return hex.EncodeToString(dataBytes)
}

// repeating key xor but on raw bytes
func RepeatingKeyXor(key string, data []byte) []byte {
	keyBytes := []byte(key)
	dataBytes := make([]byte, len(data))

	numKeyBytes := len(keyBytes)

	index := 0

	for idx, dataByte := range data {
		dataBytes[idx] = dataByte ^ keyBytes[index]
		index = (index + 1) % numKeyBytes
	}

	return dataBytes
}
