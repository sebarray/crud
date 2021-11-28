package Routers

import (
	"CrudForLandScape/DB"
	"CrudForLandScape/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user Models.User
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "error receiving customer data")
		return
	}
	fmt.Println(user)
	json.Unmarshal(reqbody, &user)

	DB.CreateUser(user)
	//	if err != nil {
	//		http.Error(w, "problems generating the user in the db"+err.Error(), http.StatusBadRequest)
	//		return
	//	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
