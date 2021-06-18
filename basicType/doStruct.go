package basictype

import "fmt"

type fruit struct {
	name  string
	color string
}

func DoStruct() {
	apple := fruit{
		name:  "apple",
		color: "red",
	}
	fmt.Println(apple)
}
