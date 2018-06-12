// review/routes.go

package review

import (
  "app/common"
)


// InitializeRoutes registers all the review related routes
func InitializeRoutes(a *common.App) {
  a.Router.HandleFunc("/review", createReview(a)).Methods("POST")
  a.Router.HandleFunc("/review/{id:[0-9]+}", getReview(a)).Methods("GET")
  a.Router.HandleFunc("/review/{id:[0-9]+}", updateReview(a)).Methods("PUT")
  a.Router.HandleFunc("/review/{id:[0-9]+}", deleteReview(a)).Methods("DELETE")
  a.Router.HandleFunc("/reviews", getReviews(a)).Methods("GET")
}
