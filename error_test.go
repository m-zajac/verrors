package verrors

import (
	"testing"

	"github.com/pkg/errors"
)

type errKey string

const (
	temporaryErrKey errKey = "temporary"
	accessErrKey    errKey = "access"
)

func NewTemporaryError(e error) error {
	return WithValue(e, temporaryErrKey, true)
}

func IsTemporary(e error) bool {
	if v := Value(e, temporaryErrKey); v == nil {
		return false
	}

	return true
}

func NewAccessError(e error) error {
	return WithValue(e, accessErrKey, true)
}

func IsForbidden(e error) bool {
	if v := Value(e, accessErrKey); v == nil {
		return false
	}

	return true
}

func TestConstructorAndCheckerFunc(t *testing.T) {
	errMsg := "test error"
	base := errors.New(errMsg)

	if IsForbidden(base) {
		t.Error("IsForbidden should return false on base error")
	}
	accErr := NewAccessError(base)
	if accErr.Error() != errMsg {
		t.Error("error message changed by NewAccessError")
	}
	if !IsForbidden(accErr) {
		t.Error("IsForbidden should return true for access error")
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
	if !IsForbidden(wrappedErr) {
		t.Error("IsForbidden returned invalid result on wrapped access error")
	}

	wrappedErr = errors.Wrap(
		NewTemporaryError(wrappedErr),
		"w3",
	)
	if !IsForbidden(wrappedErr) {
		t.Error("invalid Forbidden result on double wrapped access error")
	}
	if !IsTemporary(wrappedErr) {
		t.Error("invalid IsTemporary result on multiple wrapped temporary error")
	}
}
