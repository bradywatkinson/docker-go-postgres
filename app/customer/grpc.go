package customer

import (
  "context"

  "github.com/jinzhu/copier"
)

// CustomerServiceInterface is implemented by CustomerService
type CustomerServiceInterface struct{}

// CreateCustomer implements CustomerService.CreateCustomer
func (s *CustomerServiceInterface) CreateCustomer(ctx context.Context, in *CustomerSchema) (*CustomerSchema, error) {
  out := &CustomerSchema{}
  copier.Copy(in, out)
  return out, nil
}

// GetCustomer implements CustomerService.GetCustomer
func (s *CustomerServiceInterface) GetCustomer(ctx context.Context, in *CustomerSchema) (*CustomerSchema, error) {
  out := &CustomerSchema{}
  copier.Copy(in, out)
  return out, nil
}

// UpdateCustomer implements CustomerService.UpdateCustomer
func (s *CustomerServiceInterface) UpdateCustomer(ctx context.Context, in *CustomerSchema) (*CustomerSchema, error) {
  out := &CustomerSchema{}
  copier.Copy(in, out)
  return out, nil
}
