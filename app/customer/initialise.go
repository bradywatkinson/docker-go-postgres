// customer/initialise.go

package customer

import (
  "app/common"
)


// InitializeREST registers all the customer related routes
func InitializeREST(a *common.App) {
  a.Router.HandleFunc("/customer", postCustomer(a)).Methods("POST")
  a.Router.HandleFunc("/customer/{id:[0-9]+}", getCustomer(a)).Methods("GET")
  a.Router.HandleFunc("/customer/{id:[0-9]+}", putCustomer(a)).Methods("PUT")
  a.Router.HandleFunc("/customer/{id:[0-9]+}", deleteCustomer(a)).Methods("DELETE")
  a.Router.HandleFunc("/customers", getCustomers(a)).Methods("GET")
}


// InitializeGRPC registers the CustomerService
func InitializeGRPC(a *common.App) {
  RegisterCustomerServiceServer(a.GRPC, &CustomerServiceInterface{
    app: a,
  })
}
