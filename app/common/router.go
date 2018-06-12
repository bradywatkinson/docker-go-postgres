// common/router.go

package common

import (
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
)


// InitializeRouter setups the mux router and any additional
// middlewares
func (a *App) InitializeRouter() {
  a.Router = mux.NewRouter()
}

// RespondWithError is a helper function to respond to a request with an error
func RespondWithError(w http.ResponseWriter, code int, message string) {
  RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON is a helper function to respond to a request with a json payload
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
