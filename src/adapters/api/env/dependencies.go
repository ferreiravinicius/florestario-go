package env

import "pesthub/contracts"

type ApiDependencies struct {
	DisorderStore contracts.DisorderStore
	Messages      contracts.Messages
}

var Deps *ApiDependencies
