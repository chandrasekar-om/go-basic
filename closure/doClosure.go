package closure

func Count() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
