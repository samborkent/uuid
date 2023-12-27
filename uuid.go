package uuid

import "errors"

type UUID [16]byte

type Version uint8

const defaultVersion = Version7

const (
	Version4 Version = 4
	Version7 Version = 7
	Version8 Version = 8
)

var currentVersion = defaultVersion

var ErrInvalidVersion = errors.New("version not supported. must be 4, 7, or 8")

// SetVersion sets the version of the application.
//
// It takes an integer value representing the version and returns an error.
func SetVersion(version int) error {
	newVersion := Version(version)

	if newVersion != Version4 && newVersion != Version7 && newVersion != Version8 {
		return ErrInvalidVersion
	}

	currentVersion = newVersion

	return nil
}

func New() UUID {
	switch currentVersion {
	case Version4:
		return NewV4()
	case Version7:
		return NewV7()
	case Version8:
		return NewV8()
	default:
		return NewV7()
	}
}

func NewV4() UUID {
	uuid := make([]byte, 16)

	return UUID(uuid)
}

func NewV7() UUID {
	return UUID{}
}

func NewV8() UUID {
	return UUID{}
}
