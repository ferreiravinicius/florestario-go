package entities_test

import (
	"pesthub/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisorder(t *testing.T) {

	assert.Equal(t, entities.CauserAnimal, "animal")
	assert.Equal(t, entities.CauserFungus, "fungus")
	assert.Equal(t, entities.CauserInsect, "insect")

	assert.Equal(t, entities.ConsequenceDamage, "damage")
	assert.Equal(t, entities.ConsequenceDisease, "disease")

	_ = entities.Disorder{
		Id:           uint64(1),
		Name:         "name",
		Description:  "description",
		Causer:       "causer",
		Consequences: []string{"cause.one"},
	}
}
