package main

func RepeatStr(repetitions int, value string)(s string) {
	for i := 0; i < repetitions; i ++ {
		s += value
	}
	return s
}