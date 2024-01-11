package util

import "reflect"

func MergeArray[T any](first []T, second []T) []T {
	first = append(first, second...)
	return first
}

func MergeUniqueArray[T any](first []T, second []T) []T {
	valueMap := map[any]bool{}
	for _, value := range first {
		valueMap[value] = true
	}
	for _, value := range second {
		_, ok := valueMap[value]
		if !ok {
			first = append(first, value)
			valueMap[value] = true
		}
	}
	return first
}

func UniqueArray[T any](values []T) []T {
	newArray := make([]T, 0)
	tempMap := make(map[string][]T, len(values))
	for _, v := range values {
		tempArray, exist := tempMap[AnyToStr(v)]
		if exist {
			found := false
			for _, e := range tempArray {
				if reflect.DeepEqual(e, v) {
					found = true
					break
				}
			}
			if !found {
				tempMap[AnyToStr(v)] = append(tempArray, v)
				newArray = append(newArray, v)
			}
		} else {
			tempMap[AnyToStr(v)] = []T{v}
			newArray = append(newArray, v)
		}
	}
	return newArray
}
