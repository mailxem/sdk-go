package xem

import "fmt"

const (
	// Major version number
	VersionMajor = 0
	// Minor version number
	VersionMinor = 1
	// Patch version number
	VersionPatch = 0
)

// Version returns the full version string of the SDK
func Version() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}
