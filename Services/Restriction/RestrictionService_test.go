package RestrictionService_test

import (
	m "system_backend_go/Models"
	restrictionService "system_backend_go/Services/Restriction"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	restrictions, err := restrictionService.All()
	if err != nil {
		t.Error("Testing Of All Restrictions Has Failing")
		t.Fail()
	}

	if len(restrictions) == 0 {
		t.Error("Testing Of All Restrictions not return Data")
		t.Fail()
	} else {
		t.Log("Testing Of All Restrictions Was Succesful")
	}
}
func TestStore(t *testing.T) {
	restriction := m.Restriction{
		Name:        "Alana",
		Description: "Usuga Sepulveda",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := restrictionService.Store(restriction)
	if err != nil {
		t.Error("Testing Of Create Restriction Has Failing")
		t.Fail()
	} else {
		t.Log("Testing Of Create Restriction Was Succesful")
	}
}
func TestUpdate(t *testing.T) {

}
func TestDestroy(t *testing.T) {

}
