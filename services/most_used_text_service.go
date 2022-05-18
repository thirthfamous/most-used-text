package services

import (
	"fmt"
	"thirthfamous/golang-most-used-text/models"
	"thirthfamous/golang-most-used-text/utils"
)

func TopTenMostUsedText(text string) []models.TopTenMostUsedWord {
	purgedText := utils.PurgeString(text)
	words := utils.Split(purgedText)
	result := utils.Sort(words)
	var topTen = make([]models.TopTenMostUsedWord, 10)

	fmt.Println("Top Ten Most Used Text")
	fmt.Println("================================")
	for i := 0; i < 10; i++ {
		topTen[i] = models.TopTenMostUsedWord{
			Word:  result[i].Word,
			Times: topTen[i].Times,
		}
		fmt.Print(i+1, ". ", result[i].Word)
		fmt.Println(" used ", result[i].Times, " times")
	}
	return topTen
}
