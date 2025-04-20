package util

// ifNull returns the value of the second argument if the value of the first argument is nil
func ifNull[T any](val *T, defaultVal T) T {
	if val != nil {
		return *val
	}
	return defaultVal
}
