package test

import (
	"testing"
	"thirthfamous/golang-most-used-text/utils"

	"github.com/stretchr/testify/assert"
)

func TestFormatTextSuccess(t *testing.T) {
	text := "Hello,..\",.,."
	purgedText := utils.PurgeString(text)
	assert.Equal(t, "Hello", purgedText)
}
