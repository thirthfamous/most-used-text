package utils

import (
	"strings"
	"thirthfamous/golang-most-used-text/models"
)

func Split(text string) []models.TopTenMostUsedWord {

	words := strings.Fields(text)
	var sliceOfWords = make([]models.TopTenMostUsedWord, len(words))

	for i := 0; i < len(words); i++ { // loop of words
		words[i] = strings.ToLower(words[i])
		if !IsCommonWord(words[i]) {
			for j := 0; i < len(words); j++ { // loop of toptenwords
				// go to an index, if null then assign to it
				if len(sliceOfWords[j].Word) == 0 {
					sliceOfWords[j] = models.TopTenMostUsedWord{
						Word:  words[i],
						Times: sliceOfWords[j].Times + 1,
					}
					break
				} else if sliceOfWords[j].Word == words[i] {
					sliceOfWords[j].Times += 1
					break
				}
				// if also no and not same, go for next array
			}
		}
	}
	return sliceOfWords
}
