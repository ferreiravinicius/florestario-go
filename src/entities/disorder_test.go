package entities_test

import (
	"pesthub/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisorder(t *testing.T) {

	fungus := entities.CauserFungus
	assert.True(t, fungus.Name == "fungus" && fungus.Code == 0)

	animal := entities.CauserAnimal
	assert.True(t, animal.Name == "animal" && animal.Code == 1)

	insect := entities.CauserInsect
	assert.True(t, insect.Name == "insect" && insect.Code == 2)

	_ = entities.Disorder{
		Id:          uint64(1),
		Name:        "name",
		Description: "description",
	}
}
