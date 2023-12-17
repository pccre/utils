package shared // shared functions

// fast, doesn't keep order
func Remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// keep order
func RemoveO[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
