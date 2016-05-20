package routers

import "search/controllers"
import "github.com/astaxie/beego"

func init() {
	beego.Router("/SESearchAlgorithm", &controllers.MainController{}, "get,post:SESearchAlgorithm")
	beego.Router("/SEWiki", &controllers.MainController{}, "get,post:SEWiki")
	beego.Router("/SEAnswer", &controllers.MainController{}, "get,post:SEAnswer")
}
