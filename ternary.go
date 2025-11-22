package ternary

import "github.com/zodimo/go-lazy"

type TernaryFunc[T any] func(condition bool, value1 T, value2 T) T

func Ternary[T any](condition bool, value1 T, value2 T) T {
	if condition {
		return value1
	}
	return value2
}

type TernaryLazyFunc[T any] func(condition lazy.Value[bool], value1 lazy.Value[T], value2 lazy.Value[T]) lazy.Value[T]

func TernaryLazy[T any](condition lazy.Value[bool], value1 lazy.Value[T], value2 lazy.Value[T]) lazy.Value[T] {
	return lazy.NewLazy(func() T {
		conditionValue := condition.Get()
		if conditionValue {
			return value1.Get()
		}
		return value2.Get()
	})
}
