package basics

import (
	"encoding/csv"
	"encoding/hex"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

func CalculateEnglishScore(message string) float32 {
	letterFreq := readFrequencyTable()

	freq := make(map[rune]float32)

	total := float32(0)

	message = strings.ToLower(message)

	for _, r := range message {

		if val, exists := freq[r]; !exists {
			freq[r] = 1
		} else {
			freq[r] = val + 1
		}
	}

	for r, val := range freq {
		if freq, exists := letterFreq[r]; !exists {
			if string(r) != " " {
				total -= float32(val / float32(len(message)))
			} else {
				total -= freq - (val / float32(len(message)))
			}
		}
	}

	return total

}

func readFrequencyTable() map[rune]float32 {
	// open file
	f, err := os.Open("data/english_letter_freq.csv")
	if err != nil {
		panic(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	freqMap := make(map[rune]float32)
	// convert records to array of structs
	for i, line := range data {
		if i > 0 { // omit header line
			r, _ := utf8.DecodeRuneInString(strings.ToLower(line[0]))
			freq, err := strconv.ParseFloat(line[1], 32)

			if err != nil {
				panic("Error parsing csv float Frequency.")
			}

			freqMap[r] = float32(freq)
		}
	}

	return freqMap
}

type Pair struct {
	Key   string
	Value float32
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func RankByLetterFreq(wordFrequencies map[string]float32) string {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))

	return pl[0].Key

	// for _, v := range pl {
	// 	fmt.Println(v.Key, v.Value)
	// }
}

func XorCypherBestGuessHex(encondedHex string) string {
	data, err := hex.DecodeString(encondedHex)

	if err != nil {
		panic("Error decoding hex string")
	}

	return XorCypherBestGuess(data)
}

func XorCypherBestGuess(encondedMessage []byte) string {
	decodedStrings := make(map[string]float32)

	for x := 0; x < int(math.Pow(2, 8)); x++ {
		message := SingleXor(encondedMessage, byte(x))

		// fmt.Println(message)

		decodedStrings[message] = CalculateEnglishScore(message)

	}

	return RankByLetterFreq(decodedStrings)
}

func SingleXor(data []byte, secret byte) string {

	result := ""

	for _, val := range data {
		result += string(val ^ secret)
	}

	return result

}
