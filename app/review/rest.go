// review/rest.go

package review

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

func postReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("POST /review:"))

    c := Review{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c.Schema); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
      return
    }
    defer r.Body.Close()

    testutils.Log(fmt.Sprintf("%#v", c.Schema))

    c.copySchema()

    if err := c.Model.createReview(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
    common.RespondWithJSON(w, http.StatusCreated, c.Schema)
  }
}

func getReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /review:"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Review ID")
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Review{
      Model: &ReviewModel{ID: id},
      Schema: &ReviewSchema{},
    }
    if err := c.Model.readReview(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Review not found")
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

func putReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("PUT /review"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Review ID")
      return
    }

    c := Review{
      Model: &ReviewModel{ID: id},
      Schema: nil,
    }

    if err := c.Model.readReview(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Review not found")
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

    if err := c.Model.updateReview(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    c.copyModel()

    testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

    common.RespondWithJSON(w, http.StatusOK, c.Schema)
  }
}

func deleteReview(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("DELETE /review"))
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Review ID")
      return
    }

    testutils.Log(fmt.Sprintf("{ id: %d }", id))

    c := Review{
      Model: &ReviewModel{ID: id},
      Schema: nil,
    }
    if err := c.Model.deleteReview(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprint("Response:\n{ result: \"success\" }"))

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getReviews(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    testutils.Log(fmt.Sprint("GET /reviews"))
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
      count = 10
    }
    if start < 0 {
      start = 0
    }

    testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

    reviews, err := readReviews(a.DB, start, count)
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    testutils.Log(fmt.Sprintf("Response:\n%#v", reviews))

    common.RespondWithJSON(w, http.StatusOK, reviews)
  }
}
