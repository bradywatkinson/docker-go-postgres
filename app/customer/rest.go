// customer/rest.go

package customer

import (
  "database/sql"
  "net/http"
  "strconv"
  "encoding/json"
  "fmt"

  "github.com/gorilla/mux"

  "app/common"
  "app/test_utils"
)

func postCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("POST /customer:"))

    c := Customer{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
      return
    }
    defer r.Body.Close()

    testutils.Log(fmt.Sprintf("%#v", c.Schema))

    c.copySchema()

    if err := c.Model.createCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
    common.RespondWithJSON(w, http.StatusCreated, c.Schema)
  }
}

func getCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /customer:"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Customer{
      Model: &CustomerModel{ID: id},
      Schema: &CustomerSchema{},
    }
    if err := c.Model.readCustomer(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Customer not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

    common.RespondWithJSON(w, http.StatusOK, c.Schema)
  }
}

func putCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("PUT /customer"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
      return
    }

    c := Customer{
      Model: &CustomerModel{ID: id},
      Schema: nil,
    }

    if err := c.Model.readCustomer(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Customer not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
      return
    }
    defer r.Body.Close()

    testutils.Log(fmt.Sprintf("%#v", c.Schema))

    c.copySchema()

    if err := c.Model.updateCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

    common.RespondWithJSON(w, http.StatusOK, c.Schema)
  }
}

func deleteCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("DELETE /customer"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
      return
    }

    c := Customer{
      Model: &CustomerModel{ID: id},
      Schema: nil,
    }
    if err := c.Model.deleteCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }
    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    testutils.Log(fmt.Sprint("Response:\n{ result: \"success\""))

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getCustomers(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /customers"))
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
      count = 10
    }
    if start < 0 {
      start = 0
    }

    testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

    customers, err := readCustomers(a.DB, start, count)
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprintf("Response:\n%#v", customers))

    common.RespondWithJSON(w, http.StatusOK, customers)
  }
}
