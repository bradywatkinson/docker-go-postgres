// product/initialise.go

package product

import (
  "app/common"
)


// InitializeREST registers all the product related routes
func InitializeREST(a *common.App) {
  a.Router.HandleFunc("/product", postProduct(a)).Methods("POST")
  a.Router.HandleFunc("/product/{id:[0-9]+}", getProduct(a)).Methods("GET")
  a.Router.HandleFunc("/product/{id:[0-9]+}", putProduct(a)).Methods("PUT")
  a.Router.HandleFunc("/product/{id:[0-9]+}", deleteProduct(a)).Methods("DELETE")
  a.Router.HandleFunc("/products", getProducts(a)).Methods("GET")
}


// InitializeGRPC registers the ProductService
func InitializeGRPC(a *common.App) {
  RegisterProductServiceServer(a.GRPC, &ProductServiceInterface{
    app: a,
  })
}
