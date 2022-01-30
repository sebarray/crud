package Models

import "time"

type Relacion struct {
	IdUser  int       `json:"iduser"`
	IdPlant int    `json:"idplant"`
	HoraUsr time.Time `json:"tipo"`
	
}
