package basics

import (
	"fmt"

	"github.com/forgoer/openssl"
)

func DecryptAesEcb(data []byte, key string) string {
	keyBytes := []byte(key)
	decrypted, err := openssl.AesECBDecrypt(data, keyBytes, "")

	if err != nil {
		panic(fmt.Sprintf("Error decrypting AES ECB: %v", err))
	}
	return string(decrypted)
}
