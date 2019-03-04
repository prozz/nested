package nested

// Resolve allows for easy lookups in deeply nested structs.
// In case of any unwanted nil, it just returns defaultValue.
func Resolve(defaultValue interface{}, accessFunc func() interface{}) (v interface{}) {
	defer func() {
		recover()
	}()
	v = defaultValue
	v = accessFunc()
	return v
}
