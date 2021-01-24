package sort

import "fmt"

type Sort struct {
	arr []int
}

func NewArr(mass []int) *Sort {
	return &Sort{mass}
}

func sortSelect(mas []int) []int {
	for i := 0; i < len(mas); i++ {
		fmt.Println(i * i)
	}
	return mas
}
