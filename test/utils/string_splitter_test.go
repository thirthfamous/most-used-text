package test

import (
	"testing"
	"thirthfamous/golang-most-used-text/utils"

	"github.com/stretchr/testify/assert"
)

func TestStringSplitterSuccess(t *testing.T) {
	text := "one two three"
	result := utils.Split(text)
	assert.Equal(t, result[0].Word, "one")
}
