package search

import "irs/helpers"
import "irs/lib"

func SEAnswer(objAnswerRequest AnswerRequest) AnswerResponse {

	// declare response object
	var objAnswerResponse AnswerResponse

	// remove the stop words and the punctuations from the search string
	strSearchString := helpers.ReduceString(objAnswerRequest.SearchString)

	// get the subject of the search string
	strSearchString = helpers.Clean(strSearchString)

	// get the desired link from wiki database
	bytAnswer := lib.GetDataFromRedis(strSearchString)

	// copy link in the response object
	objAnswerResponse.Answer = string(bytAnswer)
	objAnswerResponse.Key = strSearchString

	if objAnswerResponse.Answer == " " {
		objAnswerResponse.Key = " "
	}

	// return the response object
	return objAnswerResponse

}
