package store

import "errors"

var (
	ErrBadSrcType         = errors.New("cache store: src must be a slice or pointer-to-slice")
	ErrKeysLengthNotMatch = errors.New("cache store: keys and src slices have different length")
	ErrBadDstType         = errors.New("cache store: dst must be a pointer")
	ErrBadDstMapType      = errors.New("cache store: dst must be a map or pointer-to-map")
	ErrBadDstMapValue     = errors.New("cache store: dst must not be a nil map")
	ErrSrcDstTypeMismatch = errors.New("cache store: type of fallback result mismatch error")
)

type Store interface{}
