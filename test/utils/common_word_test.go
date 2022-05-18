package test

import (
	"testing"
	"thirthfamous/golang-most-used-text/utils"

	"github.com/stretchr/testify/assert"
)

func TestCommonWordSuccess(t *testing.T) {
	text := "is"
	mustTrue := utils.IsCommonWord(text)
	assert.Equal(t, mustTrue, true)
}
