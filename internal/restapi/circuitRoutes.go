package restapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type response struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

// GetData returns simple data from the restapi
func (api *RESTAPI) GetData(w http.ResponseWriter, r *http.Request) {
	if api.attemptsSincePass == 0 {
		api.attemptsSinceFail++
	}
	if api.attemptsSinceFail >= 30 {
		api.running = false
		api.attemptsSincePass++
		if api.attemptsSincePass >= 21 {
			backUp := checkForRecovery()
			if backUp {
				api.running = true
				api.attemptsSinceFail = 0
				api.attemptsSincePass = 0
			}
		}
	}
	re := fmt.Sprintf("Server running: %v\nattempts passed = %d\nattempts failed = %d", api.running, api.attemptsSinceFail, api.attemptsSincePass)
	res := response{
		Message: re,
		Name:    r.FormValue("value"),
	}
	if !api.running {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Fail val: ", r.FormValue("value"))
		return
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.FormValue("value"))
}

func checkForRecovery() bool {
	rand.Seed(time.Now().UnixNano())
	rNum := rand.Intn(21)
	if rNum >= 17 {
		return true
	}
	return false
}
