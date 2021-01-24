package sort

import "fmt"

type Sort struct {
	arr []int
}

func NewSort(mass []int) *Sort {
	return &Sort{mass}
}

func (sort *Sort) SortSelect() bool {
	fmt.Printf("%#v\n", sort)
	for i := 0; i < len(sort.arr); i++ {
		for j := 0; j < len(sort.arr); j++ {
			if sort.arr[j] > sort.arr[i] {
				tmp := sort.arr[j]
				sort.arr[j] = sort.arr[i]
				sort.arr[i] = tmp
			}
		}
	}
	fmt.Printf("%#v\n", sort)
	return true
}
