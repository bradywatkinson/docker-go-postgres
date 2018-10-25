// customer/rest.go

package customer

import (
  "database/sql"
  "net/http"
  "strconv"
  "fmt"

  "github.com/gorilla/mux"
  "github.com/mholt/binding"

  "app/common"
  "app/test_utils"
)

func postCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("POST /customer:"))

    c := Customer{
      Schema: &CustomerSchema{},
    }

    if err := binding.Bind(r, c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
      return
    }

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
      return
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

    testutils.Log(fmt.Sprintf("Customer: %#v", c.Model))

    if err := binding.Bind(r, c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
      return
    }
    defer r.Body.Close()

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

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Customer{
      Model: &CustomerModel{ID: id},
      Schema: nil,
    }
    if err := c.Model.deleteCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprint("Response:\n{ result: \"success\" }"))

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getCustomers(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /customers"))
    q := &CustomersQuery{}
    if err := binding.Bind(r, q); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
      return
    }

    testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", q.Count, q.Start))

    customers, err := readCustomers(a.DB, int(q.Start), int(q.Count))
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError,   err.Error())
      return
    }

    testutils.Log(fmt.Sprintf("Response:\n%#v", customers))

    common.RespondWithJSON(w, http.StatusOK, customers)
  }
}
