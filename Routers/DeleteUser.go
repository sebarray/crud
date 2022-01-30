package Routers

// import (
// 	"CrudForLandScape/DB"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	err := DB.DeleteUser(vars["ID"])

// 	if err != nil {
// 		http.Error(w, "Error to delete user "+err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	users, _ := DB.ReadUser("")

// 	w.Header().Set("content-type", "application/json")
// 	json.NewEncoder(w).Encode(users)

// }
