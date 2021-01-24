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

func (sort *Sort) SortBuble() bool {
	fmt.Printf("%#v\n", sort)
	for i := 0; i < len(sort.arr); i++ {
		for j := len(sort.arr) -1; j > i; j-- {
			if sort.arr[j-1] > sort.arr[j] {
				tmp := sort.arr[j-1]
				sort.arr[j-1] =  sort.arr[j]
				sort.arr[j] = tmp
			}
		}
	}
	fmt.Printf("%#v\n", sort)
	return true
}
