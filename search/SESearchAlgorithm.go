package search

//import "fmt"
import "irs/helpers"
import "irs/lib"

func SESearchAlgorithm(objSearchRequest SearchRequest) SearchResponse {

	// declare search response object
	var objSearchResponse SearchResponse

	//fmt.Println("input " + objSearchRequest.SearchString)

	// remove the stop words from search string
	strSearchString := helpers.ReduceString(objSearchRequest.SearchString)

	//fmt.Println("reduced " + strSearchString)

	// tokenize the search string
	arrSearchString := helpers.Tokenize(strSearchString)

	// stem the search string
	arrSearchString = heloers.Stem(arrSearchString)

	//fmt.Println("stem " + strSearchString)

	//fmt.Println(arrSearchString)

	// get the document ID of the document with highest rank
	documents, err := lib.GetWeightedDocuments(arrSearchString)

	// return a false status if the ranked document was not obtained
	if err != nil {

		objSearchResponse.Status = "false"
		objSearchResponse.ExceptionCode = "2201"
		objSearchResponse.ExceptionMessage = "Could not obtain the highest rank document"

		return objSearchResponse

	}

	//fmt.Println(documents)

	// get the values of top n documents from the document map
	arrSearchResults, err := lib.GetBestDocuments(documents, int64(len(arrSearchString)))

	// return a false status if the document details were not obtained
	if err != nil {

		objSearchResponse.Status = "false"
		objSearchResponse.ExceptionCode = "2202"
		objSearchResponse.ExceptionMessage = "Could not obtain the document details"

		return objSearchResponse

	}

	objSearchResponse.SearchString = objSearchRequest.SearchString
	objSearchResponse.SearchResults = arrSearchResults
	objSearchResponse.Status = "true"
	objSearchResponse.ExceptionCode = "2200"
	objSearchResponse.ExceptionMessage = "-"

	return objSearchResponse

}
