package basics

import (
	"fmt"
	"sort"
)

// ---------------------------------- START HAMMING DISTANCE -----------------------------------

var count8bits = [256]uint8{
	0x00, 0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x03, 0x01, 0x02, 0x02, 0x03, 0x02, 0x03, 0x03, 0x04,
	0x01, 0x02, 0x02, 0x03, 0x02, 0x03, 0x03, 0x04, 0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05,
	0x01, 0x02, 0x02, 0x03, 0x02, 0x03, 0x03, 0x04, 0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05,
	0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05, 0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06,
	0x01, 0x02, 0x02, 0x03, 0x02, 0x03, 0x03, 0x04, 0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05,
	0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05, 0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06,
	0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05, 0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06,
	0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06, 0x04, 0x05, 0x05, 0x06, 0x05, 0x06, 0x06, 0x07,
	0x01, 0x02, 0x02, 0x03, 0x02, 0x03, 0x03, 0x04, 0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05,
	0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05, 0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06,
	0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05, 0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06,
	0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06, 0x04, 0x05, 0x05, 0x06, 0x05, 0x06, 0x06, 0x07,
	0x02, 0x03, 0x03, 0x04, 0x03, 0x04, 0x04, 0x05, 0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06,
	0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06, 0x04, 0x05, 0x05, 0x06, 0x05, 0x06, 0x06, 0x07,
	0x03, 0x04, 0x04, 0x05, 0x04, 0x05, 0x05, 0x06, 0x04, 0x05, 0x05, 0x06, 0x05, 0x06, 0x06, 0x07,
	0x04, 0x05, 0x05, 0x06, 0x05, 0x06, 0x06, 0x07, 0x05, 0x06, 0x06, 0x07, 0x06, 0x07, 0x07, 0x08,
}

func HammingDistance(m1 string, m2 string) int {

	// make sure m1 is always the smallest string
	if len(m1) > len(m2) {
		temp := m2
		m2 = m1
		m1 = temp
	}

	m1bytes := []byte(m1)
	m2bytes := []byte(m2)

	dist := 0

	for idx := 0; idx < len(m1); idx++ {
		byteDist := m1bytes[idx] ^ m2bytes[idx]
		// use a lookup table to find the number of
		// 1s in the current xor'd byte
		dist += int(count8bits[byteDist])
	}

	// add up the remaining bits from the string len difference
	dist += (len(m2bytes) - len(m1bytes)) * 8

	return dist
}

// --------------------- END HAMMING DISTANCE - START Base64 Decoder -----------------------

func DecodeBase64(base string) []byte {
	data := make([]byte, 0)

	// TODO: Base64 decoder code

	return data
}

// ---------------------------------- START FIND KEYSIZE -----------------------------------

const MAX_KEY_SIZE = 50

type KeysizeWeight struct {
	Key   int
	Value float32
}

type KeysizeWeightList []KeysizeWeight

func (p KeysizeWeightList) Len() int           { return len(p) }
func (p KeysizeWeightList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p KeysizeWeightList) More(i, j int) bool { return p[i].Value > p[j].Value }
func (p KeysizeWeightList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func calculateKeySizeScores(data []byte) KeysizeWeightList {
	keySizeArray := make(KeysizeWeightList, MAX_KEY_SIZE-2)

	for i := 0; i < MAX_KEY_SIZE-2; i++ {
		// +2 to start at keysize 2 (1 and 0 are not options)
		keySizeArray[i] = KeysizeWeight{i + 2, normalizedKeySizeCalculation(i+2, data)}
	}

	sort.Sort(keySizeArray)

	// // Print keysizes in order of score
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("%v - %v\n", keySizeArray[i].Key, keySizeArray[i].Value)
	// }

	return keySizeArray

}

func normalizedKeySizeCalculation(keysize int, data []byte) float32 {
	// invalid inputs
	if len(data) < keysize*4 {
		return 999.9
	}

	// take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes,
	// and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
	slice1 := data[0:keysize]
	slice2 := data[keysize : keysize*2]
	slice3 := data[keysize*2 : keysize*3]
	slice4 := data[keysize*3 : keysize*4]

	dist := float32(HammingDistance(string(slice1), string(slice2))+
		HammingDistance(string(slice2), string(slice3))+
		HammingDistance(string(slice3), string(slice4))) / float32(keysize*3)

	return dist
}

func chunkSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

// transpose the matrix and return it in order as a single slice
// It will ignore any missing elements on the last chunk after transposing
func transposeChunks(chunks [][]byte) [][]byte {
	// for transposing
	lastChunkIndex := len(chunks) - 1
	lastSize := len(chunks[lastChunkIndex])

	// make an array (slice) for data in transposed order
	transposedData := make([][]byte, 0)

	// Iterate each item in block size (keysize len)
	for i := 0; i < len(chunks[0]); i++ {

		// for each keysize byte, add a block
		transposedData = append(transposedData, make([]byte, 0))

		// Iterate each block and get that item
		// except last block, which can be incomplete
		for chunkIndex := 0; chunkIndex < len(chunks)-1; chunkIndex++ {
			// Returns the data at the current transposed location
			transposedData[i] = append(transposedData[i], chunks[chunkIndex][i])
		}
		// after iteration, check if current element exists on last block
		if i < lastSize {
			transposedData[i] = append(transposedData[i], chunks[lastChunkIndex][i])
		}
	}

	return transposedData
}

func BreakRepeatingKeyXor(data []byte) {

	// get all Scores for normalized keysize calculations
	keysizes := calculateKeySizeScores(data)

	// Iterate each keysize from most likely to less likely
	for keysizeRank, keysize := range keysizes {
		fmt.Printf("\nTrying keysize %v - %v\n", keysize.Key, keysize.Value)

		// Separate data into chunks of size 'keysize'
		chunks := chunkSlice(data, keysize.Key)

		// get transposed data
		transposedChunks := transposeChunks(chunks)

		guessedKey := ""
		guessedMessage := ""
		for _, chunk := range transposedChunks {
			guess := XorCypherBestGuess(chunk)

			guessedKey += string(guess.Key)
		}

		guessedMessage = string(RepeatingKeyXor(guessedKey, data))
		fmt.Printf("KEY %v \nProduced: \n%v\n", guessedKey, guessedMessage)

		if keysizeRank > 2 {
			break
		}

	}
}
