// merchant/initialise.go

package merchant

import (
  "app/common"
)


// InitializeREST registers all the merchant related routes
func InitializeREST(a *common.App) {
  a.Router.HandleFunc("/merchant", postMerchant(a)).Methods("POST")
  a.Router.HandleFunc("/merchant/{id:[0-9]+}", getMerchant(a)).Methods("GET")
  a.Router.HandleFunc("/merchant/{id:[0-9]+}", putMerchant(a)).Methods("PUT")
  a.Router.HandleFunc("/merchant/{id:[0-9]+}", deleteMerchant(a)).Methods("DELETE")
  a.Router.HandleFunc("/merchants", getMerchants(a)).Methods("GET")
}


// InitializeGRPC registers the MerchantService
func InitializeGRPC(a *common.App) {
  RegisterMerchantServiceServer(a.GRPC, &MerchantServiceInterface{
    app: a,
  })
}
