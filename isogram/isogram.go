//Package isogram exposes the IsIsogram function
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns a boolean indicating
// whether or not a string is an isogram (no repeating alpha characters)
func IsIsogram(w string) bool {
	w = strings.ToLower(w)
	for i, v := range w {
		if unicode.IsLetter(v) && strings.LastIndex(w, string(v)) != i {
			return false
		}
	}
	return true
}
