package DB

func DeleteUser(id string) error {
	query, err := ConDB.Prepare("DELETE FROM usuarios WHERE id=" + id)
	if err != nil {
		return err
	}
	_, err = query.Exec()
	if err != nil {
		return err
	}

	return nil
}
