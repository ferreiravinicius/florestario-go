package text

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Normalize text
// Eg. "máçÀ" => "macA"
func Normalize(text string) string {
	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		runes.Remove(runes.In(unicode.Po)), //TODO: export this to new fn
		norm.NFC,
	)
	result, _, _ := transform.String(t, text)
	return result
}

// Normalize, lowerize and generate slug based on the text
// Eg. "ábç De?f" => "abc-def"
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
