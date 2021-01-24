package main

import (
	"fmt"
	"./sort"
)

func main()  {
	slice := []int{ 31, 72, 3, 14, 15, 6, 27, 18, 59, 10 }
	fmt.Println(sort.NewArr(slice))
}

