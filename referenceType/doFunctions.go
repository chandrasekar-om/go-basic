package referenceType

import "fmt"

var add = func(a, b int) int {
	return a + b
}
var sub = func(a, b int) int {
	return a - b
}
var mul = func(a, b int) int {
	return a * b
}
var dev = func(a, b int) int {
	return a / b
}

type divide func(i, j int) int

func DoFunction() {
	func() {
		fmt.Println("Welcome to Function !")
	}()

	func(n string) {
		fmt.Println(n, "Welcome to Function !")
	}("Developer")

	printFunc(addition(1, 2))
	a, b := swap(1, 2)
	printFunc(a, b)
	subs(10, 2, printFunc)
	simple(2, 2, mul, printFunc)

	var c divide = func(i, j int) int {
		return i / j
	}
	r := c(10, 5)
	printFunc(r)
	printFunc(math(1)(100, 200))

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	nextInt := intSeq()
	for i := 0; i < 10; i++ {
		printFunc(nextInt())
	}
}

func addition(a int, b int) int {
	return a + b
}

func printFunc(r ...int) {
	fmt.Println(r)
}

func swap(x, y int) (int, int) {
	return y, x
}

func subs(a, b int, pr func(p ...int)) {
	pr(a - b)
}

func simple(i, j int, x func(a, b int) int, y func(p ...int)) {
	y(x(i, j))
}

func math(f int) func(a, b int) int {
	switch f {
	case 1:
		return add
	case 2:
		return sub
	case 3:
		return mul
	case 4:
		return dev
	}
	return nil
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
