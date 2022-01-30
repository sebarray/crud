package DB

import (
	"CrudForLandScape/Models"
)

func CreateUser(user Models.Relacion, idplant string) (Models.Planta, error) {

	var planta Models.Planta

	//query := " INSERT INTO usuarios (NAME, EMAIL) VALUES (\"" + user.IdUser + "\", \"" + user.TIPO + "\")"
	query := idplant
	create, err := ConDB.Prepare(query)
	if err != nil {
		return planta, err
	}
	create.Exec()
	return planta, nil
}
