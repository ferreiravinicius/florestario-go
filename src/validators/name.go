package validators

import (
	"pesthub/contracts"
	"pesthub/env"
	"pesthub/failures"
)

const (
	MsgNameMinimumSize = "name.minimum.size"
	dfFieldName        = "Name"
)

func Name(name string, field ...string) error {
	f := append(field, dfFieldName)[0] //first or df
	n := sanitize(name)
	if len(n) < 3 {
		msg := env.MessageProvider.Get(MsgNameMinimumSize, contracts.ArgMap{fieldArg: f})
		return failures.Validation(msg, f)
	}
	return nil
}
