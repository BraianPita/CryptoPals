package block

// Adds PKCS#7 Padding to an array of bytes based on the block size
func AddPkcs7Padding(data []byte, blockSize uint8) []byte {

	// calculate needed padding
	padding := int(blockSize) - (len(data) % int(blockSize))

	// fmt.Printf("%v %v %v", len(data), blockSize, padding)

	// if no padding, the padding would be a complete new block
	if padding == 0 {
		padding = int(blockSize)
	}

	// Create padding data (slice)
	paddingData := make([]byte, padding)
	for i := range paddingData {
		paddingData[i] = byte(padding)
	}

	return append(data, paddingData...)

}
