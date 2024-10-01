package option

type Option[T any] struct {
	Value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{Value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) Unwrap() T {
	if o.Value == nil {
		panic("unwrap of None")
	}
	return *o.Value
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.Value == nil {
		return def
	}
	return *o.Value
}

func (o Option[T]) UnwrapOrNil() *T {
	if o.Value == nil {
		return nil
	}
	return o.Value
}

func (o Option[T]) IsSome() bool {
	return o.Value != nil
}

func (o Option[T]) IsNone() bool {
	return o.Value == nil
}
