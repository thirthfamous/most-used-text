package utils

import "thirthfamous/golang-most-used-text/models"

func Sort(topTen []models.TopTenMostUsedWord) []models.TopTenMostUsedWord {
	for i := 0; i < len(topTen)-1; i++ {
		for j := 0; j < len(topTen)-i-1; j++ {
			if topTen[j].Times < topTen[j+1].Times {
				topTen[j], topTen[j+1] = topTen[j+1], topTen[j]
			}
		}
	}
	return topTen
}
