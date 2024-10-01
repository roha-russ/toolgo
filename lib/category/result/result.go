package result

type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Value
}

func (r Result[T]) UnwrapOr(def T) T {
	if r.Err != nil {
		return def
	}
	return r.Value
}

func (r Result[T]) UnwrapOrNil() *T {
	if r.Err != nil {
		return nil
	}
	return &r.Value
}

func (r Result[T]) UnwrapErr() error {
	return r.Err
}

func (r Result[T]) UnwrapErrOr(def error) error {
	if r.Err != nil {
		return r.Err
	}
	return def
}

func (r Result[T]) UnwrapErrOrNil() *error {
	if r.Err != nil {
		return &r.Err
	}
	return nil
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}
