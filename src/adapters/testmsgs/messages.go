package testmsgs

import "pesthub/contracts"

type TestableMessages struct {
}

func (TestableMessages) GetText(code string, _ ...map[string]string) string {
	return code
}

func NewTestableMessages() contracts.Messages {
	return &TestableMessages{}
}
