package entities

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
