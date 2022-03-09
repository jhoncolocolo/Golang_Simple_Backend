package RestrictionService

import (
	RestrictionRepository "system_backend_go/Repositories/Restriction"

	m "system_backend_go/Models"
)

func All() (m.Restrictions, error) {
	restrictions, err := RestrictionRepository.All()
	if err != nil {
		return nil, err
	}
	return restrictions, nil
}

func Find(restrictionId string) (m.Restriction, error) {
	restriction, err := RestrictionRepository.Find(restrictionId)
	if err != nil {
		return restriction, err
	}
	return restriction, nil
}

func Store(restriction m.Restriction) error {
	err := RestrictionRepository.Store(restriction)
	if err != nil {
		return err
	}
	return nil
}

func Update(restriction m.Restriction, restrictionId string) error {
	err := RestrictionRepository.Update(restriction, restrictionId)
	if err != nil {
		return err
	}
	return nil
}

func Destroy(restrictionId string) error {
	err := RestrictionRepository.Destroy(restrictionId)
	if err != nil {
		return err
	}
	return nil
}
