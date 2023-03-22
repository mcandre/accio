//go:build !windows
// +build !windows

package accio

// Delve applies OS-specific binary file extensions.
func Delve(executable string) string {
	return executable
}
