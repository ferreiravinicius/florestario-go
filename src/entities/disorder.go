package entities

// type Culture struct {
// }

// type Symptom struct {
// }

// type ControlMethod struct {
// }

// type Damage struct {
// }

type Disorder struct {
	// // Unique
	Code uint64
	Name string

	// BinomialName string //maybe list ?
	Slug   string
	Kind   string //desease, pest
	Causer string //insect, fungae
	// Description    string
	// // Damages        []Damage ignore for now
	// CommonCultures []Culture
	// ControlMethods []ControlMethod
	// Symptoms       []Symptom
}
