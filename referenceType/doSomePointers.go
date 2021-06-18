package referenceType

import "fmt"

func DoPointer() {
	v := 0
	fmt.Println("initial : ", v)
	updateValue(v)
	fmt.Println("Update Value : ", v)
	updatePointer(&v)
	fmt.Println("Update Pointer : ", v)
	fmt.Println("Pointer : ", &v)
}

func updateValue(v int) {
	v = 10
}

func updatePointer(v *int) {
	*v = 100
}
