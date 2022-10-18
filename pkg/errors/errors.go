package errors

import "github.com/pkg/errors"

func New(text string) error {
	return errors.New(text)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func Wrap(err error, text string) error {
	return errors.Wrap(err, text)
}
