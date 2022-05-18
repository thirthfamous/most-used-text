package test

import (
	"testing"
	"thirthfamous/golang-most-used-text/services"

	"github.com/stretchr/testify/assert"
)

func TestMostUsedTextSuccess(t *testing.T) {
	text := "The Witcher 3: Wild Hunt was released for Microsoft Windows, PlayStation 4, and Xbox One in May 2015, with a Nintendo Switch version released in October 2019, and PlayStation 5 and Xbox Series X/S versions planned for release in 2022. The game received critical acclaim, with praise for its gameplay, narrative, world design, combat, and visuals, although it received minor criticism due to technical issues. It received numerous game of the year awards and has been cited as one of the best video games ever made. Two expansions were also released to critical acclaim: Hearts of Stone and Blood and Wine. A Game of the Year edition was released in August 2016, with the base game, expansions and all downloadable content. The game shipped over 40 million copies, making it one of the best selling video games of all time."
	mostUsedText := services.TopTenMostUsedText(text)
	assert.Equal(t, mostUsedText[0].Word, "game")
}

func TestMostUsedTextFail(t *testing.T) {
	text := ""
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	services.TopTenMostUsedText(text)
}
