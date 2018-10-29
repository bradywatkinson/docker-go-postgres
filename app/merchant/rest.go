// merchant/rest.go

package merchant

import (
  "database/sql"
  "net/http"
  "strconv"
  "fmt"

  "github.com/gorilla/mux"
  "github.com/mholt/binding"
  log "github.com/sirupsen/logrus"

  "app/common"
)

func postMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {

    m := Merchant{
      Schema: &MerchantSchema{},
    }

    if err := binding.Bind(req, m.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    log.Debug(fmt.Sprintf("Merchant: %#v", m.Schema))

    m.copySchema()

    if err := m.Model.createMerchant(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    m.copyModel()

    log.Debug(fmt.Sprintf("Response:\n%#v", m.Schema))
    common.RespondWithJSON(w, http.StatusCreated, m.Schema)
  }
}

func getMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Merchant ID")
      return
    }

    log.Debug(fmt.Sprintf("{ id: %d }", id))

    m := Merchant{
      Model: &MerchantModel{ID: id},
      Schema: &MerchantSchema{},
    }
    if err := m.Model.readMerchant(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Merchant not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    m.copyModel()

    log.Debug(fmt.Sprintf("Response:\n%#v", m.Schema))

    common.RespondWithJSON(w, http.StatusOK, m.Schema)
  }
}

func putMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Merchant ID")
      return
    }

    log.Debug(fmt.Sprintf("{ id: %d }", id))

    m := Merchant{
      Model: &MerchantModel{ID: id},
      Schema: &MerchantSchema{},
    }

    if err := m.Model.readMerchant(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Merchant not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    log.Debug(fmt.Sprintf("Merchant: %#v", m.Model))

    if err := binding.Bind(req, m.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }
    defer req.Body.Close()

    m.copySchema()

    if err := m.Model.updateMerchant(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    m.copyModel()

    log.Debug(fmt.Sprintf("Response:\n%#v", m.Schema))

    common.RespondWithJSON(w, http.StatusOK, m.Schema)
  }
}

func deleteMerchant(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Merchant ID")
      return
    }

    log.Debug(fmt.Sprintf("{ id: %d }", id))

    m := Merchant{
      Model: &MerchantModel{ID: id},
      Schema: nil,
    }
    if err := m.Model.deleteMerchant(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }


    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getMerchants(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    q := &MerchantsQuery{}
    if err := binding.Bind(req, q); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    log.Debug(fmt.Sprintf("{ count: %d, start: %d }", q.Count, q.Start))

    merchants, err := readMerchants(a.DB, int(q.Start), int(q.Count))
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    log.Debug(fmt.Sprintf("Response:\n%#v", merchants))

    common.RespondWithJSON(w, http.StatusOK, merchants)
  }
}
