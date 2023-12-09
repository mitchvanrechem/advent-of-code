package utils

func Pop[s any](slice *[]s) s {
	sliceLen := len(*slice)
	popped := (*slice)[sliceLen-1]

	*slice = (*slice)[:sliceLen-1]

	return popped
}
