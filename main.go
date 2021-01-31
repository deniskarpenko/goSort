package main

import (
	//"./sort"
	"./socials"
)

func main()  {
	social := socials.NewFacebook();
	social.SetToken("fb");
	println(social.GetToken())
	//slice := []int{ 31, 72, 3, 14, 15, 6, 27, 18, 59, 10 }
	//sort := sort.NewSort(slice)
	//sort.SortSelect()
	//sort.SortBuble()
	//sort.SortInsert()
}

