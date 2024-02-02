package route

import (

	"bitbucket.org/juancolamendydevteam/prod-apisvr-starter/controller/indexctl"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	// Create router
	rtr := mux.NewRouter()

	// index_controller
	rtr.HandleFunc("/", indexctl.GetIndexPage).Methods("GET")

	// Return router
	return rtr
}
