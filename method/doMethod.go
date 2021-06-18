package method

import "fmt"

type pattu struct {
	name string
	age  int
	food string
	game string
}

func (p pattu) eat() string {
	return p.food
}

func (p pattu) play() string {
	return p.game
}

func DoMethod() {
	p := pattu{"Strinikethan", 1, "veg food", "toy play"}
	fmt.Println(p.play())
	fmt.Println(p.eat())
}
