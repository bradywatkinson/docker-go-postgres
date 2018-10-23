package customer

import (
  "database/sql"
  "context"
  "fmt"

  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  wrappers "github.com/golang/protobuf/ptypes/wrappers"

  "app/common"
  "app/test_utils"
)

// CustomerServiceInterface is implemented by CustomerService
type CustomerServiceInterface struct{
  app *common.App
}

// CreateCustomer implements CustomerService.CreateCustomer
func (s *CustomerServiceInterface) CreateCustomer(ctx context.Context, req *CustomerSchema) (*CustomerSchema, error) {
  testutils.Log(fmt.Sprint("CustomerService.CreateCustomer"))
  c := Customer{
    Schema: req,
    Model: nil,
  }

  c.copySchema()

  if err := c.Model.createCustomer(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// GetCustomer implements CustomerService.CustomerRequest
func (s *CustomerServiceInterface) GetCustomer(ctx context.Context, req *CustomerRequest) (*CustomerSchema, error) {
  testutils.Log(fmt.Sprint("CustomerService.GetCustomer"))
  c := Customer{
    Model: &CustomerModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := c.Model.readCustomer(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Customer not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// UpdateCustomer implements CustomerService.UpdateCustomer
func (s *CustomerServiceInterface) UpdateCustomer(ctx context.Context, req *CustomerRequest) (*CustomerSchema, error) {
  testutils.Log(fmt.Sprint("CustomerService.UpdateCustomer"))
  c := Customer{
    Schema: req.Customer,
    Model: &CustomerModel{ID: int(req.Id)},
  }

  if err := c.Model.readCustomer(s.app.DB); err != nil {
    switch err {
    case sql.ErrNoRows:
      return nil, status.Error(codes.NotFound, "Customer not found")
    default:
      return nil, status. Error(codes.Internal, err.Error())
    }
  }

  testutils.Log(fmt.Sprintf("%#v", c.Schema))
  c.copySchema()

  if err := c.Model.updateCustomer(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
  return c.Schema, nil
}

func (s *CustomerServiceInterface) DeleteCustomer(ctx context.Context, req *CustomerRequest) (*wrappers.StringValue, error) {
  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

func (s *CustomerServiceInterface) GetCustomers(ctx context.Context, req *CustomersRequest) (*CustomersResponse, error) {
  customers := &CustomersResponse{}
  return customers, nil
}
