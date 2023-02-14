package jslikeutils

func smap[T any](arr []T, f func(e T, i int) T) []T {
	nArr := make([]T, len(arr))
	for i, e := range arr {
		nArr[i] = f(e, i)
	}
	return nArr
}

func filter[T any](arr []T, f func(e T, i int) bool) []T {
	nArr := make([]T, 0, len(arr))
	for i, e := range arr {
		if f(e, i) {
			nArr = append(nArr, e)
		}
	}
	return nArr
}

func reduce[T any](arr []T, f func(e T, n T, i int) T) T {
	var res T
	for i, e := range arr {
		res = f(res, e, i)
	}
	return res
}
