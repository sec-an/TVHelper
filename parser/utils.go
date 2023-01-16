package parser

import "fmt"

func duplicateRemoving[T any](s []T) []T {
	res := make([]T, 0, len(s))
	mySet := make(map[any]struct{})
	for _, t := range s {
		key := fmt.Sprint(t)
		if _, ok := mySet[key]; !ok {
			res = append(res, t)
			mySet[key] = struct{}{}
		}
	}
	return res
}

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
