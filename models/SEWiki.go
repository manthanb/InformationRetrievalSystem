package models

func SEWiki(objWikiRequest WikiRequest) WikiResponse {

	// declare response object
	var objWikiResponse WikiResponse

	// remove the stop words and the punctuations from the search string
	strSearchString := ReduceString(objWikiRequest.SearchString)

	// get the subject of the search string
	strSearchString = Clean(strSearchString)

	// get the desired link from wiki database
	strLink, strContent := GetLinkFromWiki(strSearchString)

	// copy link in the response object
	objWikiResponse.Link = strLink
	objWikiResponse.Content = strContent

	if GetDataFromRedis(strSearchString) == nil {

		// store in redis
		SetDataInRedis(strSearchString, strContent+" "+strLink)

	}

	// return the response object
	return objWikiResponse

}
