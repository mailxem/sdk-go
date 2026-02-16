package xem

import "fmt"

const (
	// Major version number
	VersionMajor = 0
	// Minor version number
	VersionMinor = 1
	// Patch version number
	VersionPatch = 1
)

var Version = fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
