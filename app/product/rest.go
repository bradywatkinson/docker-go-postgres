// product/rest.go

package product

import (
  "net/http"
  "strconv"
  "fmt"

  "github.com/gorilla/mux"
  "github.com/mholt/binding"
  "github.com/jinzhu/gorm"

  "app/common"
)

func postProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {

    p := Product{
      Schema: &ProductSchema{},
    }

    if err := binding.Bind(req, p.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    p.copySchema()

    if err := p.Model.createProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    p.copyModel()

    common.RespondWithJSON(w, http.StatusCreated, p.Schema)
  }
}

func getProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
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
      case gorm.ErrRecordNotFound:
        common.RespondWithError(w, http.StatusNotFound, "Product not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    p.copyModel()

    common.RespondWithJSON(w, http.StatusOK, p.Schema)
  }
}

func putProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
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
      case gorm.ErrRecordNotFound:
        common.RespondWithError(w, http.StatusNotFound, "Product not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

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

    common.RespondWithJSON(w, http.StatusOK, p.Schema)
  }
}

func deleteProduct(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
      return
    }

    p := Product{
      Model: &ProductModel{ID: id},
      Schema: nil,
    }
    if err := p.Model.deleteProduct(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getProducts(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    q := &ProductsQuery{}
    if err := binding.Bind(req, q); err != nil {
      common.RespondWithError(w, http.StatusBadRequest,fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    products, err := readProducts(a.DB, int(q.Start), int(q.Count))
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, products)
  }
}
