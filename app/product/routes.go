// product/routes.go

package product

import (
  "app/common"
)


// InitializeRoutes registers all the product related routes
func InitializeRoutes(a *common.App) {
  a.Router.HandleFunc("/product", createProduct(a)).Methods("POST")
  a.Router.HandleFunc("/product/{id:[0-9]+}", getProduct(a)).Methods("GET")
  a.Router.HandleFunc("/product/{id:[0-9]+}", updateProduct(a)).Methods("PUT")
  a.Router.HandleFunc("/product/{id:[0-9]+}", deleteProduct(a)).Methods("DELETE")
  a.Router.HandleFunc("/products", getProducts(a)).Methods("GET")
}
