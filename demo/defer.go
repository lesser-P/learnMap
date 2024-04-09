package demo

func DeferFunc() int {
	i := 0
	defer func() {
		i++
	}()
	return i
}
func DeferFuncNameBack() (i int) {
	defer func() {
		i++
	}()
	return i
}
