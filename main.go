package main

import (
	"fmt"
	"math"

	"github.com/Lunarisnia/n-gram/internal/profiler"
)

func main() {
	indonesianProfile, err := profiler.CreateProfile("./tests/indonesian_dataset")
	if err != nil {
		panic(err)
	}
	englishProfile, err := profiler.CreateProfile("./tests/english_dataset")
	if err != nil {
		panic(err)
	}

	test := "Climate change poses a significant threat to global food security"
	inputProfile, err := profiler.CreateProfileUsingRawData(test)
	if err != nil {
		panic(err)
	}

	indonesianResult := 0
	englishResult := 0
	for _, testToken := range inputProfile.Data {
		indFound := false
		for _, token := range indonesianProfile.Data {
			if token.Token == testToken.Token {
				indFound = true
				indonesianResult += int(math.Abs(float64(token.Rank - testToken.Rank)))
			}
		}
		if !indFound {
			indonesianResult += 999_999
		}

		engFound := false
		for _, token := range englishProfile.Data {
			if token.Token == testToken.Token {
				engFound = true
				englishResult += int(math.Abs(float64(token.Rank - testToken.Rank)))
			}
		}
		if !engFound {
			englishResult += 999_999
		}
	}

	fmt.Println("Token: ", test)
	fmt.Println("English Score: ", englishResult)
	fmt.Println("Indonesian Score: ", indonesianResult)
}
