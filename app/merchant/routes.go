// merchant/routes.go

package merchant

import (
  "app/common"
)


// InitializeRoutes registers all the merchant related routes
func InitializeRoutes(a *common.App) {
  a.Router.HandleFunc("/merchant/{id:[0-9]+}", getMerchant(a)).Methods("GET")
  a.Router.HandleFunc("/merchant", createMerchant(a)).Methods("POST")
  a.Router.HandleFunc("/merchant/{id:[0-9]+}", updateMerchant(a)).Methods("PUT")
  a.Router.HandleFunc("/merchant/{id:[0-9]+}", deleteMerchant(a)).Methods("DELETE")
  a.Router.HandleFunc("/merchants", getMerchants(a)).Methods("GET")
}
