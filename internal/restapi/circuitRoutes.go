package restapi

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// GetData returns simple data from the restapi
func (api *RESTAPI) GetData(w http.ResponseWriter, r *http.Request) {
	if api.attemptsSincePass == 0 {
		api.attemptsSinceFail++
	}
	if api.attemptsSinceFail >= 30 {
		api.running = false
		api.attemptsSincePass++
		fmt.Println("Server Closed")
		if api.attemptsSincePass >= 21 {
			backUp := checkForRecovery()
			if backUp {
				fmt.Println("Server Opened")
				api.running = true
				api.attemptsSinceFail = 0
				api.attemptsSincePass = 0
			}
		}
	}
	response := fmt.Sprintf("Server running: %v\nattempts passed = %d\nattempts failed = %d", api.running, api.attemptsSinceFail, api.attemptsSincePass)
	if !api.running {
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, response)
}

func checkForRecovery() bool {
	rand.Seed(time.Now().UnixNano())
	rNum := rand.Intn(21)
	if rNum >= 17 {
		return true
	}
	return false
}
