package testmsgs

import "pesthub/contracts"

type TestableMessageProvider struct {
}

func (TestableMessageProvider) Get(code string, _ ...contracts.ArgMap) string {
	return code
}

func NewTestableMessageProvider() contracts.MessageProvider {
	return TestableMessageProvider{}
}
