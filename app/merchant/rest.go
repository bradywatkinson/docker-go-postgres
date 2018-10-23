// merchant/rest.go

package merchant

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

func postMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("POST /merchant:"))

    c := Merchant{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
      return
    }
    defer r.Body.Close()

    testutils.Log(fmt.Sprintf("%#v", c.Schema))

    c.copySchema()

    if err := c.Model.createMerchant(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
    common.RespondWithJSON(w, http.StatusCreated, c.Schema)
  }
}

func getMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /merchant:"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Merchant ID")
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Merchant{
      Model: &MerchantModel{ID: id},
      Schema: &MerchantSchema{},
    }
    if err := c.Model.readMerchant(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Merchant not found")
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

func putMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("PUT /merchant"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Merchant ID")
      return
    }

    c := Merchant{
      Model: &MerchantModel{ID: id},
      Schema: nil,
    }

    if err := c.Model.readMerchant(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Merchant not found")
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

    c.copySchema()

    if err := c.Model.updateMerchant(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

    common.RespondWithJSON(w, http.StatusOK, c.Schema)
  }
}

func deleteMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("DELETE /merchant"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Merchant ID")
      return
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Merchant{
      Model: &MerchantModel{ID: id},
      Schema: nil,
    }
    if err := c.Model.deleteMerchant(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprint("Response:\n{ result: \"success\" }"))

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getMerchants(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /merchants"))
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
      count = 10
    }
    if start < 0 {
      start = 0
    }

    testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

    merchants, err := readMerchants(a.DB, start, count)
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprintf("Response:\n%#v", merchants))

    common.RespondWithJSON(w, http.StatusOK, merchants)
  }
}
