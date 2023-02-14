package jslikeutils

// Map applies a function to each element of array and returns it.
// Function f receives the parameter e, which is the element of
// the array that is currently being processed, and parameter i,
// which is the index. It should return the processed element.
func Map[T any](arr []T, f func(e T, i int) T) []T {
	nArr := make([]T, len(arr))
	for i, e := range arr {
		nArr[i] = f(e, i)
	}
	return nArr
}

// Filter returns only elements of array which satisfies a condition.
// Function f receives the parameter e, which is the element of
// the array that is currently being processed, and parameter i,
// which is the index. It should return a boolean indicating whether
// that element must be included or not.
func Filter[T any](arr []T, f func(e T, i int) bool) []T {
	nArr := make([]T, 0, len(arr))
	for i, e := range arr {
		if f(e, i) {
			nArr = append(nArr, e)
		}
	}
	return nArr
}

// Reduce returns the result of procesing all elements in given array.
// Function f receives the parameter e, which is the element that contains
// the result of previous iteration, parameter n, that is the element of
// the array that is currently being processed, and parameter i,
// which is the index. It should return the processed element.
func Reduce[T any](arr []T, f func(e T, n T, i int) T) T {
	var res T
	for i, e := range arr {
		res = f(res, e, i)
	}
	return res
}
