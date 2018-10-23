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

  c.copySchema()

  if err := c.Model.updateCustomer(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))
  return c.Schema, nil
}

// DeleteCustomer implements CustomerService.DeleteCustomer
func (s *CustomerServiceInterface) DeleteCustomer(ctx context.Context, req *CustomerRequest) (*wrappers.StringValue, error) {
  testutils.Log(fmt.Sprint("CustomerService.DeleteCustomer"))

  c := Customer{
    Model: &CustomerModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := c.Model.deleteCustomer(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  testutils.Log(fmt.Sprint("Response:\n{ value: \"success\" }"))

  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

// GetCustomers implements CustomerService.GetCustomers
func (s *CustomerServiceInterface) GetCustomers(ctx context.Context, req *CustomersRequest) (*CustomersResponse, error) {
  testutils.Log(fmt.Sprint("CustomerService.GetCustomers"))

  count, start := int(req.Count), int(req.Start)

  if count > 10 || count < 1 {
    count = 10
  }
  if start < 0 {
    start = 0
  }

  testutils.Log(fmt.Sprintf("{ count: %d, start: %d }", count, start))

  customers, err := readCustomers(s.app.DB, start, count)
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  res := &CustomersResponse{
    Customers: []*CustomerSchema{},
  }

  for _, c := range customers {
    tmp := &CustomerSchema{}
    copyModel(&c, tmp)
    res.Customers = append(res.Customers, tmp)
  }

  testutils.Log(fmt.Sprintf("Response:\n%#v", customers))

  return res, nil
}
