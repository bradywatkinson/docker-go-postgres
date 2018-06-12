// Customer/routes.go

package customer

import (
  "app/common"
)


// InitializeRoutes registers all the customer related routes
func InitializeRoutes(a *common.App) {
  a.Router.HandleFunc("/customer", createCustomer(a)).Methods("POST")
  a.Router.HandleFunc("/customer/{id:[0-9]+}", getCustomer(a)).Methods("GET")
  a.Router.HandleFunc("/customer/{id:[0-9]+}", updateCustomer(a)).Methods("PUT")
  a.Router.HandleFunc("/customer/{id:[0-9]+}", deleteCustomer(a)).Methods("DELETE")
  a.Router.HandleFunc("/customers", getCustomers(a)).Methods("GET")
}
