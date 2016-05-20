package models

type SearchRequest struct {
	SearchString string `json:"SearchString"`
}

type SearchResult struct {
	QuestionText string `json:"questionText"`
	AnswerText   string `json:"answerText"`
}

type SearchResponse struct {
	SearchString     string         `json:"searchString"`
	SearchResults    []SearchResult `json:"searchResults"`
	Status           string         `json:"status"`
	ExceptionCode    string         `json:"exceptionCode"`
	ExceptionMessage string         `json:"exceptionMessage"`
}

type Documents struct {
	Id           int64   `json:"quesId" bson:"quesId"`
	QuestionText string  `json:"ques" bson:"ques"`
	AnswerText   string  `json:"ans" bson:"ans"`
	Score        float64 `json:"score" bson:"score"`
	Terms        int64   `json:"terms" bson:"terms"`
	TermRatio    float64 `json:"termRatio" bson:"termRatio"`
}

type Terms struct {
	Id     string `json:"_id" bson:"_id"`
	TermId int64  `json:"termId" bson:"termId"`
	Term   string `json:"term" bson:"term"`
}

type Posting struct {
	DocId  int64   `json:"docId" bson:"docId"`
	Weight float64 `json:"weight" bson:"weight"`
}

type CollTerm struct {
	Id       string    `json:"_id" bson:"_id"`
	Pid      int64     `json:"pid" bson:"pid"`
	Term     string    `json:"term" bson:"term"`
	Idf      float64   `json:"idf" bson:"idf"`
	Df       int64     `json:"df" bson:"df"`
	Postings []Posting `json:"postings" bson:"postings"`
}

type WikiData struct {
	Id      string `json:"_id" bson:"_id"`
	Key     string `json:"key" bson:"key"`
	Link    string `json:"link" bson:"link"`
	Content string `json:"content" bson:"content"`
}

type WikiRequest struct {
	SearchString string `json:"searchString" bson:"searchString"`
}

type WikiResponse struct {
	Link    string `json:"link" bson:"link"`
	Content string `json:"content" bson:"content"`
}

type AnswerRequest struct {
	SearchString string `json:"searchString" bson:"searchString"`
}

type AnswerResponse struct {
	Key    string `json:"key" bson:"key"`
	Answer string `json:"answer" bson:"answer"`
}

var countries map[string]string
