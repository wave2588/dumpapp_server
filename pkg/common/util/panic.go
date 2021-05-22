package util

import (
	"github.com/pkg/errors"
)

func PanicIf(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}
