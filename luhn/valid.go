package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid returns a boolean indicating whether or not a
// a string passes the luhn algo
func Valid(s string) bool {
	var valid bool

	s = tokenize(s)
	if len(s) <= 1 {
		return valid
	}

	var sum int
	var shouldDouble bool

	for idx := (len(s) - 1); idx >= 0; idx-- {
		if !unicode.IsNumber(rune(s[idx])) {
			return valid
		}

		var acc int
		v, _ := strconv.Atoi(string((s[idx])))

		if shouldDouble == true {
			acc += v * 2
		} else {
			acc += v
		}
		if acc > 9 {
			acc -= 9
		}
		sum += acc
		shouldDouble = !shouldDouble
	}

	if sum%10 == 0 {
		valid = true
	}

	return valid
}

// filter out spaces
func tokenize(s string) string {
	r := strings.NewReplacer(" ", "")
	return r.Replace(s)
}
