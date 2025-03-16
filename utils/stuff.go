package utils

import (
	"fmt"
	"slices"
)

func RemoveElement[T any](slice []T, fromIndex int) []T {
	return slices.Delete(slice, fromIndex, fromIndex+1)
}

func GetArrayElementAt[T any](array []T, i int) (T, error) {
	if i >= 0 && i < len(array) {
		return array[i], nil
	}

	var z T
	return z, fmt.Errorf("array out of bounds")
}

func IsElementExist[T any](array []T, index int) bool {
	return len(array) > index
}
