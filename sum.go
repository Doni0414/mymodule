package mymodule

func Sum(a ...int) int {
	var res = 0
	for _, num := range a {
		res += num
	}
	return res
}
