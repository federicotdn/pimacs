package core

import (
	"runtime"
)

type osType int

const (
	osUnknown osType = iota
	osLinux
	osWindows
	osMacOS
)

func getDefault[K comparable, V any](m map[K]V, key K, default_ V) V {
	val, ok := m[key]
	if !ok {
		return default_
	}
	return val
}

func getOS() osType {
	switch runtime.GOOS {
	case "linux":
		return osLinux
	case "windows":
		return osWindows
	case "darwin":
		return osMacOS
	default:
		return osUnknown
	}

}
