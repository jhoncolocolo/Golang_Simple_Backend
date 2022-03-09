package Controllers

import (
	"net/http"
	h "system_backend_go/Helpers"
	m "system_backend_go/Models"
	restrictionService "system_backend_go/Services/Restriction"
	"time"
)

func IndexRestriction(w http.ResponseWriter, r *http.Request) {
	restrictions, err := restrictionService.All()
	if err != nil {
		panic(err.Error())
	}

	if len(restrictions) == 0 {
		panic(err.Error())
	}

	h.ParseHtmlTemplates().ExecuteTemplate(w, "restriction_index", restrictions) //Call te layout*/
}

func CreateRestriction(w http.ResponseWriter, r *http.Request) {
	h.ParseHtmlTemplates().ExecuteTemplate(w, "restriction_create", nil) //Call te layout
}

func StoreRestriction(w http.ResponseWriter, r *http.Request) {
	restriction := m.Restriction{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := restrictionService.Store(restriction)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, "/restrictions", 301)
}

func EditRestriction(w http.ResponseWriter, r *http.Request) {
	registryId := r.URL.Query().Get("id")
	restriction, err := restrictionService.Find(registryId)
	if err != nil {
		panic(err.Error())
	}
	h.ParseHtmlTemplates().ExecuteTemplate(w, "restriction_edit", restriction) //Call te layout
}

func UpdateRestriction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		restriction := m.Restriction{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		id := r.FormValue("id")
		err := restrictionService.Update(restriction, id)
		if err != nil {
			panic(err.Error())
		}
		http.Redirect(w, r, "/restrictions", 301)
	}
}
func DestroyRestriction(w http.ResponseWriter, r *http.Request) {
	restrictionId := r.URL.Query().Get("id")
	err := restrictionService.Destroy(restrictionId)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, "/restrictions", 301)
}