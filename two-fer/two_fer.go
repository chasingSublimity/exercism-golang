// Package twofer exposes the ShareWith function
package twofer

// ShareWith returns the string "One for X, one for me.",
// where X is a string passed in via the "name" argument.
// If "name" is an empty string, X will be set to the string "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
