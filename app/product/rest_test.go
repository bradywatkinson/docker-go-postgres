// product/routes_test.go

package product_test

import (
  "fmt"
  "os"
  "testing"
  "net/http"
  "encoding/json"
  "bytes"

  "app/common"
  "app/product"
  "app/test_utils"
)

var a common.App

func TestMain(m *testing.M) {
  a = common.App{}
  connectionString :=
    fmt.Sprintf("sslmode=disable host=%s user=%s password=%s dbname=%s",
      os.Getenv("APP_DB_HOST"),
      os.Getenv("APP_DB_USERNAME"),
      os.Getenv("APP_DB_PASSWORD"),
      os.Getenv("APP_DB_NAME"))
  a.InitializeDB(connectionString)

  testutils.SetupDB(&a)

  a.InitializeRouter()
  product.InitializeREST(&a)

  code := m.Run()

  testutils.TeardownDB(&a)
  os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
  testutils.ClearTable(&a, "product")

  req, _ := http.NewRequest("GET", "/products", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func TestGetNonExistentProduct(t *testing.T) {
  testutils.ClearTable(&a, "product")

  req, _ := http.NewRequest("GET", "/product/11", nil)
  response := testutils.ExecuteJSONRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusNotFound, response)

  var m map[string]string
  json.Unmarshal(response.Body.Bytes(), &m)
  if m["error"] != "Product not found" {
    t.Errorf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", m["error"])
  }
}

func TestCreateProduct(t *testing.T) {
  testutils.ClearTable(&a, "product")

  testutils.AddMerchants(&a, 1)

  payload := []byte(`{"name":"test product","price":11.22,"merchant_id": 1}`)

  req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(payload))
  response := testutils.ExecuteJSONRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusCreated, response)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["name"] != "test product" {
    t.Errorf("Expected product name to be 'test product'. Got '%v'", m["name"])
  }

  if m["price"] != 11.22 {
    t.Errorf("Expected product price to be '11.22'. Got '%v'", m["price"])
  }

  // the id is compared to 1.0 because JSON unmarshaling converts numbers to
  // floats, when the target is a map[string]interface{}
  if m["id"] != 1.0 {
    t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
  }
}

func TestGetProduct(t *testing.T) {
  testutils.ClearTable(&a, "product")
  _, err := testutils.AddProducts(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/product/1", nil)
  response := testutils.ExecuteJSONRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response)
}

func TestUpdateProduct(t *testing.T) {
  testutils.ClearTable(&a, "product")
  _, err := testutils.AddProducts(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/product/1", nil)
  response := testutils.ExecuteJSONRequest(&a, req)
  var originalProduct map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &originalProduct)

  payload := []byte(`{"name":"test product - updated name","price":11.22, "merchant_id": 1}`)

  req, _ = http.NewRequest("PUT", "/product/1", bytes.NewBuffer(payload))
  response = testutils.ExecuteJSONRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["id"] != originalProduct["id"] {
    t.Errorf("Expected the id to remain the same (%v). Got %v", originalProduct["id"], m["id"])
  }

  if m["name"] == originalProduct["name"] {
    t.Errorf("Expected the name to change from '%v' to 'test product - updated name'. Got '%v'", originalProduct["name"], m["name"])
  }

  if m["price"] == originalProduct["price"] {
    t.Errorf("Expected the price to change from '%v' to '11.22'. Got '%v'", originalProduct["price"], m["price"])
  }
}

func TestDeleteProduct(t *testing.T) {
  testutils.ClearTable(&a, "product")
  _, err := testutils.AddProducts(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/product/1", nil)
  response := testutils.ExecuteJSONRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusOK, response)

  req, _ = http.NewRequest("DELETE", "/product/1", nil)
  response = testutils.ExecuteJSONRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response)

  req, _ = http.NewRequest("GET", "/product/1", nil)
  response = testutils.ExecuteJSONRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusNotFound, response)
}
