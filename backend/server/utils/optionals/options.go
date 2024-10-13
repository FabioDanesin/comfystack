package optional

type Optional[T any] struct {
	value *T
}

func (opt *Optional[T]) HasValue() bool {
	return opt != nil && opt.value != nil
}

func (opt *Optional[T]) TryGetValue() *T {
	if opt.HasValue() {
		return opt.value
	} else {
		return nil
	}
}

func BuildOptional[T any](value *T) *Optional[T] {
	newOpt := new(Optional[T])
	newOpt.value = value
	return newOpt
}
