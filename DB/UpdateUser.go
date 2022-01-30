package DB

import "CrudForLandScape/Models"

func UpdateUser(user Models.Planta, id string) error {
	query := "UPDATE usuarios SET email =\"" + user.TIPO + "\", name =\"" + user.NAME + "\" WHERE id =" + id
	update, err := ConDB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = update.Exec()
	if err != nil {
		return err
	}
	return nil
}
