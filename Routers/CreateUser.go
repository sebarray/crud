package Routers

import (
	"CrudForLandScape/DB"
	"CrudForLandScape/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var rPlant Models.Relacion
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, "error receiving client data")
		return
	}
	idplant := vars["hora"]

	json.Unmarshal(reqbody, &rPlant)

	plant, err2 := DB.CreateUser(rPlant, idplant)

	if err2 != nil {
		http.Error(w, "problems generating the user in the db"+err2.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(plant)
}
