package main

type Fighter struct {
	Name, aka string
}

func Find[T any](iterable []T, got func(T) bool) (value T, found bool) {
	for _, item := range iterable {
		if (got(item)) {
			return item, true
		}
	}
	return 
}