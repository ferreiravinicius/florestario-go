package entities

const (
	CauserAnimal = "animal"
	CauserFungus = "fungus"
	CauserInsect = "insect"
)

const (
	ConsequenceDamage  = "damage"
	ConsequenceDisease = "disease"
)

type Disorder struct {
	Id           uint64
	Name         string
	Causer       string
	Description  string
	Consequences []string
}

// BinomialName [?]string
// Slug   string
// CommonCultures []Culture
// ControlMethods []ControlMethod
// Symptoms       []Symptom
