package utils

var commonWords = []string{
	"yes",
	"no",
	"the",
	"was",
	"for",
	"and",
	"in",
	"with",
	"a",
	"in",
	"it",
	"of",
	"has",
	"been",
	"as",
	"is",
}

func IsCommonWord(word string) bool {

	for i := 0; i < len(commonWords); i++ {
		if commonWords[i] == word {
			return true
		}
	}
	return false
}
