package interfaceType

import (
	"fmt"

	"github.com/chandrasekar-om/go-basic/closure"
)

type operation interface {
	drive(direction string, speed int) string
	fillTank(fuel int) bool
	getSeatCapacity() int
	getTankCapacity() int
}

type car struct {
	tankCap  int
	seatCap  int
	maxSpeed int
}

func (c car) drive(direction string, speed int) string {
	result := fmt.Sprintf("%s%d%s%d", "Car going in ",
		direction, " at a speed of ", speed)
	return result
}

func (c car) fillTank(fuel int) bool {
	if c.tankCap == fuel {
		return true
	}
	return false
}

func (c car) getSeatCapacity() int {
	return c.seatCap
}

func (c car) getTankCapacity() int {
	return c.tankCap
}

func DoDrive() {
	nextInt := closure.Count()
	c := car{40, 5, 180}
	fmt.Println(nextInt())
	fmt.Println(c.drive("EAST", 100))
	fmt.Println(nextInt())
	fmt.Println(c.getSeatCapacity())
	fmt.Println(nextInt())
	fmt.Println(c.getTankCapacity())
	fmt.Println(nextInt())
}
