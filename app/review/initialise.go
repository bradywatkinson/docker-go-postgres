// review/initialise.go

package review

import (
  "app/common"
)


// InitializeREST registers all the review related routes
func InitializeREST(a *common.App) {
  a.Router.HandleFunc("/review", postReview(a)).Methods("POST")
  a.Router.HandleFunc("/review/{id:[0-9]+}", getReview(a)).Methods("GET")
  a.Router.HandleFunc("/review/{id:[0-9]+}", putReview(a)).Methods("PUT")
  a.Router.HandleFunc("/review/{id:[0-9]+}", deleteReview(a)).Methods("DELETE")
  a.Router.HandleFunc("/reviews", getReviews(a)).Methods("GET")
}


// InitializeGRPC registers the ReviewService
func InitializeGRPC(a *common.App) {
  RegisterReviewServiceServer(a.GRPC, &ReviewServiceInterface{
    app: a,
  })
}
