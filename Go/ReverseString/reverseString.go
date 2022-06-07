package ReverseString

func ReverseString(str string) string {
	var r string

	for i := len(str) - 1; i >= 0; i-- {
		r = r + string(str[i])
	}
	return r
}

func ReverseStr(str string) string {
	var rvs string

	for _, r := range str {
		rvs = string(r) + rvs
	}
	return rvs
}
