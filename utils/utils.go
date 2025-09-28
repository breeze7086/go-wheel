package utils

// GetPtr return the pointer of the object
func GetPtr[T any](v T) *T {
	return &v
}
