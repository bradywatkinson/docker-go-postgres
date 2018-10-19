package customer

import (
  "context"
)

// CustomerServiceInterface is implemented by CustomerService
type CustomerServiceInterface struct{}

// CreateCustomer implements CustomerService.CreateCustomer
func (s *CustomerServiceInterface) CreateCustomer(ctx context.Context, in *CustomerSchema) (*CustomerSchema, error) {
  return &CustomerSchema{in}, nil
}

// GetCustomer implements CustomerService.GetCustomer
func (s *CustomerServiceInterface) GetCustomer(ctx context.Context, in *CustomerSchema) (*CustomerSchema, error) {
  return &CustomerSchema{in}, nil
}

// UpdateCustomer implements CustomerService.UpdateCustomer
func (s *CustomerServiceInterface) UpdateCustomer(ctx context.Context, in *CustomerSchema) (*CustomerSchema, error) {
  return &CustomerSchema{&in}, nil
}
