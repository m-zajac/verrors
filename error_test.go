package verrors

import (
	"testing"

	"github.com/pkg/errors"
)

type errKey string

const (
	invalidRequestErrKey errKey = "access"
	accessErrKey         errKey = "access"
)

func NewInvalidRequestError(e error) error {
	return WithValue(e, invalidRequestErrKey, true)
}

func InvalidRequest(e error) bool {
	if v := Value(e, invalidRequestErrKey); v == nil {
		return false
	}

	return true
}

func NewAccessError(e error) error {
	return WithValue(e, accessErrKey, true)
}

func Forbidden(e error) bool {
	if v := Value(e, accessErrKey); v == nil {
		return false
	}

	return true
}

func TestConstructorAndCheckerFunc(t *testing.T) {
	errMsg := "test error"
	base := errors.New(errMsg)

	if Forbidden(base) {
		t.Error("base error is not access error")
	}
	accErr := NewAccessError(base)
	if accErr.Error() != errMsg {
		t.Error("error message changed by NewAccessError")
	}
	if !Forbidden(accErr) {
		t.Error("invalid Forbidden result on access error")
	}
}

func TestWrapping(t *testing.T) {
	errMsg := "test error"
	base := errors.New(errMsg)

	wrappedErr := errors.Wrap(
		NewAccessError(
			errors.Wrap(base, "w1"),
		),
		"w2",
	)
	if !Forbidden(wrappedErr) {
		t.Error("Forbidden returned invalid result on wrapped access error")
	}

	wrappedErr = errors.Wrap(
		NewInvalidRequestError(wrappedErr),
		"w3",
	)
	if !Forbidden(wrappedErr) {
		t.Error("invalid Forbidden result on double wrapped access error")
	}
	if !InvalidRequest(wrappedErr) {
		t.Error("invalid InvalidRequest result on multiple wrapped inv. req. error")
	}
}
