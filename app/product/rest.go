// product/rest.go

package product

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

func postProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    testutils.Log(fmt.Sprint("POST /product:"))

    p := Product{
      Schema: &ProductSchema{},
    }

    if err := binding.Bind(req, p.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    testutils.Log(fmt.Sprintf("Product: %#v", p.Schema))

    p.copySchema()

    if err := p.Model.createProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    p.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", p.Schema))
    common.RespondWithJSON(w, http.StatusCreated, p.Schema)
  }
}

func getProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    testutils.Log(fmt.Sprint("GET /product:"))
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
      return
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    p := Product{
      Model: &ProductModel{ID: id},
      Schema: &ProductSchema{},
    }
    if err := p.Model.readProduct(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Product not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    p.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", p.Schema))

    common.RespondWithJSON(w, http.StatusOK, p.Schema)
  }
}

func putProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    testutils.Log(fmt.Sprint("PUT /product"))
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
      return
    }

    p := Product{
      Model: &ProductModel{ID: id},
      Schema: &ProductSchema{},
    }

    if err := p.Model.readProduct(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Product not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    testutils.Log(fmt.Sprintf("Product: %#v", p.Model))

    if err := binding.Bind(req, p.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }
    defer req.Body.Close()

    p.copySchema()

    if err := p.Model.updateProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    p.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", p.Schema))

    common.RespondWithJSON(w, http.StatusOK, p.Schema)
  }
}

func deleteProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    testutils.Log(fmt.Sprint("DELETE /product"))
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
      return
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    p := Product{
      Model: &ProductModel{ID: id},
      Schema: nil,
    }
    if err := p.Model.deleteProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprint("Response:\n{ result: \"success\" }"))

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getProducts(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    testutils.Log(fmt.Sprint("GET /products"))
    q := &ProductsQuery{}
    if err := binding.Bind(req, q); err != nil {
      common.RespondWithError(w, http.StatusBadRequest,fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", q.Count, q.Start))

    products, err := readProducts(a.DB, int(q.Start), int(q.Count))
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprintf("Response:\n%#v", products))

    common.RespondWithJSON(w, http.StatusOK, products)
  }
}
