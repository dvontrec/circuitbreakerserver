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
	response := fmt.Sprintf("Attempts = %d", api.attemptsSinceFail)
	if api.attemptsSinceFail >= 10 {
		api.attemptsSincePass++
		response = fmt.Sprintf("Server crashed\nattempts passed = %d\nattempts failed = %d", api.attemptsSinceFail, api.attemptsSincePass)
		if api.attemptsSincePass >= 13 {
			backUp := checkForRecovery()
			if backUp {
				api.attemptsSinceFail = 0
				api.attemptsSincePass = 0
			}
		}
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
