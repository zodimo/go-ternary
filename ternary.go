package ternary

import "github.com/zodimo/go-lazy"

func Ternary[T any](condition bool, value1 T, value2 T) T {
	if condition {
		return value1
	}
	return value2
}

func TernaryLazy[T any](condition lazy.Value[bool], value1 lazy.Value[T], value2 lazy.Value[T]) lazy.Value[T] {
	return lazy.NewLazy(func() T {
		conditionValue := condition.Get()
		if conditionValue {
			return value1.Get()
		}
		return value2.Get()
	})
}
