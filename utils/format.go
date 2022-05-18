package utils

import "strings"

func PurgeString(text string) string {
	// format the comma
	text = strings.Replace(text, ",", "", -1)
	// format the dot
	text = strings.Replace(text, ".", "", -1)
	// format the double quotes
	text = strings.Replace(text, "\"", "", -1)
	// format the double colon
	text = strings.Replace(text, ":", "", -1)

	return text

}
