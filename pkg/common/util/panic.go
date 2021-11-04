package util

import (
	"github.com/pkg/errors"
)

func PanicIf(err error) {
	if err != nil {
		e := errors.WithStack(err)
		panic(e)
	}
}
