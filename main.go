package main

//	@title			Raising
//	@version		1.0
//	@description	Golang work.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	zhangxin
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		39.108.145.195:3000
//	@BasePath	/api/v1
import (
	"Raising/conf"
	"Raising/router"
	"Raising/util"
)

func main() {
	conf.Init()
	util.Init()
	r := router.NewRouter()
	r.Run(conf.Port)
}
