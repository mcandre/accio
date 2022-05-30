//go:build windows

package accio

import (
	"fmt"
	"strings"
)

// Delve applies OS-specific binary file extensions.
func Delve(executable string) string {
	e := strings.ToLower(executable)

	if strings.HasSuffix(e, ".exe") {
		return e
	}

	return fmt.Sprintf("%s.exe", e)
}
