package profiler

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Lunarisnia/n-gram/internal/strgram"
	"github.com/Lunarisnia/n-gram/internal/tokenizer"
)

func CreateProfile(datasetPath string) (*HashMap, error) {
	fmt.Println("Processing: ", datasetPath)
	fmt.Print("Reading Files: ")
	f, err := os.Open(datasetPath)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	dataset := string(b)
	fmt.Println("COMPLETE")

	profile, err := processDataset(dataset)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func CreateProfileUsingRawData(dataset string) (*HashMap, error) {
	fmt.Println("Processing: ", "Raw Dataset")
	profile, err := processDataset(dataset)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func processDataset(dataset string) (*HashMap, error) {
	fmt.Print("Tokenizing: ")
	sanitized := tokenizer.Sanitize(dataset)
	tokens, err := tokenizer.Tokenize(sanitized)
	if err != nil {
		return nil, err
	}
	ngrams := make([]string, 0)
	for _, token := range tokens {
		for i := range 5 {
			ngram, err := strgram.Encode(strings.Join(token, ""), i+1)
			if err != nil {
				return nil, err
			}
			// FIXME: I don't think I can mash all ngram together, it has to be kept separately
			// NOTE: From what I think I understand I seems that this ngram has to be matched by the corresponding same size ngram
			// NOTE: and from those 5 ngram the lowest are taken which mean each language has 5 chance to be the lowest
			// NOTE: I don't think this matters
			for _, gram := range ngram {
				ngrams = append(ngrams, strings.Join(gram, ""))
			}
		}
	}
	fmt.Println("COMPLETE")

	fmt.Print("Profiling: ")
	profile := NewHashMap()
	for _, ngram := range ngrams {
		profile.InsertNode(ngram)
	}
	fmt.Println("COMPLETE")

	fmt.Print("CALCULATING RANKS: ")
	profile.CalculateAllRanks()
	fmt.Println("COMPLETE")

	return profile, nil
}
