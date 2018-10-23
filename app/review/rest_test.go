// review/routes_test.go

package review_test

import (
  "fmt"
  "os"
  "testing"
  "net/http"
  "encoding/json"
  "bytes"

  "app/common"
  "app/review"
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
  review.InitializeREST(&a)

  code := m.Run()

  testutils.TeardownDB(&a)
  os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
  testutils.ClearTable(&a, "review")

  req, _ := http.NewRequest("GET", "/reviews", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  if body := response.Body.String(); body != "[]" {
    t.Errorf("Expected an empty array. Got %s", body)
  }
}

func TestGetNonExistentReview(t *testing.T) {
  testutils.ClearTable(&a, "review")

  req, _ := http.NewRequest("GET", "/review/11", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusNotFound, response.Code)

  var m map[string]string
  json.Unmarshal(response.Body.Bytes(), &m)
  if m["error"] != "Review not found" {
    t.Errorf("Expected the 'error' key of the response to be set to 'Review not found'. Got '%s'", m["error"])
  }
}

func TestCreateReview(t *testing.T) {
  testutils.ClearTable(&a, "customer")
  testutils.ClearTable(&a, "product")
  testutils.ClearTable(&a, "review")

  testutils.AddCustomers(&a, 1)
  testutils.AddProducts(&a, 1)

  payload := []byte(`{"rating":4,"review":"this is a review","customer_id":1,"product_id":1}`)

  req, _ := http.NewRequest("POST", "/review", bytes.NewBuffer(payload))
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusCreated, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["rating"] != 4.0 {
    t.Errorf("Expected review.rating to be '4'. Got '%v'", m["rating"])
  }

  if m["review"] != "this is a review" {
    t.Errorf("Expected review.review to be 'this is a review'. Got '%v'", m["review"])
  }

  if m["customer_id"] != 1.0 {
    t.Errorf("Expected review.customer_id to be '1'. Got '%v'", m["customer_id"])
  }

  if m["product_id"] != 1.0 {
    t.Errorf("Expected review.product_id to be '1'. Got '%v'", m["product_id"])
  }

  // the id is compared to 1.0 because JSON unmarshaling converts numbers to
  // floats, when the target is a map[string]interface{}
  if m["id"] != 1.0 {
    t.Errorf("Expected review ID to be '1'. Got '%v'", m["id"])
  }
}

func TestGetReview(t *testing.T) {
  testutils.ClearTable(&a, "review")
  _, err := testutils.AddReviews(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/review/1", nil)
  response := testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateReview(t *testing.T) {
  testutils.ClearTable(&a, "review")
  _, err := testutils.AddReviews(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/review/1", nil)
  response := testutils.ExecuteRequest(&a, req)
  var originalReview map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &originalReview)

  payload := []byte(`{"rating":3,"review":"this is a review - updated"}`)

  req, _ = http.NewRequest("PUT", "/review/1", bytes.NewBuffer(payload))
  response = testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  var m map[string]interface{}
  json.Unmarshal(response.Body.Bytes(), &m)

  if m["id"] != originalReview["id"] {
    t.Errorf("Expected the id to remain the same (%v). Got %v", originalReview["id"], m["id"])
  }

  if m["rating"] == originalReview["rating"] {
    t.Errorf("Expected the rating to change from '%v' to '%v'. Got '%v'", originalReview["name"], m["name"], m["name"])
  }

  if m["review"] == originalReview["review"] {
    t.Errorf("Expected the review to change from '%v' to '%v'. Got '%v'", originalReview["price"], m["price"], m["price"])
  }
}

func TestDeleteReview(t *testing.T) {
  testutils.ClearTable(&a, "review")
  _, err := testutils.AddReviews(&a, 1)
  if err != nil {
    t.Errorf(err.Error())
  }

  req, _ := http.NewRequest("GET", "/review/1", nil)
  response := testutils.ExecuteRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("DELETE", "/review/1", nil)
  response = testutils.ExecuteRequest(&a, req)

  testutils.CheckResponseCode(t, http.StatusOK, response.Code)

  req, _ = http.NewRequest("GET", "/review/1", nil)
  response = testutils.ExecuteRequest(&a, req)
  testutils.CheckResponseCode(t, http.StatusNotFound, response.Code)
}
