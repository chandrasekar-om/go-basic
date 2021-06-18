package commonFunctions

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func IsAvaliable(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func FMain() {
	fruits := []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(fruits, "pear"))
	fmt.Println(IsAvaliable(fruits, "apple"))
	f := func(s string) bool {
		return strings.HasPrefix(s, "p")
	}
	fmt.Println(Any(fruits, f))
	fmt.Println(All(fruits, f))
	fmt.Println(Filter(fruits, func(s string) bool {
		return strings.Contains(s, "e")
	}))
	fmt.Println(Map(fruits, strings.ToUpper))
}
