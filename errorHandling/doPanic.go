package errorhandling

import "os"

func DoPanic() {

	//panic("a problem")

	_, err := os.Create("hello/one.html")
	if err != nil {
		//panic("Unable to create a file")
	}
}
