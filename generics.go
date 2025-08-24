package main

func getValue[T any](data map[string]interface{}, key string) (T, bool) {
	val, ok := data[key]
	if !ok {
		var zero T
		return zero, false
	}
	converted, ok := val.(T)
	return converted, ok
}
