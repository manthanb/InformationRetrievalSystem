package models

import "github.com/astaxie/beego"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "errors"
import "math"
import "strings"

var GlobalSession *mgo.Session
var GlobalSession2 *mgo.Session

// set up a connection to mongo
func ConnectToMongo() *mgo.Session {

	// create a global connection if it does not exist
	if GlobalSession == nil {

		err := errors.New("")

		// read Mongo IP and Mongo Port from config
		strMongoIp := beego.AppConfig.String("mongoip")
		strMongoPort := beego.AppConfig.String("mongoport")
		strMongoLoc := strMongoIp + ":" + strMongoPort

		// create the global session
		GlobalSession, err = mgo.Dial(strMongoLoc)
		HandleError(err)

	}

	// return a copy of the global session
	return GlobalSession.Copy()

}

// set up a connection to mongo
func ConnectToMongo2() *mgo.Session {

	// create a global connection if it does not exist
	if GlobalSession2 == nil {

		err := errors.New("")

		// create the global session
		GlobalSession2, err = mgo.Dial("127.0.0.1:27018")
		HandleError(err)

	}

	// return a copy of the global session
	return GlobalSession2.Copy()

}

// closes the global session
func CloseGlobalSession() {
	GlobalSession.Close()
	GlobalSession2.Close()
}

// returns the document with highest rank
func GetWeightedDocuments(arrSearchString []string) (map[int64]float64, error) {

	// get a mongo session
	mongoSession := ConnectToMongo()
	defer mongoSession.Close()

	// create object of collTerm and map
	var objCollTerm CollTerm
	var documents map[int64]float64

	// make the documents map
	documents = make(map[int64]float64)

	// connect to database
	strDB := beego.AppConfig.String("dbname")
	strCollection := beego.AppConfig.String("collection")
	collTermCollection := mongoSession.DB(strDB).C(strCollection)

	// for every term of the search string, add documents to the hash map
	for _, term := range arrSearchString {

		// get the postings corresponding to a particular term
		_ = collTermCollection.Find(bson.M{"term": term}).One(&objCollTerm)

		// store the inverse document frequency in a temporary variable
		fltIdf := objCollTerm.Idf

		// parse all postings of a term
		for _, posting := range objCollTerm.Postings {

			// check if the document id is already present in the hash map
			_, ok := documents[posting.DocId]

			// if not present then add a new entry and initialize it
			if !ok {
				documents[posting.DocId] = posting.Weight * fltIdf
				continue
			}

			// increment the value of corresponfing doc ID
			documents[posting.DocId] = documents[posting.DocId] + posting.Weight*fltIdf

		}

	}

	// return the documents map
	return documents, nil

}

func GetBestDocuments(documents map[int64]float64, intLength int64) ([]SearchResult, error) {

	// get a mongo session
	mongoSession := ConnectToMongo()
	defer mongoSession.Close()

	// create array and object of search result
	var arrBestDocuments []SearchResult
	var objBestDocument SearchResult
	var arrDocuments []Documents
	var objDocuments Documents

	// declare a max
	fltMax := -1.0

	// connect to database
	strDB := beego.AppConfig.String("dbname")
	strCollection := beego.AppConfig.String("documents")
	documentsCollection := mongoSession.DB(strDB).C(strCollection)

	//find the maximum weight from hash map
	for i, _ := range documents {

		if fltMax < documents[i] {
			fltMax = documents[i]
		}

	}

	// calculate the minimum length that should match - currently 60%
	fltMinMatchLength := float64(intLength) * float64(0.8)

	// calculte the minimum weight assuming each matching term occurs once
	fltMinWeight := fltMinMatchLength * math.Log(2)

	// calculate the cut off factor by multiplying by lower average idf
	fltCutOff := fltMinWeight * 3.1

	// get the documents with maximum weight
	for i, _ := range documents {

		// if the weight of the document is maximum
		if documents[i] == fltMax || documents[i] > fltCutOff {

			// get the document
			err := documentsCollection.Find(bson.M{"quesId": i}).One(&objDocuments)
			HandleError(err)

			// return error if document was not retrieved
			if err != nil {
				return nil, err
			}

			// calculate the term ratio
			objDocuments.TermRatio = math.Abs(1.0 - (float64(objDocuments.Terms) / float64(intLength)))

			// append the document to document array to sort later
			arrDocuments = append(arrDocuments, objDocuments)

		}

	}

	// sort documents by terms if more than one document has maximum weight
	if len(arrDocuments) > 1 {
		arrDocuments = SortDocuments(arrDocuments)
	}

	// copy the question and answer strings to the data structure that is to be returned
	// add answers to redis
	for i, _ := range arrDocuments {

		objBestDocument.QuestionText = arrDocuments[i].QuestionText
		objBestDocument.AnswerText = arrDocuments[i].AnswerText

		// add to redis only if the score is atleast half of the average score
		if arrDocuments[i].Score > 0.5 {

			strQuestion := ReduceString(arrDocuments[i].QuestionText)
			strQuestion = Clean(strQuestion)

			if len(strings.Split(strQuestion, " ")) <= 6 {
				SetDataInRedis(strQuestion, arrDocuments[i].AnswerText)
			}

		}

		arrBestDocuments = append(arrBestDocuments, objBestDocument)

	}

	// return the best documents
	return arrBestDocuments, nil
}

// gets link from wiki database related to the search query
func GetLinkFromWiki(strSearchString string) (string, string) {

	// declare an object of wiki data and link to be returned
	var objWikiData WikiData
	var strLink string
	var strContent string

	// get a mongo session
	mongoSession := ConnectToMongo2()
	defer mongoSession.Close()

	// connect to database
	strDB := beego.AppConfig.String("dbname")
	strCollection := beego.AppConfig.String("wikipedia")
	wikiCollection := mongoSession.DB(strDB).C(strCollection)

	// split the search string to get terms
	arrSearchString := strings.Split(strSearchString, " ")

	// declare the regex string
	var strRegex string

	// for earch term in the search string, get a link
	for _, term := range arrSearchString {

		// no point searching if the term is small or empty
		if term == " " {
			continue
		}

		// keep appending the terms to regex
		strRegex = strRegex + term + ".*"

	}

	// add prefix to the regex
	strRegex = ".*" + strRegex

	// write bson map for the search criteria
	mapSearchCriteria := bson.M{"$regex": bson.RegEx{strRegex, "i"}}

	// find a document
	err := wikiCollection.Find(bson.M{"key": mapSearchCriteria}).One(&objWikiData)

	// no need to go further if a document is found
	if err != nil {
		strLink = ""
		strContent = ""
		return strLink, strContent
	}

	// copy the link
	strLink = objWikiData.Link
	strContent = objWikiData.Content

	// return the link
	return strLink, strContent

}
