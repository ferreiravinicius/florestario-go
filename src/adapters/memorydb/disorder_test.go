package memorydb_test

import (
	"pesthub/adapters/memorydb"
	"pesthub/contracts/store"
	"pesthub/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sut() store.DisorderStore {
	s := memorydb.NewMemoryDisorderStore()
	return s
}

func TestSave(t *testing.T) {

	t.Run("it should generate unique code", func(t *testing.T) {
		sut := sut()
		disorder := entities.Disorder{}
		saved, err := sut.Save(&disorder)
		assert.NoError(t, err)
		assert.Greater(t, saved.Code, uint64(0))
	})

}

func TestFindAll(t *testing.T) {
	sut := sut()
	d1 := &entities.Disorder{Name: "name"}
	d2 := &entities.Disorder{Name: "name"}
	sut.Save(d1)
	sut.Save(d2)
	disorders, err := sut.FindAll()
	assert.NoError(t, err)
	assert.Len(t, disorders, 2)
}

func TestExistsName(t *testing.T) {
	sut := sut()
	sut.Save(&entities.Disorder{Name: "name"})
	exists, err := sut.ExistsName("name")
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = sut.ExistsName("name that does not exists")
	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestDirectAccessShouldNotBeAllowed(t *testing.T) {

	sut := sut()
	d1 := &entities.Disorder{Name: "name"}
	d2 := &entities.Disorder{Name: "name"}
	saved1, _ := sut.Save(d1)
	saved2, _ := sut.Save(d2)

	// find all then change all
	all1, _ := sut.FindAll()
	for _, d := range all1 {
		d.Name = "changed"
	}

	all2, _ := sut.FindAll()
	for _, d := range all2 {
		assert.Equal(t, "name", d.Name)
	}

	// change original and saved
	saved1.Name = "oops"
	saved2.Name = "oops"
	d1.Name = "oops"
	d2.Name = "oops"

	all3, _ := sut.FindAll()
	for _, d := range all3 {
		assert.Equal(t, "name", d.Name)
	}

}
