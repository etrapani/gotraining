package main

import (
	"fmt"
	users "github.com/etrapani/gotraining/back-end/api/user/infraestructure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	log.Println("start api")
	setRoutes()
}

func setRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/", homeLink)
	var userRouter = users.NewUserRouter()
	userRouter.SetRoutes(router.PathPrefix("/api/users").Subrouter())

	log.Println("main")
	log.Fatal(http.ListenAndServe(":8080", router))
}
