package utils

func Contains[T byte | string | int | rune](list []T, element T) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}

	return false
}
