package customer

import (
  "context"
  "fmt"

  "github.com/jinzhu/copier"

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

  _ = c.Model.createCustomer(s.app.DB);

  c.copyModel()

  testutils.Log(fmt.Sprintf("Response:\n%#v", c.Schema))

  return c.Schema, nil
}

// GetCustomer implements CustomerService.GetCustomer
func (s *CustomerServiceInterface) GetCustomer(ctx context.Context, req *CustomerSchema) (*CustomerSchema, error) {
  res := &CustomerSchema{}
  copier.Copy(res, req)
  return res, nil
}

// UpdateCustomer implements CustomerService.UpdateCustomer
func (s *CustomerServiceInterface) UpdateCustomer(ctx context.Context, req *CustomerSchema) (*CustomerSchema, error) {
  res := &CustomerSchema{}
  copier.Copy(res, req)
  return res, nil
}
