package util

// ifTernary returns the first value when the condition evaluates to true, otherwise it returns the second value
func ifTernary[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	}
	return ifFalse
}
