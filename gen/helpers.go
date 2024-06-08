package gen

import (
	"fmt"
	"strings"
)

// IsOneOf returns true if the input string is one of the options.
// func isOneOf(input string, options []string) bool {
// 	for _, o := range options {
// 		if o == input {
// 			return true
// 		}
// 	}

// 	return false
// }

// removeAll removes all occurrences of items from input.
// func removeAll(input string, items ...string) string {
// 	out := input
// 	for _, item := range items {
// 		out = strings.Replace(out, item, "", -1)
// 	}

// 	return out
// }

func wrapInValToRef(input string) string {
	return fmt.Sprintf("utils.ValToRef(%s)", input)
}

func wrapInRefToVal(input string) string {
	return fmt.Sprintf("utils.RefToVal(%s)", input)
}

func getFQObjectName(outPackage string, objectPackage string, name string) string {
	if strings.EqualFold(objectPackage, outPackage) {
		return name
	}

	return fmt.Sprintf("%s.%s", objectPackage, name)
}
