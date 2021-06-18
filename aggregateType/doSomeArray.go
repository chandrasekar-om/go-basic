package aggregateType

import (
	"fmt"
)

func DoArray() {
	fmt.Println("hello array")
	var l [2]string
	l[0] = "A"
	l[1] = "B"
	fmt.Println(l[0], l[1])
	fmt.Println(l)

	n := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	var r []int = n[1:4]
	fmt.Println(r)
}
