package customer

import (
  "context"

  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  wrappers "github.com/golang/protobuf/ptypes/wrappers"
  "github.com/jinzhu/gorm"

  "app/common"
)

// CustomerServiceInterface is implemented by CustomerService
type CustomerServiceInterface struct{
  app *common.App
}

// CreateCustomer implements CustomerService.CreateCustomer
func (s *CustomerServiceInterface) CreateCustomer(ctx context.Context, req *CustomerSchema) (*CustomerSchema, error) {
  c := Customer{
    Schema: req,
    Model: nil,
  }

  c.copySchema()

  if err := c.Model.createCustomer(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  c.copyModel()

  return c.Schema, nil
}

// GetCustomer implements CustomerService.CustomerQuery
func (s *CustomerServiceInterface) GetCustomer(ctx context.Context, req *CustomerQuery) (*CustomerSchema, error) {
  c := Customer{
    Model: &CustomerModel{ID: int(req.Id)},
    Schema: nil,
  }

  if err := c.Model.readCustomer(s.app.DB); err != nil {
    switch err {
    case gorm.ErrRecordNotFound:
      return nil, status.Error(codes.NotFound, "Customer not found")
    default:
      return nil, status.Error(codes.Internal, err.Error())
    }
  }

  c.copyModel()

  return c.Schema, nil
}

// UpdateCustomer implements CustomerService.UpdateCustomer
func (s *CustomerServiceInterface) UpdateCustomer(ctx context.Context, req *CustomerQuery) (*CustomerSchema, error) {
  c := Customer{
    Schema: req.Customer,
    Model: &CustomerModel{ID: int(req.Id)},
  }

  if err := c.Model.readCustomer(s.app.DB); err != nil {
    switch err {
    case gorm.ErrRecordNotFound:
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

  return c.Schema, nil
}

// DeleteCustomer implements CustomerService.DeleteCustomer
func (s *CustomerServiceInterface) DeleteCustomer(ctx context.Context, req *CustomerQuery) (*wrappers.StringValue, error) {

  c := Customer{
    Model: &CustomerModel{ID: int(req.Id)},
    Schema: nil,
  }
  if err := c.Model.deleteCustomer(s.app.DB); err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }


  m := &wrappers.StringValue{Value: "success"}
  return m, nil
}

// GetCustomers implements CustomerService.GetCustomers
func (s *CustomerServiceInterface) GetCustomers(ctx context.Context, req *CustomersQuery) (*CustomersResponse, error) {

  count, start := int(req.Count), int(req.Start)

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

  return res, nil
}
