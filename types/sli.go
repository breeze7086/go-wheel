package types

type Slice[T comparable] []T

// Contains check the slice contains the value, if it exists, return true, otherwise return false
func (s Slice[T]) Contains(val T) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

// Index get the index number of the value
func (s Slice[T]) Index(val T) int {
	for i, v := range s {
		if v == val {
			return i
		}
	}
	return -1
}

// Remove remove first item as per given value
func (s Slice[T]) Remove(val T) Slice[T] {
	idx := s.Index(val)
	if idx == -1 {
		return s // 没找到，原样返回
	}
	return append(s[:idx], s[idx+1:]...)
}

// RemoveAll remove all items as per given value
func (s Slice[T]) RemoveAll(val T) Slice[T] {
	result := make(Slice[T], 0, len(s))
	for _, v := range s {
		if v != val {
			result = append(result, v)
		}
	}
	return result
}

// Unique remove duplicate items, only keep 1
func (s Slice[T]) Unique() Slice[T] {
	seen := make(map[T]struct{})
	result := make(Slice[T], 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
