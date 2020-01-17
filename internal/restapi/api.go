package restapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RESTAPI allws exported api functionality
type RESTAPI struct {
	attemptsSinceFail int
	attemptsSincePass int
	running           bool
}

// New returns a new instance of the RESTAPI giving access to the CRUD
// opporations
func New() *RESTAPI {
	return &RESTAPI{}
}

// Start starts up the RESTAPI to be consumed
func (api *RESTAPI) Start() {
	router := mux.NewRouter()
	api.running = true
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, from the circuit")
	}).Methods("GET")
	router.HandleFunc("/circuit", api.GetData).Methods("GET")
	http.ListenAndServe(":8109", router)
}
