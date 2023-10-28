package basics

// Pass in an array of bytes, returns the
// number of repeating blocks of size bs,
// which can help detect AES in ECB mode
func DetectAesEcb(bs int, data []byte) int {
	blockCount := make(map[string]int)

	for i := 0; i < len(data); i += bs {
		block := EncodeBase64(data[i : i+bs])

		// if it does not exits, create entry,
		// if it exists, add one to it
		if _, ok := blockCount[block]; !ok {
			blockCount[block] = 0
		} else {
			blockCount[block] += 1
		}

	}

	// sum all repeated blocks
	sum := 0
	for _, val := range blockCount {
		sum += val
	}

	return sum

}

type AesEcbDetection struct {
	data           string
	repeatedBlocks int
}

func FindEncryptedAesEcb(options []string, blockSize int) AesEcbDetection {
	maxReps := AesEcbDetection{data: "", repeatedBlocks: 0}

	for _, cipher := range options {
		data := Base64ToBytes([]string{cipher})
		currReps := DetectAesEcb(blockSize, data)

		if maxReps.repeatedBlocks < currReps {
			maxReps = AesEcbDetection{
				data:           cipher,
				repeatedBlocks: currReps,
			}
		}
	}

	return maxReps
}
