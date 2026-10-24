package services

import (
	"fmt"
	"strings"
)

// if flag is set, append underscore
func AddSepIfValue(f string) string {
	if len(f) > 0 {
		return fmt.Sprintf("%s_", f)
	}
	return ""
}

// make sure no flags contain underscore
func ValidateFlagValues(p string, n string, e string) bool {
	if strings.Contains(p, "_") {
		return false
	}
	if strings.Contains(n, "_") {
		return false
	}
	if strings.Contains(e, "_") {
		return false
	}
	return true
}
