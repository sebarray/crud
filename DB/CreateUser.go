package DB

import (
	"CrudForLandScape/Models"
)

func CreateUser(user Models.User) error {

	create, err := ConDB.Prepare(" INSERT INTO usuarios (NAME, EMAIL) VALUES (\"" + user.NAME + "\", \"" + user.EMAIL + "\")")
	if err != nil {
		return err
	}
	create.Exec()
	return nil
}
