// review/rest.go

package review

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

func postReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {

    r := Review{
      Schema: &ReviewSchema{},
    }
    if err := binding.Bind(req, r.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }
    defer req.Body.Close()

    log.Debug(fmt.Sprintf("%#v", r.Schema))

    r.copySchema()

    if err := r.Model.createReview(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    r.copyModel()

    log.Debug(fmt.Sprintf("Response:\n%#v", r.Schema))
    common.RespondWithJSON(w, http.StatusCreated, r.Schema)
  }
}

func getReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Review ID")
      return
    }

    log.Debug(fmt.Sprintf("{ id: %d }", id))

    r := Review{
      Model: &ReviewModel{ID: id},
      Schema: &ReviewSchema{},
    }
    if err := r.Model.readReview(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Review not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    r.copyModel()

    log.Debug(fmt.Sprintf("Response:\n%#v", r.Schema))

    common.RespondWithJSON(w, http.StatusOK, r.Schema)
  }
}

func putReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Review ID")
      return
    }

    r := Review{
      Model: &ReviewModel{ID: id},
      Schema: &ReviewSchema{},
    }

    if err := r.Model.readReview(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Review not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    log.Debug(fmt.Sprintf("Review: %#v", r.Model))

    if err := binding.Bind(req, r.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }
    defer req.Body.Close()

    r.copySchema()

    if err := r.Model.updateReview(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    r.copyModel()

    log.Debug(fmt.Sprintf("Response:\n%#v", r.Schema))

    common.RespondWithJSON(w, http.StatusOK, r.Schema)
  }
}

func deleteReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Review ID")
      return
    }

    log.Debug(fmt.Sprintf("{ id: %d }", id))

    r := Review{
      Model: &ReviewModel{ID: id},
      Schema: nil,
    }
    if err := r.Model.deleteReview(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }


    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getReviews(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    q := &ReviewsQuery{}
    if err := binding.Bind(req, q); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err.Error()))
      return
    }

    log.Debug(fmt.Sprintf("{ count: %d, start: %d }", q.Count, q.Start))

    reviews, err := readReviews(a.DB, int(q.Start), int(q.Count))
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    log.Debug(fmt.Sprintf("Response:\n%#v", reviews))

    common.RespondWithJSON(w, http.StatusOK, reviews)
  }
}
