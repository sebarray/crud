package Handlers

import (
	"CrudForLandScape/Routers"
	"net/http"

	"github.com/gorilla/mux"
)

func Manager() {
	router := mux.NewRouter()
	router.HandleFunc("/create", Routers.CreateUser).Methods("POST")
	router.HandleFunc("/readUsers", m).Methods("GET")
	router.HandleFunc("/readUsers/{id}", m).Methods("GET")
	router.HandleFunc("/deleteUser/{id}", m).Methods("DELETE")
	router.HandleFunc("/updateUser/{id}", m).Methods("PUT")

	http.ListenAndServe(":8083", router)
}

func m(w http.ResponseWriter, r *http.Request) {

}
