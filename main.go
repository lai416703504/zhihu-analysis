package main

import (
	"zhihu-analysis/app/route"

	"github.com/gin-gonic/gin"
)

func main() {
	//show version
	// ver := flagutil.GetVersion()
	// if *ver {
	// 	version.Version()
	// 	return
	// }

	// //init conf
	// confutil.InitConfig()

	// defer recovery()
	r := gin.Default()
	// s := ginhttp.NewServer()
	// engine := s.GetGinEngine()

	//Add middleware
	//You can customize the middleware according to your actual needs
	// engine.Use()

	route.RegisterRouter(r)

	r.Run()

	//Front hook for service startup
	// s.AddBeforeServerStartFunc(
	// bs.InitLoggerWithConf(),
	// bs.InitTraceLogger("Gaea", "1.0"),
	// s.InitConfig(),
	// )

	//Exec hook Funcs before the service to closing
	// s.AddAfterServerStopFunc(bs.CloseLogger())

	// er := s.Serve()
	// if er != nil {
	// 	log.Printf("Server stop err:%v", er)
	// } else {
	// 	log.Printf("Server exit")
	// }
}
