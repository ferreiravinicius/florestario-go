package contracts

type Messages interface {
	GetText(code string, args ...map[string]string) string
}
