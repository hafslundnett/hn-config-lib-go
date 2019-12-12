package util

// MapToSlice takes a map and converts it to a slice on the form of {key1, value1, key2, value2, .....}.
// Go's maps are intentionally random in order, thus the order of the slice will also be random,
// with the exception that keys will always come right before their corresponding values.
func MapToSlice(m map[interface{}]interface{}) (s []interface{}) {
	for k, v := range m {
		s = append(s, k, v)
	}
	return s
}
