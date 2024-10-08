package common

import (
	"github.com/pkg/errors"
)

func PanicOnError(err error, msg string) {
	if err != nil {
		panic(errors.WithMessage(err, msg))
	}
}
