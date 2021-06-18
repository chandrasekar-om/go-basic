package sorting

import (
	"fmt"
	"sort"
)

func DoSorting() {
	sl := []string{"d", "z", "a", "o"}
	sort.Strings(sl)
	fmt.Println(sl)

	il := []int{8, 3, 9, 5}
	sort.Ints(il)
	fmt.Println(il)
	f := sort.IntsAreSorted(il)
	fmt.Println(f)

}

type list []string

func (l list) Len() int {
	return len(l)
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l list) Less(i, j int) bool {
	return len(l[i]) < len(l[j])
}

func DoSortingByFunctions() {
	names := []string{"harish", "srinikethan", "harini", "srikrithi"}
	sort.Strings(list(names))
	fmt.Println(names)
}
