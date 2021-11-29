package DB

import (
	"CrudForLandScape/Models"
)

func CreateUser(user Models.User) error {
	query := " INSERT INTO usuarios (NAME, EMAIL) VALUES (\"" + user.NAME + "\", \"" + user.EMAIL + "\")"
	create, err := ConDB.Prepare(query)
	if err != nil {
		return err
	}
	create.Exec()
	return nil
}
