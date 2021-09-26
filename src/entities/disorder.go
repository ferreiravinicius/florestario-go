package entities

import "errors"

var (
	CauserFungus = causer{Code: 0, Name: "fungus"}
	CauserAnimal = causer{Code: 1, Name: "animal"}
	CauserInsect = causer{Code: 2, Name: "insect"}
)

type Disorder struct {
	Id          uint64
	Name        string
	Causer      causer
	Description string
}

// Private struct that holds @causer attributes.
type causer struct {
	Code int
	Name string
}

// Prints the name of the causer.
func (c causer) String() string {
	return c.Name
}

// Parse/get @causer from given string.
// Returns error if no @causer matches the string.
func ParseCauser(name string) (causer, error) {
	switch name {
	case CauserAnimal.Name:
		return CauserAnimal, nil
	case CauserFungus.Name:
		return CauserFungus, nil
	case CauserInsect.Name:
		return CauserInsect, nil
	default:
		return causer{}, errors.New("no causer matches given string")
	}
}

// BinomialName [?]string
// Slug   string
// CommonCultures []Culture
// ControlMethods []ControlMethod
// Symptoms       []Symptom
// Consequences []string //disesae? make bool?
