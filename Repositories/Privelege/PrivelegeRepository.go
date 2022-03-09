package PrivelegeRepository

import (
	m "system_backend_go/Models"
	connection "system_backend_go/connections"
)

var database = connection.ConnectMySQLDatabase()

func All() (m.Priveleges, error) {

	var priveleges m.Priveleges
	registries, err := database.Query("SELECT id,name,description FROM priveleges")

	if err != nil {
		return nil, err
	}

	for registries.Next() {
		var privelege m.Privelege
		err := registries.Scan(&privelege.Id, &privelege.Name, &privelege.Description)
		if err != nil {
			panic(err.Error())
		}
		priveleges = append(priveleges, privelege)
	}
	return priveleges, nil
}

func Find(privilegeId string) (m.Privelege, error) {
	var privelege m.Privelege
	err := database.QueryRow("SELECT id,name,description FROM priveleges WHERE id = ?", privilegeId).Scan(&privelege.Id, &privelege.Name, &privelege.Description)
	if err != nil {
		return privelege, err
	} else {
		return privelege, nil
	}
}

func Store(privilege m.Privelege) error {
	var err error

	insertRecords, err := database.Prepare("INSERT INTO priveleges(name,description,created_at) VALUES(?,?,?)")

	if err != nil {
		return err
	}

	insertRecords.Exec(privilege.Name, privilege.Description, privilege.CreatedAt)
	return nil
}

func Update(privelege m.Privelege, privelegeId string) error {
	var err error
	updateRecords, err := database.Prepare("UPDATE priveleges SET name=?,description=?,updated_at=? WHERE id=?")

	if err != nil {
		return err
	}
	updateRecords.Exec(privelege.Name, privelege.Description, privelege.UpdatedAt, privelegeId)
	return nil
}
func Destroy(privelegeId string) error {
	var err error
	deleteRecord, err := database.Prepare("DELETE FROM priveleges WHERE id = ?")

	if err != nil {
		return err
	}
	deleteRecord.Exec(privelegeId)
	return nil
}
