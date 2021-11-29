package DB

import (
	"CrudForLandScape/Models"
	"fmt"
)

func ReadUser(id string) ([]Models.User, error) {
	var users []Models.User
	var user Models.User

	where := ""

	if id != "" {
		where = "where id=" + id
	}

	registry, err := ConDB.Query("SELECT * FROM usuarios " + where)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for registry.Next() {
		var iddb int
		var name, email string
		err := registry.Scan(&iddb, &name, &email)
		if err != nil {
			fmt.Println(err.Error())
		}

		user.ID = iddb
		user.EMAIL = email
		user.NAME = name

		users = append(users, user)
	}

	return users, nil

}
