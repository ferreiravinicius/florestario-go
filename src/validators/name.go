package validators

import (
	"errors"
	"pesthub/contracts"
)

const (
	MsgNameMinimumSize = "name.minimum.size"
	dfFieldName        = "Name"
)

func Name(messages contracts.Messages, name string, field ...string) error {
	f := append(field, dfFieldName)[0]
	n := sanitize(name)
	if len(n) < 3 {
		msg := messages.GetText(MsgNameMinimumSize, contracts.ArgMap{fieldArg: f})
		return errors.New(msg)
	}
	return nil
}
