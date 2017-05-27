# verrors [![Build Status](https://travis-ci.org/m-zajac/verrors.svg?branch=master)](https://travis-ci.org/m-zajac/verrors) [![Go Report Card](https://goreportcard.com/badge/github.com/m-zajac/verrors)](https://goreportcard.com/report/github.com/m-zajac/verrors) [![GoDoc](https://godoc.org/github.com/m-zajac/verrors?status.svg)](http://godoc.org/github.com/m-zajac/verrors)

Package verrors provides simple helpers for creating errors, which can be checked for their behaviour.

Package exposes only 2 functions: "WithValue" and "Value", which helps creating error constructors and behaviour checks.

This package is created as extension of github.com/pkg/errors, but pkg/errors is not required. You can wrap errors by errors.Wrap and behaviour checks will still work.

Example:

```go
import (
    "github.com/pkg/errors"
    "github.com/m-zajac/verrors"
)

// type for error keys
type errKey string

// key for "temporary" error
const temporaryErrKey errKey = "temporary"

// NewTemporaryError creates temporary error
func NewTemporaryError(e error) error {
    return verrors.WithValue(e, temporaryErrKey, true)
}

// IsTemporary checks if error has "temporary" nature
func IsTemporary(e error) bool {
    if v := verrors.Value(e, temporaryErrKey); v == nil {
	   return false
    }
    return true
}

// ...

err := NewTemporaryError(errors.New("tmp error"))

// Fake some wraps. In real code wraps would happen while moving up the call stack.
err = errors.Wrap(err, "some context 1")
err = errors.Wrap(err, "some context 2")

// ...

if IsTemporary(err) {
    println("Should retry...") // error cause is created by NewTemporaryError, so this is the case
} else {
    panic("I'll end here.")
}
```
