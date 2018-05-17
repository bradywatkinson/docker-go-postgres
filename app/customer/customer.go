// customer/customer.go

package customer

import (
  "errors"
  "net/http"

  "github.com/jinzhu/copier"
  "github.com/go-ozzo/ozzo-validation"
  "github.com/go-ozzo/ozzo-validation/is"
  "github.com/mholt/binding"
)

// Customer holds all information about a customer
type Customer struct {
  Schema *CustomerSchema
  Model  *CustomerModel
}

// ID is read by copier and used to populate the CustomerSchema's
// Id field. ID is enforced by golint but protoc generates the field
// name as Id, hence the need for conversion
func (schema *CustomerSchema) ID(id int) {
  schema.Id = int32(id)
}

// Id (see above)
func (model *CustomerModel) Id(id int32) {
  if id != 0 {
    model.ID = int(id)
  }
}

func (c *Customer) copySchema() {
  if c.Schema == nil {
    panic(errors.New("Failed to copy schema: Empty Schema"))
  }

  if c.Model == nil {
    c.Model = &CustomerModel{}
  }

  copier.Copy(c.Model, c.Schema)
}

func (c *Customer) copyModel() {
  if c.Schema == nil {
    c.Schema = &CustomerSchema{}
  }
  copyModel(c.Model, c.Schema)
}

func copyModel(model *CustomerModel, schema *CustomerSchema) {
 if model == nil {
    panic(errors.New("Failed to copy model: Empty Model"))
  }

  copier.Copy(schema, model)
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (c *CustomerSchema) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &c.Id:        "id",
    &c.FirstName: "first_name",
    &c.LastName:  "last_name",
  }
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (c *CustomersQuery) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &c.Start: "start",
    &c.Count: "count",
  }
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (c *CustomerQuery) Validate() error {
  return validation.ValidateStruct(c,
    validation.Field(&c.Id, validation.Required),
    validation.Field(&c.Customer, validation.NilOrNotEmpty),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (c *CustomerSchema) Validate() error {
  return validation.ValidateStruct(c,
    validation.Field(&c.Id, validation.In(nil).Error("Cannot update id")),
    validation.Field(&c.FirstName, is.Alpha),
    validation.Field(&c.LastName, is.Alpha),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (c *CustomersQuery) Validate() error {
  if c.Count > 10 || c.Count < 1 {
    c.Count = 10
  }
  if c.Start < 0 {
    c.Start = 0
  }
  return nil
}
