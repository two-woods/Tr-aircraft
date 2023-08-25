package cerror

import "github.com/pkg/errors"

func WErr(err error, message string) {

}

func WErrF(err error, message string, params ...interface{}) {

}

func WithMessage(err error, message string) {

}

func WithMessagef(err error, message string, params ...interface{}) {

}

func Cause(err error) error {
	return errors.Cause(err)
}
