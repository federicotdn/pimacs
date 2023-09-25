package core

func getDefault[K comparable, V any](m map[K]V, key K, default_ V) V {
	val, ok := m[key]
	if !ok {
		return default_
	}
	return val
}
