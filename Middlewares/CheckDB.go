package Middlewares

import (
	"CrudForLandScape/DB"
	"net/http"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if DB.Checkconnection() == false {
			http.Error(w, "Lost connection with the database", 500)
return
				}
		next.ServeHTTP(w, r)
	}
}
