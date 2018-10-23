// customer/customer.go

package customer

import (
  "errors"

  "github.com/jinzhu/copier"
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
