package basics

import (
	"encoding/hex"
)

func RepeatingKeyXor(key string, data string) string {
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
