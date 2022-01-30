package Models

type Planta struct {
	ID         int    `json:"id"`
	NAME       string `json:"name"`
	TIPO       string `json:"tipo"`
	HORA_RIEGO string `json:"horariego"`
	HORA_LUZ   string `json:"horaluz"`
}
