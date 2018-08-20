//Package isogram exposes the IsIsogram function
package isogram

import (
	"regexp"
	"sort"
	"strings"
)

// IsIsogram returns a boolean indicating
// whether or not a string is an isogram (no repeating alpha characters)
func IsIsogram(w string) bool {
	token := tokenize(w)
	ss := sort.StringSlice(strings.Split(token, ""))
	ss.Sort()

	for i, v := range ss {
		if i > 0 {
			if v == ss[i-1] {
				return false
			}
		}
	}
	return true
}

func tokenize(s string) string {
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-z]+")
	return reg.ReplaceAllString(s, "")
}
