package RestrictionRepository

import (
	m "system_backend_go/Models"
	connection "system_backend_go/connections"
)

var database = connection.ConnectMySQLDatabase()

func All() (m.Restrictions, error) {

	var restrictions m.Restrictions
	registries, err := database.Query("SELECT * FROM restrictions")

	if err != nil {
		return nil, err
	}

	for registries.Next() {
		var restriction m.Restriction
		err := registries.Scan(&restriction.Id, &restriction.Name, &restriction.Description)
		if err != nil {
			panic(err.Error())
		}
		restrictions = append(restrictions, restriction)
	}
	return restrictions, nil
}

func Find(restrictionId string) (m.Restriction, error) {
	var restriction m.Restriction
	err := database.QueryRow("SELECT id,name,description FROM restrictions WHERE id = ?", restrictionId).Scan(&restriction.Id, &restriction.Name, &restriction.Description)
	if err != nil {
		return restriction, err
	} else {
		return restriction, nil
	}
}

func Store(restriction m.Restriction) error {
	var err error

	insertRecords, err := database.Prepare("INSERT INTO restrictions(name,description) VALUES(?,?)")

	if err != nil {
		return err
	}

	insertRecords.Exec(restriction.Name, restriction.Description)
	return nil
}

func Update(restriction m.Restriction, restrictionId string) error {
	var err error
	updateRecords, err := database.Prepare("UPDATE restrictions SET name=?,description=? WHERE id=?")

	if err != nil {
		return err
	}
	updateRecords.Exec(restriction.Name, restriction.Description, restrictionId)
	return nil
}
func Destroy(restrictionId string) error {
	var err error
	deleteRecord, err := database.Prepare("DELETE FROM restrictions WHERE id = ?")

	if err != nil {
		return err
	}
	deleteRecord.Exec(restrictionId)
	return nil
}
