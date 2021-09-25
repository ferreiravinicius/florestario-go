package validators

import (
	"errors"
	"pesthub/contracts"
	"pesthub/env"
)

const (
	MsgNameMinimumSize = "name.minimum.size"
	dfFieldName        = "Name"
)

func Name(name string, field ...string) error {
	f := append(field, dfFieldName)[0]
	n := sanitize(name)
	if len(n) < 3 {
		msg := env.MessageProvider.Get(MsgNameMinimumSize, contracts.ArgMap{fieldArg: f})
		return errors.New(msg)
	}
	return nil
}
