package Request

import (
	"fmt"
	"net/http"
	c "system_backend_go/Controllers"
)

func Launch() {
	http.HandleFunc("/priveleges", c.IndexPrivelege)
	http.HandleFunc("/priveleges/create", c.CreatePrivelege)
	http.HandleFunc("/priveleges/edit", c.EditPrivelege)
	http.HandleFunc("/priveleges/store", c.StorePrivelege)
	http.HandleFunc("/priveleges/update", c.UpdatePrivelege)
	http.HandleFunc("/priveleges/destroy", c.DestroyPrivelege)
	http.HandleFunc("/restrictions", c.IndexRestriction)
	http.HandleFunc("/restrictions/create", c.CreateRestriction)
	http.HandleFunc("/restrictions/edit", c.EditRestriction)
	http.HandleFunc("/restrictions/store", c.StoreRestriction)
	http.HandleFunc("/restrictions/update", c.UpdateRestriction)
	http.HandleFunc("/restrictions/destroy", c.DestroyRestriction)
	fmt.Println("Server running")
	http.ListenAndServe(":7777", nil)
}
