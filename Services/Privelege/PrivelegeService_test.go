package PrivelegeService_test

import (
	m "system_backend_go/Models"
	privelegeService "system_backend_go/Services/Privelege"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	priveleges, err := privelegeService.All()
	if err != nil {
		t.Error("Testing Of All Privileges Has Failing")
		t.Fail()
	}

	if len(priveleges) == 0 {
		t.Error("Testing Of All Privileges not return Data")
		t.Fail()
	} else {
		t.Log("Testing Of All Privileges Was Succesful")
	}
}
func TestStore(t *testing.T) {
	privelege := m.Privelege{
		Name:        "Alana",
		Description: "Usuga Sepulveda",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := privelegeService.Store(privelege)
	if err != nil {
		t.Error("Testing Of Create Privilege Has Failing")
		t.Fail()
	} else {
		t.Log("Testing Of Create Privilege Was Succesful")
	}
}
func TestUpdate(t *testing.T) {
	privelege := m.Privelege{
		Name:        "Alana",
		Description: "Usuga Sepulveda",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := privelegeService.Update(privelege, 1)
	if err != nil {
		t.Error("Testing Of Update Privilege Has Failing")
		t.Fail()
	} else {
		t.Log("Testing Of Update Privilege Was Succesful")
	}
}
func TestDestroy(t *testing.T) {

	err := privelegeService.Destroy(1)
	if err != nil {
		t.Error("Testing Of Delete Privilege Has Failing")
		t.Fail()
	} else {
		t.Log("Testing Of Delete Privilege Was Succesful")
	}
}
