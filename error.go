package verrors

import "context"

func WithValue(e error, k interface{}, v interface{}) error {
	if err, ok := e.(ctxError); ok {
		return err.withValue(k, v)
	}

	return ctxError{
		err: e,
		ctx: context.WithValue(context.Background(), k, v),
	}
}

func Value(e error, k interface{}) interface{} {
	for {
		if err, ok := e.(ctxError); ok {
			if v := err.ctx.Value(k); v != nil {
				return v
			}
		}
		if c, ok := e.(causer); ok {
			e = c.Cause()
			continue
		}

		return nil
	}
}

type causer interface {
	Cause() error
}

type ctxError struct {
	err error
	ctx context.Context
}

func (e ctxError) Error() string {
	return e.err.Error()
}

func (e ctxError) Cause() error {
	return e.err
}

func (e ctxError) withValue(k interface{}, v interface{}) ctxError {
	e.ctx = context.WithValue(e.ctx, k, v)
	return e
}

func (e ctxError) get(k interface{}) interface{} {
	return e.ctx.Value(k)
}
