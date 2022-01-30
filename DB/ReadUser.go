package DB

import (
	"CrudForLandScape/Models"
	"fmt"
)

func ReadUser(id string) ([]Models.Planta, error) {
	var plantas []Models.Planta
	var planta Models.Planta

	where := ""

	if id != "" {
		where = "where id=" + id
	}

	registry, err := ConDB.Query("SELECT * FROM plant " + where)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for registry.Next() {

		err := registry.Scan(&planta.ID, &planta.NAME, &planta.TIPO, &planta.HORA_RIEGO, &planta.HORA_LUZ)
		if err != nil {
			fmt.Println(err.Error())
		}

		plantas = append(plantas, planta)
	}

	return plantas, nil

}
