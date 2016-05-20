package controllers

import "github.com/astaxie/beego"
import "search/models"
import "encoding/json"

type MainController struct {
	beego.Controller
}

func (c *MainController) SESearchAlgorithm() {

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	// declare request and response objects for search algorithm
	var objSearchRequest models.SearchRequest
	var objSearchResponse models.SearchResponse

	// get the JSON parameter from the request
	strJsonQuery := c.Ctx.Input.Query(beego.AppConfig.String("requestJsonQuery"))

	// decode the JSON
	err := json.Unmarshal([]byte(strJsonQuery), &objSearchRequest)

	// if unmarshal was unsuccessfull, return error message
	if err != nil {

		objSearchResponse.Status = "false"
		objSearchResponse.ExceptionCode = "2101"
		objSearchResponse.ExceptionMessage = "Could not decode the request JSON"

		jsonObjSearchResponse, _ := json.Marshal(objSearchResponse)

		c.Ctx.WriteString(string(jsonObjSearchResponse))

		return

	}

	// call the models method for search algorithm
	objSearchResponse = models.SESearchAlgorithm(objSearchRequest)

	// return the algorithm reponse
	jsonObjSearchResponse, _ := json.Marshal(objSearchResponse)
	c.Ctx.WriteString(string(jsonObjSearchResponse))

}

func (c *MainController) SEWiki() {

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	// declare request and response objects for search algorithm
	var objWikiRequest models.WikiRequest
	var objWikiResponse models.WikiResponse

	// get the JSON parameter from the request
	strJsonQuery := c.Ctx.Input.Query(beego.AppConfig.String("requestJsonQuery"))

	// decode the JSON
	err := json.Unmarshal([]byte(strJsonQuery), &objWikiRequest)

	// if unmarshal was unsuccessfull, return error message
	if err != nil {

		objWikiResponse.Link = ""

		jsonObjWikiResponse, _ := json.Marshal(objWikiResponse)

		c.Ctx.WriteString(string(jsonObjWikiResponse))

		return

	}

	// call the models method for search algorithm
	objWikiResponse = models.SEWiki(objWikiRequest)

	// return the algorithm reponse
	jsonObjWikiResponse, _ := json.Marshal(objWikiResponse)
	c.Ctx.WriteString(string(jsonObjWikiResponse))

}

func (c *MainController) SEAnswer() {

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

	// declare request and response objects for search algorithm
	var objAnswerRequest models.AnswerRequest
	var objAnswerResponse models.AnswerResponse

	// get the JSON parameter from the request
	strJsonQuery := c.Ctx.Input.Query(beego.AppConfig.String("requestJsonQuery"))

	// decode the JSON
	err := json.Unmarshal([]byte(strJsonQuery), &objAnswerRequest)

	// if unmarshal was unsuccessfull, return error message
	if err != nil {

		objAnswerResponse.Answer = ""

		jsonObjAnswerResponse, _ := json.Marshal(objAnswerResponse)

		c.Ctx.WriteString(string(jsonObjAnswerResponse))

		return

	}

	// call the models method for search algorithm
	objAnswerResponse = models.SEAnswer(objAnswerRequest)

	// return the algorithm reponse
	jsonObjAnswerResponse, _ := json.Marshal(objAnswerResponse)
	c.Ctx.WriteString(string(jsonObjAnswerResponse))

}
