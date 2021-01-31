package main

import (
	//"./sort"
	"./socials"
)

func main()  {
	social := socials.NewFacebook();
	social.SetToken("EAAXNr3F3gqYBAKJ0w27NKJJZAVxUXKZBTztpxTExQRLXno9088BifiBYe7OjJnWuiZAToFjosZBXZAZCfnbjl9WDAN5XENqneUUx8AiqAaFzIKyGeetZAC8ZAAMkH2SGWVl9epZCa6BWhrE7kEVQWZAUyyozEQl1WfGl7egL7wZCrJup1wl5vA3MC64lhxMzRyjZBkgU09gzGhpM0wZDZD");
	meInfo,status := social.GetMe()
	if status {
		println(meInfo)
	} else {
		println(social.GetToken())
	}
	//slice := []int{ 31, 72, 3, 14, 15, 6, 27, 18, 59, 10 }
	//sort := sort.NewSort(slice)
	//sort.SortSelect()
	//sort.SortBuble()
	//sort.SortInsert()
}

