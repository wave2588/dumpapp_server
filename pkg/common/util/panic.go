package util

import (
	"io"

	"github.com/pkg/errors"
)

func PanicIf(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		e := errors.WithStack(err)
		panic(e)
	}
}
