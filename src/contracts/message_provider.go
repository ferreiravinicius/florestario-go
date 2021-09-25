package contracts

type ArgMap map[string]string

type MessageProvider interface {
	Get(code string, args ...ArgMap) string
}
