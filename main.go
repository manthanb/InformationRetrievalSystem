package main

import _ "irs/routers"
import "irs/lib"
import "github.com/astaxie/beego"

func main() {
	lib.ConnectToCache()
	beego.Run()
	lib.CloseGlobalSession()
}
