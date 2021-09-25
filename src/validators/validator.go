package validators

import "strings"

const fieldArg = "field"

func sanitize(value string) string {
	return strings.TrimSpace(value)
}
