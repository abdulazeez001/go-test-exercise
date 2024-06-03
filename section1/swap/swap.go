package swap

import (
	"errors"
	"reflect"
)

// Swap swaps the values of two variables of any type without using a temporary variable
func Swap[T any](firstVar, secondVar *T) error {
	// Edge case validation: Check if either pointer is nil
	if firstVar == nil || secondVar == nil {
		return errors.New("nil pointer provided")
	}

	// Edge case validation: Check if the inputs are pointers
	if reflect.TypeOf(firstVar).Kind() != reflect.Ptr || reflect.TypeOf(secondVar).Kind() != reflect.Ptr {
		return errors.New("both arguments must be pointers")
	}

	// Swap the values
	*firstVar, *secondVar = *secondVar, *firstVar
	return nil
}
