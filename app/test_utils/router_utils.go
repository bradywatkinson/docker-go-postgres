// router_utils.go

package testutils

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "encoding/json"

  "app/common"
)


// ExecuteJSONRequest is a utility function to send a request
// to the router
func ExecuteJSONRequest(a *common.App, req *http.Request) *httptest.ResponseRecorder {
  req.Header.Set("Content-Type", "application/json")

  return ExecuteRequest(a, req)
}

func ExecuteRequest(a *common.App, req *http.Request) *httptest.ResponseRecorder {
  rr := httptest.NewRecorder()
  a.Router.ServeHTTP(rr, req)

  return rr
}

// CheckResponseCode is a utility function to check
// the response from the router
func CheckResponseCode(t *testing.T, expected int, response *httptest.ResponseRecorder) {
  if expected != response.Code {
    var r map[string]interface{}
    json.Unmarshal(response.Body.Bytes(), &r)
    t.Errorf("Expected response code %d. Got %d\nResponse:%#v",
      expected,
      response.Code,
      r,
    )
  }
}
