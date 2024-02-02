package main

import (
	"bitbucket.org/juancolamendydevteam/golib/utils/logger"
	"bitbucket.org/juancolamendydevteam/prod-apisvr-starter/conf"
	"bitbucket.org/juancolamendydevteam/prod-apisvr-starter/route"
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
)

func main() {
	// get application config
	config := conf.GetConfig()
	config.Dump()

	// init HTTP server
	rtr := route.GetRouter()

	// create an HTTP server and start listening
	logger.Info.Printf("--- Start HTTP server. HTTP_PORT:[%s]\n", config.HTTP_PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", config.HTTP_PORT),
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(rtr))
}
