// product/service.go

package product

import (
  "database/sql"
  "net/http"
  "strconv"
  "encoding/json"
  "fmt"

  "github.com/gorilla/mux"

  "app/common"
)


func createProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    var p product
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&p); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
      return
    }
    defer r.Body.Close()

    if err := p.createProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusCreated, p)
  }
}

func getProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
    }

    p := product{ID: id}
    if err := p.getProduct(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Product not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    common.RespondWithJSON(w, http.StatusOK, p)
  }
}

func updateProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
      return
    }

    var p product
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&p); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
      return
    }
    defer r.Body.Close()
    p.ID = id

    if err := p.updateProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, p)
  }
}

func deleteProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
      return
    }

    p := product{ID: id}
    if err := p.deleteProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getProducts(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
      count = 10
    }
    if start < 0 {
      start = 0
    }

    var p product
    products, err := p.getProducts(a.DB, start, count)
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, products)
  }
}
