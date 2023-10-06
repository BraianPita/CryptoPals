package basics

import (
	"encoding/hex"
)

func safeByte(bytes []byte, index int) byte {
	if index >= len(bytes) {
		return 0
	} else {
		return bytes[index]
	}
}

func calculatePadding(data []byte, i int, remains int) []int {

	d4 := 64 // 64 is padding
	d3 := 64
	d2 := 64
	d1 := 64

	temp := byte(0)

	d1 = int((data[i] >> 2))
	temp = (data[i] << 4) | (safeByte(data, i+1) >> 4)
	temp = temp & 0x3F
	d2 = int(temp)
	if remains == 1 {
		return []int{d1, d2, d3, d4}
	}

	temp = (safeByte(data, i+1) << 2) | (safeByte(data, i+2) >> 6)
	temp = temp & 0x3F
	d3 = int(temp)
	if remains == 2 {
		return []int{d1, d2, d3, d4}
	}

	temp = safeByte(data, i+2)
	temp = temp & 0x3F
	d4 = int(temp)
	if remains == 3 {
		return []int{d1, d2, d3, d4}
	}

	return []int{d1, d2, d3, d4}
}

// import "fmt"
// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
func HexToBase64Challenge(input string) string {
	output := ""

	map64 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

	data, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("len = %b", data)

	indexes := make([]int, 0, 0)

	i := 0

	for i = 0; i <= len(data)-3; i += 3 {
		d1 := data[i] >> 2
		d2 := (data[i] << 4) | (data[i+1] >> 4)
		d2 = d2 & 0x3F
		d3 := (data[i+1] << 2) | (data[i+2] >> 6)
		d3 = d3 & 0x3F
		d4 := data[i+2] & 0x3F

		indexes = append(indexes, int(d1), int(d2), int(d3), int(d4))
	}

	// Calculating padding if necessary
	remains := len(data) - i

	if remains > 0 {
		indexes = append(indexes, calculatePadding(data, i, remains)...)
	}

	for _, mapIndex := range indexes {
		output += string(map64[mapIndex])
	}

	return output
}
