package hfunk

func Includes(arr []any, find any) bool {
	for _, item := range arr {
		if item == find {
			return true
		}
	}

	return false
}
