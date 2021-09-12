package text

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func Normalize(text string) string {
	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		runes.Remove(runes.In(unicode.Po)),
		norm.NFC,
	)
	result, _, _ := transform.String(t, text)
	return result
}

func Slugfy(text string) string {
	words := strings.Split(text, " ")
	slug := make([]string, 0, len(words))
	for _, word := range words {
		normalized := Normalize(word)
		lowered := strings.ToLower(normalized)
		slug = append(slug, lowered)
	}
	return strings.Join(slug, "-")
}
