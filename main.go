package main

import (
	"./sort"
	"fmt"
)

func main()  {
	slice := []int{ 31, 72, 3, 14, 15, 6, 27, 18, 59, 10 }
	sort := sort.NewSort(slice)
	fmt.Printf("%#v\n", sort)
	sort.SortSelect()
}

