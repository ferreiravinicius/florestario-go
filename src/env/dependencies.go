package env

import (
	"pesthub/contracts"
)

var DisorderStore contracts.DisorderStore
var Messages contracts.Messages

func CheckDependencies() {
	deps := []interface{}{
		DisorderStore,
		Messages,
	}

	for _, dep := range deps {
		check(dep)
	}
}

func check(dep interface{}) {
	if dep == nil {
		panic("All dependencies must be initialized")
	}
}
