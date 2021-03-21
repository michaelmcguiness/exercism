// Package twofer provides methods for the exercism.io Two Fer challenge.
package twofer

import "fmt"

// ShareWith returns a message that varies with the name parameter.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
