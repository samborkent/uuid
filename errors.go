package uuid

import "errors"

var (
	ErrCompareVersions = errors.New("cannot compare different uuid versions")
	ErrInvalidVersion  = errors.New("version not supported. must be 4, 7, or 8")
	ErrNotTimeOrdered  = errors.New("uuid version is not time ordered")
)
