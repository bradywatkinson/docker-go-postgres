// router_utils.go

package testutils

import (
  "testing"
  "net/http"
  "net/http/httptest"

  "app/common"
)


// ExecuteRequest is a utility function to send a request
// to the router
func ExecuteRequest(a *common.App, req *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()
  a.Router.ServeHTTP(rr, req)

  return rr
}

// CheckResponseCode is a utility function to check
// the response from the router
func CheckResponseCode(t *testing.T, expected, actual int) {
  if expected != actual {
    t.Errorf("Expected response code %d. Got %d\n", expected, actual)
  }
}
