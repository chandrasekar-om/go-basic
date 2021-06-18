package ios

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const path = "/Users/chandrasekarj/dev/git/filesystem/go"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func DoRead() {

	dat, err := ioutil.ReadFile(path + "/data.txt")
	check(err)
	fmt.Println(string(dat))

	f, err := os.Open(path + "/data.txt")
	check(err)
	b1 := make([]byte, 5)
	c, err := f.Read(b1)
	check(err)
	fmt.Println(string(b1[:c]))
	c1, err := f.Seek(10, 0)
	check(err)

	b2 := make([]byte, 5)
	c2, err1 := f.Read(b2)
	check(err1)
	fmt.Println("%d bytes @ %d: %s\n", c2, c1, string(b2))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()

}

func DoWrite() {
	d := []byte("hello\ngo\nhello\nchandrasekarj\n")
	err := ioutil.WriteFile(path+"/data.txt", d, 0644)
	check(err)

	f, err := os.Create(path + "/data1.txt")
	check(err)
	defer f.Close()

	d1 := []byte{8, 10, 12, 13, 14, 15, 16, 30}
	n, err := f.Write(d1)
	check(err)
	fmt.Println("Write : ", n)

	n1, err := f.WriteString("Hello\nSrinikethan\n")
	check(err)
	fmt.Println("Write : ", n1)

	n2, err := f.WriteAt(d, 10)
	check(err)
	fmt.Println("Write AT : ", n2)

	f.Sync()

	w := bufio.NewWriter(f)
	n3, err := w.WriteString("Hello\nHarini\nHello\nSrikrithi\n")
	check(err)
	fmt.Println("Write Buff : ", n3)
	w.Flush()
}
