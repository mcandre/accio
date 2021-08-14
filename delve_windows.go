// +build windows

package accio

import (
	"fmt"
)

// Delve applies OS-specific binary file extensions.
func Delve(executable string) string {
	e = strings.ToLower(executable)

	if strings.EndsWith(e, ".exe") {
		return e
	}

	return fmt.Sprintf("%s.exe", e)
}
