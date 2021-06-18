package referenceType

import (
	"fmt"
	"strings"
)

func DoSlice() {
	names := []string{"Harish", "Harini", "Srinikethan", "Srikrithi"}
	fmt.Println(names)

	a := names[1:3]
	b := names[1:2]
	fmt.Println(a, b)

	b[0] = "Love"
	fmt.Println(a, b)
	fmt.Println(names)

	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q)
	r := []bool{true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		j bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
	}
	fmt.Println(s)

	t := []int{1, 2, 3, 4, 5, 6}
	t = t[1:4]
	fmt.Println(t)
	t = t[:2]
	fmt.Println(t)
	t = t[1:]
	fmt.Print(t)

	u := []int{1, 2, 3, 4, 5, 6}
	printSlice(u)
	u = u[:0]
	printSlice(u)
	u = u[:4]
	printSlice(u)
	u = u[2:]
	printSlice(u)

	var v []int
	printSlice(v)
	if v == nil {
		fmt.Println("nil!")
	}
}

func DoSecondSlice() {
	a := make([]int, 5)
	fmt.Println(a)
	b := make([]int, 0, 5)
	fmt.Println(b)
	c := b[:2]
	fmt.Println(c)
	d := c[2:]
	fmt.Println(d)

}

func printSlice(a []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

}

func DoSliceOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	//fmt.Println(board)
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(board)
	for i := 0; i < len(board); i++ {
		fmt.Println("%s\n", strings.Join(board[i], ""))
	}
}

func DoSliceAppend() {
	s := []int{1, 2, 3, 4, 5}
	printSlice(s)
	s = append(s, 20, 30)
	printSlice(s)
}
