// customer/routes_test.go

package customer_test

import (
  "fmt"
  "os"
  "testing"
  "net/http"
  "encoding/json"
  "bytes"

  "app/common"
  "app/customer"
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
  customer.InitializeRoutes(&a)

  code := m.Run()

  testutils.TeardownDB(&a)
  os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
  testutils.ClearTable(&a, "customer")

  req, _ := http.NewRequest("GET", "/customers", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func TestGetNonExistentcustomer(t *testing.T) {
  testutils.ClearTable(&a, "customer")

  req, _ := http.NewRequest("GET", "/customer/11", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusNotFound, response.Code)

  var m map[string]string
  json.Unmarshal(response.Body.Bytes(), &m)
  if m["error"] != "Customer not found" {
    t.Errorf("Expected the 'error' key of the response to be set to 'Customer not found'. Got '%s'", m["error"])
  }
}

func TestCreateCustomer(t *testing.T) {
  testutils.ClearTable(&a, "customer")

  payload := []byte(`{"first_name":"test first name","last_name":"test last name"}`)

  req, _ := http.NewRequest("POST", "/customer", bytes.NewBuffer(payload))
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusCreated, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["first_name"] != "test first name" {
    t.Errorf("Expected customer name to be 'test first customer'. Got '%v'", m["first_name"])
  }

  if m["last_name"] != "test last name" {
    t.Errorf("Expected customer first_name to be 'test last name'. Got '%v'", m["last_name"])
  }

  // the id is compared to 1.0 because JSON unmarshaling converts numbers to
  // floats, when the target is a map[string]interface{}
  if m["id"] != 1.0 {
    t.Errorf("Expected customer ID to be '1'. Got '%v'", m["id"])
  }
}

func TestGetCustomer(t *testing.T) {
  testutils.ClearTable(&a, "customer")
  _, err := testutils.AddCustomers(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/customer/1", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateCustomer(t *testing.T) {
  testutils.ClearTable(&a, "customer")
  _, err := testutils.AddCustomers(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/customer/1", nil)
  response := testutils.ExecuteRequest(&a, req)
  var originalcustomer map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &originalcustomer)

  payload := []byte(`{"first_name":"test customer - updated name"}`)

  req, _ = http.NewRequest("PUT", "/customer/1", bytes.NewBuffer(payload))
  response = testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["id"] != originalcustomer["id"] {
    t.Errorf("Expected the id to remain the same (%v). Got %v", originalcustomer["id"], m["id"])
  }

  if m["first_name"] == originalcustomer["first_name"] {
    t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalcustomer["first_name"], m["first_name"], m["first_name"])
  }
}

func TestDeleteCustomer(t *testing.T) {
  testutils.ClearTable(&a, "customer")
  _, err := testutils.AddCustomers(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/customer/1", nil)
  response := testutils.ExecuteRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("DELETE", "/customer/1", nil)
  response = testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("GET", "/customer/1", nil)
  response = testutils.ExecuteRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusNotFound, response.Code)
}

