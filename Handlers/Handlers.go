package Handlers

import (
	"CrudForLandScape/Middlewares"
	"CrudForLandScape/Routers"
	"net/http"

	"github.com/gorilla/mux"
)

func Manager() {
	router := mux.NewRouter()
	router.HandleFunc("/create", Middlewares.CheckBD(Routers.CreateUser)).Methods("POST")
	router.HandleFunc("/readall", Middlewares.CheckBD(Routers.ReadUsersAll)).Methods("GET")
	router.HandleFunc("/read/{ID}", Middlewares.CheckBD(Routers.ReadUser)).Methods("GET")
	router.HandleFunc("/delete/{ID}", Middlewares.CheckBD(Routers.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/update/{ID}", Middlewares.CheckBD(Routers.UpdateUser)).Methods("PUT")

	http.ListenAndServe(":8083", router)
}

func m(w http.ResponseWriter, r *http.Request) {

}
