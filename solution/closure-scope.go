package solution

func CreateCallbacks(num int) []func() int {
	sliceFunction := []func() int{}
	for i := 0; i < num; i++ {
		j := i
		sliceFunction = append(sliceFunction, func() int {
			return j
		})
	}
	return sliceFunction
}
