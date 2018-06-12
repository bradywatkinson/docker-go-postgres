// merchant/routes_test.go

package merchant_test

import (
  "fmt"
  "os"
  "testing"
  "net/http"
  "encoding/json"
  "bytes"

  "app/common"
  "app/merchant"
  "app/test_utils"
)

var a common.App

func TestMain(m *testing.M) {
  a = common.App{}
  connectionString :=
    fmt.Sprintf("sslmode=disable host=postgres user=%s password=%s dbname=%s",
      os.Getenv("APP_DB_USERNAME"),
      os.Getenv("APP_DB_PASSWORD"),
      os.Getenv("APP_DB_NAME"))
  a.InitializeDB(connectionString)

  testutils.SetupDB(&a)

  a.InitializeRouter()
  merchant.InitializeRoutes(&a)

  code := m.Run()

  testutils.TeardownDB(&a)
  os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
  testutils.ClearTable(&a, "merchant")

  req, _ := http.NewRequest("GET", "/merchants", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func TestGetNonExistentMerchant(t *testing.T) {
  testutils.ClearTable(&a, "merchant")

  req, _ := http.NewRequest("GET", "/merchant/11", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusNotFound, response.Code)

  var m map[string]string
  json.Unmarshal(response.Body.Bytes(), &m)
  if m["error"] != "Merchant not found" {
    t.Errorf("Expected the 'error' key of the response to be set to 'Merchant not found'. Got '%s'", m["error"])
  }
}

func TestCreateMerchant(t *testing.T) {
  testutils.ClearTable(&a, "merchant")

  payload := []byte(`{"name":"test merchant"}`)

  req, _ := http.NewRequest("POST", "/merchant", bytes.NewBuffer(payload))
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusCreated, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["name"] != "test merchant" {
    t.Errorf("Expected merchant name to be 'test merchant'. Got '%v'", m["name"])
  }

  // the id is compared to 1.0 because JSON unmarshaling converts numbers to
  // floats, when the target is a map[string]interface{}
  if m["id"] != 1.0 {
    t.Errorf("Expected merchant ID to be '1'. Got '%v'", m["id"])
  }
}

func TestGetMerchant(t *testing.T) {
  testutils.ClearTable(&a, "merchant")
  _, err := testutils.AddMerchants(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/merchant/1", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateMerchant(t *testing.T) {
  testutils.ClearTable(&a, "merchant")
  _, err := testutils.AddMerchants(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/merchant/1", nil)
  response := testutils.ExecuteRequest(&a, req)
  var originalMerchant map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &originalMerchant)

  payload := []byte(`{"name":"test merchant - updated name"}`)

  req, _ = http.NewRequest("PUT", "/merchant/1", bytes.NewBuffer(payload))
  response = testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["id"] != originalMerchant["id"] {
    t.Errorf("Expected the id to remain the same (%v). Got %v", originalMerchant["id"], m["id"])
  }

  if m["name"] == originalMerchant["name"] {
    t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalMerchant["name"], m["name"], m["name"])
  }
}

func TestDeleteMerchant(t *testing.T) {
  testutils.ClearTable(&a, "merchant")
  _, err := testutils.AddMerchants(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/merchant/1", nil)
  response := testutils.ExecuteRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("DELETE", "/merchant/1", nil)
  response = testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("GET", "/merchant/1", nil)
  response = testutils.ExecuteRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusNotFound, response.Code)
}
