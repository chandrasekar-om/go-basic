package ios

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DoLineFilters() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := strings.ToUpper(s.Text())
		fmt.Println("-->", v)
	}
	if e := s.Err(); e != nil {
		fmt.Fprintln(os.Stderr, " error ", e)
		os.Exit(1)
	}
}
