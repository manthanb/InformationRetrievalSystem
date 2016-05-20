package main

import _ "search/routers"
import "search/models"
import "github.com/astaxie/beego"

func main() {
	models.ConnectToCache()
	beego.Run()
	models.CloseGlobalSession()
}
