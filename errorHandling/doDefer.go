package errorhandling

import (
	"fmt"
	"os"
)

func DoDefer() {
	f := createFile("/Users/chandrasekarj/dev/a.txt")
	defer closeFile(f)
	writeFile(f, "Hello File writer.")
}

func createFile(p string) *os.File {
	fmt.Println("Create a File.")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, data string) {
	fmt.Println("Write a file.")
	fmt.Fprintln(f, data)
}

func closeFile(f *os.File) {
	fmt.Println("Close a file")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %v\n", err)
		os.Exit(1)
	}
}
