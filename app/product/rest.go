// product/rest.go

package product

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

func postProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("POST /product:"))

    c := Product{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
      return
    }
    defer r.Body.Close()

    testutils.Log(fmt.Sprintf("%#v", c.Schema))

    c.copySchema()

    if err := c.Model.createProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
    common.RespondWithJSON(w, http.StatusCreated, c.Schema)
  }
}

func getProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /product:"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Product{
      Model: &ProductModel{ID: id},
      Schema: &ProductSchema{},
    }
    if err := c.Model.readProduct(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Product not found")
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

func putProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("PUT /product"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
      return
    }

    c := Product{
      Model: &ProductModel{ID: id},
      Schema: nil,
    }

    if err := c.Model.readProduct(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Product not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
      return
    }
    defer r.Body.Close()

    c.copySchema()

    if err := c.Model.updateProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

    common.RespondWithJSON(w, http.StatusOK, c.Schema)
  }
}

func deleteProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("DELETE /product"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
      return
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Product{
      Model: &ProductModel{ID: id},
      Schema: nil,
    }
    if err := c.Model.deleteProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprint("Response:\n{ result: \"success\" }"))

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getProducts(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /products"))
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
      count = 10
    }
    if start < 0 {
      start = 0
    }

    testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

    products, err := readProducts(a.DB, start, count)
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprintf("Response:\n%#v", products))

    common.RespondWithJSON(w, http.StatusOK, products)
  }
}
