package uuid

const defaultVersion = Version7

type Version uint8

const (
	Version4 Version = 4
	Version7 Version = 7
	Version8 Version = 8
)

var currentVersion = defaultVersion

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
