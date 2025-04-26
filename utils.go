package utils

func Filter[T any](list []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, t := range list {
		if fn(t) {
			result = append(result, t)
		}
	}
	return result
}

func FilterNotNil[T any](list []*T) []*T {
	result := make([]*T, 0)
	for _, t := range list {
		if t != nil {
			result = append(result, t)
		}
	}
	return result
}

func Map[T, R any](list []T, fn func(T) R) []R {
	result := make([]R, 0, len(list))
	for _, t := range list {
		result = append(result, fn(t))
	}
	return result
}

func Any[T any](list []T, fn func(T) bool) bool {
	for _, t := range list {
		if fn(t) {
			return true
		}
	}
	return false
}

func ToMap[T any, K comparable](list []T, fn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, t := range list {
		k := fn(t)
		if _, ok := result[k]; !ok {
			result[k] = []T{t}
		} else {
			result[k] = append(result[k], t)
		}
	}
	return result
}

func Distinct[T any](list []T) []T {
	seen := make(map[any]struct{})
	result := make([]T, 0)
	for _, t := range list {
		if _, ok := seen[t]; !ok {
			seen[t] = struct{}{}
			result = append(result, t)
		}
	}
	return result
}
