package entities

import (
	"errors"
	"pesthub/commons/text"
	"strings"
)

type Culture struct {
}

type Symptom struct {
}

type ControlMethod struct {
}

type Damage struct {
}

type Pest struct {
	// // Unique
	Name         string
	BinomialName string
	Slug         string
	// Kind           string //insect, fungae
	// Description    string
	// PopularNames   []string
	// // Duplicated NxN
	// // Damages        []Damage ignore for now
	// CommonCultures []Culture
	// ControlMethods []ControlMethod
	// Symptoms       []Symptom
}

func (pest *Pest) FillSlug() error {
	name := pest.Name
	if len(name) == 0 {
		return errors.New("name is required to generate slug")
	}

	trimmed := strings.TrimSpace(name)
	words := strings.Split(trimmed, " ")
	result := make([]string, len(words))
	for _, word := range words {
		normalized := text.Normalize(word)
		lowered := strings.ToLower(normalized)
		result = append(result, lowered)
	}
}
