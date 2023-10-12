package basics

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func calculateEnglishScore(message string) float64 {
	letterFreq := readFrequencyTable()

	freq := make(map[rune]float64)

	total := float64(0)

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
				total -= float64(val / float64(len(message)))
			} else {
				total -= freq - (val / float64(len(message)))
			}
		}
	}

	return total

}

func readFrequencyTable() map[rune]float64 {
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

	freqMap := make(map[rune]float64)
	// convert records to array of structs
	for i, line := range data {
		if i > 0 { // omit header line
			r, _ := utf8.DecodeRuneInString(strings.ToLower(line[0]))
			freq, err := strconv.ParseFloat(line[1], 32)

			if err != nil {
				panic("Error parsing csv float Frequency.")
			}

			freqMap[r] = float64(freq)
		}
	}

	return freqMap
}

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func rankByWordFreq(wordFrequencies map[string]float64) {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)

	for _, v := range pl {
		fmt.Println(v.Key, v.Value)
	}
}

func TestSingleXorCypher(t *testing.T) {
	secretMessage := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decodedStrings := make(map[string]float64)

	for x := 0; x < int(math.Pow(2, 8)); x++ {
		message := SingleXor(secretMessage, byte(x))

		// fmt.Println(message)

		decodedStrings[message] = calculateEnglishScore(message)

	}

	rankByWordFreq(decodedStrings)

}
