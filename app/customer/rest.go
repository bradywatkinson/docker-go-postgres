// customer/rest.go

package customer

import (
  "net/http"
  "strconv"
  "fmt"

  "github.com/gorilla/mux"
  "github.com/mholt/binding"
  "github.com/jinzhu/gorm"

  "app/common"
)

func postCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {

    c := Customer{
      Schema: &CustomerSchema{},
    }

    if err := binding.Bind(req, c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    c.copySchema()

    if err := c.Model.createCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    common.RespondWithJSON(w, http.StatusCreated, c.Schema)
  }
}

func getCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
      return
    }

    c := Customer{
      Model: &CustomerModel{ID: int(id)},
      Schema: &CustomerSchema{},
    }
    if err := c.Model.readCustomer(a.DB); err != nil {
      switch err {
      case gorm.ErrRecordNotFound:
        common.RespondWithError(w, http.StatusNotFound, "Customer not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    c.copyModel()

    common.RespondWithJSON(w, http.StatusOK, c.Schema)
  }
}

func putCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
      return
    }

    c := Customer{
      Model: &CustomerModel{ID: int(id)},
      Schema: &CustomerSchema{},
    }

    if err := c.Model.readCustomer(a.DB); err != nil {
      switch err {
      case gorm.ErrRecordNotFound:
        common.RespondWithError(w, http.StatusNotFound, "Customer not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    if err := binding.Bind(req, c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }
    defer req.Body.Close()

    c.copySchema()

    if err := c.Model.updateCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    common.RespondWithJSON(w, http.StatusOK, c.Schema)
  }
}

func deleteCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
      return
    }

    c := Customer{
      Model: &CustomerModel{ID: int(id)},
      Schema: nil,
    }
    if err := c.Model.deleteCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getCustomers(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    q := &CustomersQuery{}
    if err := binding.Bind(req, q); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    customers, err := readCustomers(a.DB, int(q.Start), int(q.Count))
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, customers)
  }
}
