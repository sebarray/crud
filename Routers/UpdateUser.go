package Routers

import (
	"CrudForLandScape/DB"
	"CrudForLandScape/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user Models.User
	reqbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintln(w, "error receiving client data")
		return
	}

	json.Unmarshal(reqbody, &user)

	err = DB.UpdateUser(user, vars["ID"])

	if err != nil {
		http.Error(w, "the user with that id does not exist"+err.Error(), http.StatusBadRequest)
		return
	}

	user.ID, err = strconv.Atoi(vars["ID"])
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}
