package restapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RESTAPI allws exported api functionality
type RESTAPI struct {
	attemptsSincePass int
}

// New returns a new instance of the RESTAPI giving access to the CRUD
// opporations
func New() *RESTAPI {
	return &RESTAPI{}
}

// Start starts up the RESTAPI to be consumed
func (api *RESTAPI) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, from the circuit")
	})
	http.ListenAndServe(":8109", router)
}
