package helpers

import "strings"
import "github.com/kljensen/snowball"
import "fmt"

// it is called whenever an error occurs
// input - error
func HandleError(err error) {

	//panic the error if it is not nil
	if err != nil {
		panic(err)
	}

}

// removes the punctuations and stop words from a string
// input - string
// output - string without stop words
func ReduceString(strSearchString string) string {

	// convert the string to lower case
	strReducedString := strings.ToLower(strSearchString)

	// remove the punctuation marks
	strReducedString = strings.Replace(strReducedString, ".", "", -1)
	strReducedString = strings.Replace(strReducedString, ",", "", -1)
	strReducedString = strings.Replace(strReducedString, "!", "", -1)
	strReducedString = strings.Replace(strReducedString, "?", "", -1)
	strReducedString = strings.Replace(strReducedString, ":", "", -1)

	// remove the stop words
	strReducedString = strings.Replace(strReducedString, " a ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " an ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " and ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " are ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " as ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " at ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " be ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " by ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " for ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " from ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " has ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " he ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " in ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " is ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " it ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " its ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " of ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " on ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " that ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " the ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " to ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " was ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " were ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " will ", " ", -1)
	strReducedString = strings.Replace(strReducedString, " with ", " ", -1)

	// return the string
	return strReducedString

}

// stems a string
// input - string
// output - stemmed string
func Stem(arrSearchString []string) []string {

	// call the porter's algorithm
	// return string(porterstemmer.StemWithoutLowerCasing([]rune(strSearchString)))

	var arrStemmedString []string

	for _, token := range arrSearchString {

		strStemmedString, err := snowball.Stem(token, "english", true)
		fmt.Println(err)

		arrStemmedString = append(arrStemmedString, strStemmedString)

	}

	return arrStemmedString

}

// tokenizes a string
// input - string
// output - array of tokens of the input string
func Tokenize(strSearchString string) []string {

	// split the string on blank spaces
	return strings.Split(strSearchString, " ")

}

// remove unimportant words from the search string
// input - string
// output - string with words removed
func Clean(strSearchString string) string {

	// remove words
	strCleanString := strings.Replace(strSearchString, "what", "", -1)
	strCleanString = strings.Replace(strCleanString, "who", "", -1)
	strCleanString = strings.Replace(strCleanString, "when", "", -1)
	strCleanString = strings.Replace(strCleanString, "how", "", -1)
	strCleanString = strings.Replace(strCleanString, "how many", "", -1)
	strCleanString = strings.Replace(strCleanString, "did", "", -1)
	strCleanString = strings.Replace(strCleanString, "does", "", -1)
	strCleanString = strings.Replace(strCleanString, "define", "", -1)
	strCleanString = strings.Replace(strCleanString, "defination", "", -1)
	strCleanString = strings.Replace(strCleanString, "means", "", -1)
	strCleanString = strings.Replace(strCleanString, " do ", "", -1)
	strCleanString = strings.Replace(strCleanString, "meaning", "", -1)
	strCleanString = strings.Replace(strCleanString, "mean", "", -1)
	strCleanString = strings.Replace(strCleanString, "capital", "", -1)
	strCleanString = strings.Replace(strCleanString, "'s", "", -1)
	strCleanString = strings.Replace(strCleanString, " i", "", -1)
	strCleanString = strings.Replace(strCleanString, " we", "", -1)
	strCleanString = strings.Replace(strCleanString, "i ", "", -1)
	strCleanString = strings.Replace(strCleanString, "we ", "", -1)
	strCleanString = strings.Replace(strCleanString, " i ", "", -1)
	strCleanString = strings.Replace(strCleanString, " we ", "", -1)
	strCleanString = strings.Replace(strCleanString, "you", "", -1)
	strCleanString = strings.Replace(strCleanString, "   ", " ", -1)
	strCleanString = strings.Replace(strCleanString, "  ", " ", -1)

	// remove leading and trailing spaces
	strCleanString = strings.TrimSpace(strCleanString)

	return strCleanString

}
