package models

func SEAnswer(objAnswerRequest AnswerRequest) AnswerResponse {

	// declare response object
	var objAnswerResponse AnswerResponse

	// remove the stop words and the punctuations from the search string
	strSearchString := ReduceString(objAnswerRequest.SearchString)

	// get the subject of the search string
	strSearchString = Clean(strSearchString)

	// get the desired link from wiki database
	bytAnswer := GetDataFromRedis(strSearchString)

	// copy link in the response object
	objAnswerResponse.Answer = string(bytAnswer)
	objAnswerResponse.Key = strSearchString

	if objAnswerResponse.Answer == " " {
		objAnswerResponse.Key = " "
	}

	// return the response object
	return objAnswerResponse

}
