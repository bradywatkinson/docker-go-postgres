// common/router.go

package common

import (
  "fmt"
  "net/http"
  "encoding/json"
  "bytes"

  "github.com/gorilla/mux"
  "github.com/urfave/negroni"
)

type myResponseWriter struct {
  http.ResponseWriter
  buf *bytes.Buffer
}

func (mrw *myResponseWriter) Write(p []byte) (int, error) {
  return mrw.buf.Write(p)
}

type JsonPanicFormatter struct{}

func (j *JsonPanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *negroni.PanicInformation) {
  e := map[string]string{
    "error": fmt.Sprintf("%s", infos.RecoveredPanic),
  }
  response, _ := json.Marshal(e)
  rw.Write(response)
}

// InitializeRouter setups the mux router and any additional
// middlewares
func (a *App) InitializeRouter() *negroni.Negroni {
  a.Router = mux.NewRouter()
  n := negroni.New()
  n.Use(&loggingMiddleware{})
  recovery := negroni.NewRecovery()
  recovery.Formatter = &JsonPanicFormatter{}
  n.Use(recovery)
  return n
}


// RespondWithError is a helper function to respond to a request with an error
func RespondWithError(rw http.ResponseWriter, code int, message string) {
  RespondWithJSON(rw, code, map[string]string{"error": message})
}

// RespondWithJSON is a helper function to respond to a request with a json payload
func RespondWithJSON(rw http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  rw.Header().Set("Content-Type", "application/json")
  rw.WriteHeader(code)
  rw.Write(response)
}
