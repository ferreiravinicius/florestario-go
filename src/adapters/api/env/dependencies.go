package env

import "pesthub/contracts"

type ApiDependencies struct {
	DisorderStore   contracts.DisorderStore
	MessageProvider contracts.MessageProvider
}

var Deps *ApiDependencies
