package Routers

import (
	"CrudForLandScape/DB"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ReadUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	users, err := DB.ReadUser(vars["ID"])

	if err != nil {
		http.Error(w, "the user with that id does not exist"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func ReadUsersAll(w http.ResponseWriter, r *http.Request) {
	users, err := DB.ReadUser("")
	if err != nil {
		http.Error(w, "the user with that id does not exist"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(users)

}
