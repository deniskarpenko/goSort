package main

import (
	//"./sort"
	"./socials"
)

func main()  {
	idFacebookApp := "1633528290181798"
	secretFacebookApp := "6e09c6911662367916a88d58dff59bcb"
	facebookAppToken := idFacebookApp + "|" + secretFacebookApp
	//facebookAppToken := "EAAXNr3F3gqYBAM5XnucaDc6WKZAgAM2ZBtJlx7d1bIZCfVhd93SJWWZAZBynYypz8HxHovryBFseZCk7StPg8iFn1ZAqsCDLi18dkjRDMb9T1INWvaPCsiKlNSZCLrC7niOqzeCKxlFdwUpCpOEUQxHDzeQRkbOkehW8mkZCEpbrSdoQeXDFlIrjQt1Ys7by0T7bnIq4iiwb29BxZCimHKVShf"
	social := socials.FactorySocial("fb");
	social.SetToken(facebookAppToken)
	//followers,status := social.GetMe()
	followers,status := social.GetFollowers("netpeak")
	if status {
		println(followers)
	} else {
		println(social.GetToken())
	}
	/*
		{"error":{"message":"(#100) Object does not exist, cannot be loaded due to missing permission or reviewable feature, or does not support this operation. This endpoint requires the 'pag
	es_read_engagement' permission or the 'Page Public Content Access' feature or the 'Page Public Metadata Access' feature. Refer to https:\/\/developers.facebook.com\/docs\/apps\/review\
	/login-permissions#manage-pages, https:\/\/developers.facebook.com\/docs\/apps\/review\/feature#reference-PAGES_ACCESS and https:\/\/developers.facebook.com\/docs\/apps\/review\/featur
	e#page-public-metadata-access for details.","type":"OAuthException","code":100,"fbtrace_id":"AEd7sp_0tc6-T98MZ3HWJ47"}}
	*/
	//slice := []int{ 31, 72, 3, 14, 15, 6, 27, 18, 59, 10 }
	//sort := sort.NewSort(slice)
	//sort.SortSelect()
	//sort.SortBuble()
	//sort.SortInsert()
}

