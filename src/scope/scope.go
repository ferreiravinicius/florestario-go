//TODO: rename files and separate them like disorder.go, message.go and etc
package scope

import (
	"pesthub/contracts"
	"pesthub/contracts/store"
)

// Disorder
var DisorderStore store.DisorderStore

// Message
var Messages contracts.Messages
