package contracts

type ArgMap map[string]string

type Messages interface {
	GetText(code string, args ...ArgMap) string
}
