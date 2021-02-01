package main

import (
	//"./sort"
	"./socials"
)

func main()  {
	social := socials.FactorySocial("instagram");
	social.SetToken("EAAXNr3F3gqYBAGplA5DQpiKdGgsietyMPVJnKZBBAT8TlZASZCbpIkMZBGOWDoftzjk1QS350X6dcdmxQzsd0DY2ubkgFzsaxodxfF9BreAZCI9lvJtMkoBlWcSdZCj1KB1mlFUXZCh2sIM6lqY6P071FHS4w9VZAbOBxZCAZC0je5bp4gYJkStfyHLv9rxcYV18uvA1I2h9dgBHkZBZC7OG9oUwfifrO1tzLMhSSyR55JZCObW9d3NlpZBPk3qsM2ChEZBTJ4ZD")
	followers,status := social.GetFollowers("sargonApp")
	if status {
		println(followers)
	} else {
		println(social.GetToken())
	}
	//slice := []int{ 31, 72, 3, 14, 15, 6, 27, 18, 59, 10 }
	//sort := sort.NewSort(slice)
	//sort.SortSelect()
	//sort.SortBuble()
	//sort.SortInsert()
}

