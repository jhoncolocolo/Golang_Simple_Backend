package PrivelegeService

import (
	PrivelegeRepository "system_backend_go/Repositories/Privelege"

	m "system_backend_go/Models"
)

func All() (m.Priveleges, error) {
	priveleges, err := PrivelegeRepository.All()
	if err != nil {
		return nil, err
	}
	return priveleges, nil
}

func Find(privelegeId string) (m.Privelege, error) {
	privelege, err := PrivelegeRepository.Find(privelegeId)
	if err != nil {
		return privelege, err
	}
	return privelege, nil
}

func Store(privelege m.Privelege) error {
	err := PrivelegeRepository.Store(privelege)
	if err != nil {
		return err
	}
	return nil
}

func Update(privelege m.Privelege, privelegeId string) error {
	err := PrivelegeRepository.Update(privelege, privelegeId)
	if err != nil {
		return err
	}
	return nil
}

func Destroy(privelegeId string) error {
	err := PrivelegeRepository.Destroy(privelegeId)
	if err != nil {
		return err
	}
	return nil
}
