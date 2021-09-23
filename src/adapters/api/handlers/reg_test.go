package handlers

import (
	"pesthub/adapters/api/env"
	"pesthub/adapters/testmsgs"
	"pesthub/entities"
	"pesthub/usecases/disorder"
	"testing"
)

type mock struct{}

func (*mock) Save(_ *entities.Disorder) (*entities.Disorder, error) {
	return &entities.Disorder{
		Code: 123,
	}, nil
}
func (*mock) ExistsName(_ string) (bool, error) {
	return false, nil
}

func (*mock) FindAll() ([]entities.Disorder, error) {
	return nil, nil
}

func getInput() *disorder.RegisterDisorderInput {
	return &disorder.RegisterDisorderInput{}
}

func Direct() *disorder.RegisterDisorderOutput {
	usecase := disorder.NewRegisterDisorder(env.Deps.DisorderStore, env.Deps.Messages)
	output, _ := usecase.Execute(getInput())
	return output
}

func Dep(usecase disorder.RegisterDisorderUseCase) *disorder.RegisterDisorderOutput {
	output, _ := usecase.Execute(getInput())
	return output
}

func resetEnv() {
	store := &mock{}
	messages := testmsgs.NewTestableMessages()

	env.Deps = &env.ApiDependencies{
		DisorderStore: store,
		Messages:      messages,
	}
}

var resultDirect *disorder.RegisterDisorderOutput

func BenchmarkDirect(b *testing.B) {

	resetEnv()

	var r *disorder.RegisterDisorderOutput
	for i := 0; i < b.N; i++ {
		r = Direct()
	}
	resultDirect = r
}

var resultDep *disorder.RegisterDisorderOutput

func BenchmarkCached(b *testing.B) {

	resetEnv()
	cached := disorder.NewRegisterDisorder(env.Deps.DisorderStore, env.Deps.Messages)

	var r *disorder.RegisterDisorderOutput
	for i := 0; i < b.N; i++ {
		r = Dep(cached)
	}
	resultDep = r
}

var resultDepNoCache *disorder.RegisterDisorderOutput

func BenchmarkOnTheFly(b *testing.B) {
	resetEnv()
	var r *disorder.RegisterDisorderOutput
	for i := 0; i < b.N; i++ {
		usecase := disorder.NewRegisterDisorder(env.Deps.DisorderStore, env.Deps.Messages)
		r = Dep(usecase)
	}
	resultDepNoCache = r
}
