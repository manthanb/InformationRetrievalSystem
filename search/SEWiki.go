package search

import "search/helpers"
import "search/lib"

func SEWiki(objWikiRequest WikiRequest) WikiResponse {

	// declare response object
	var objWikiResponse WikiResponse

	// remove the stop words and the punctuations from the search string
	strSearchString := helpers.ReduceString(objWikiRequest.SearchString)

	// get the subject of the search string
	strSearchString = helpers.Clean(strSearchString)

	// get the desired link from wiki database
	strLink, strContent := lib.GetLinkFromWiki(strSearchString)

	// copy link in the response object
	objWikiResponse.Link = strLink
	objWikiResponse.Content = strContent

	if lib.GetDataFromRedis(strSearchString) == nil {

		// store in redis
		lib.SetDataInRedis(strSearchString, strContent+" "+strLink)

	}

	// return the response object
	return objWikiResponse

}
