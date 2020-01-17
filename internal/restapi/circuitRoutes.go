package restapi

import (
	"fmt"
	"net/http"
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
	}
	fmt.Fprintf(w, response)
}
