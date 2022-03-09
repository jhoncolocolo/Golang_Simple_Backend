package Controllers

import (
	"net/http"
	h "system_backend_go/Helpers"
	m "system_backend_go/Models"
	privelegeService "system_backend_go/Services/Privelege"
	"time"
)

var views = h.ParseHtmlTemplates()

func IndexPrivelege(w http.ResponseWriter, r *http.Request) {
	priveleges, err := privelegeService.All()
	if err != nil {
		panic(err.Error())
	}

	if len(priveleges) == 0 {
		panic(err.Error())
	}

	views.ExecuteTemplate(w, "privelege_index", priveleges) //Call te layout*/
}

func CreatePrivelege(w http.ResponseWriter, r *http.Request) {
	views.ExecuteTemplate(w, "privelege_create", nil) //Call te layout
}

func StorePrivelege(w http.ResponseWriter, r *http.Request) {
	privelege := m.Privelege{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := privelegeService.Store(privelege)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, "/priveleges", 301)
}

func EditPrivelege(w http.ResponseWriter, r *http.Request) {
	registryId := r.URL.Query().Get("id")
	privelege, err := privelegeService.Find(registryId)
	if err != nil {
		panic(err.Error())
	}
	views.ExecuteTemplate(w, "privelege_edit", privelege) //Call te layout
}

func UpdatePrivelege(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		privelege := m.Privelege{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		id := r.FormValue("id")
		err := privelegeService.Update(privelege, id)
		if err != nil {
			panic(err.Error())
		}
		http.Redirect(w, r, "/priveleges", 301)
	}
}
func DestroyPrivelege(w http.ResponseWriter, r *http.Request) {
	privelegeId := r.URL.Query().Get("id")
	err := privelegeService.Destroy(privelegeId)
	if err != nil {
		panic(err.Error())
	}
	http.Redirect(w, r, "/priveleges", 301)
}