// Package raindrops exposes the Convert function
package raindrops

import "strconv"

// Convert takes in an integer and conditionally returns a string
// based on the factors of said integer
func Convert(n int) string {
	var reply string

	if n%3 == 0 {
		reply += "Pling"
	}
	if n%5 == 0 {
		reply += "Plang"
	}
	if n%7 == 0 {
		reply += "Plong"
	}

	if reply == "" {
		reply = strconv.Itoa(n)
	}
	return reply
}
