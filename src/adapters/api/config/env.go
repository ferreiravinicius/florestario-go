package config

import (
	"log"
	"pesthub/contracts"
)

type ApiEnv struct {
	DisorderStore contracts.DisorderStore
	Messages      contracts.Messages
}

var instance *ApiEnv

func SetEnv(deps *ApiEnv) {
	instance = deps
}

func Env() *ApiEnv {
	if instance == nil {
		log.Fatal("env must be initialized @ api/config/env")
	}
	return instance
}
