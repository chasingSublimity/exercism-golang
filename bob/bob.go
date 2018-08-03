// Package bob should provide the Hey function
package bob

import (
	"strings"
)

// Hey should take a remark string and conditionally return a sassy teenage response
func Hey(remark string) string {

	trimmed := strings.TrimSpace(remark)

	if trimmed == "" {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(trimmed, "?") {
		if isYelling(trimmed) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if isYelling(trimmed) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isYelling(remark string) bool {
	const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return remark == strings.ToUpper(remark) && strings.ContainsAny(remark, alpha)
}
