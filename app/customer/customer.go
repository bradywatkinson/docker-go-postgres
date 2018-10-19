// customer/customer.go

package customer

import (
  "errors"

  "github.com/jinzhu/copier"
)

type Customer struct {
  Schema *CustomerSchema
  Model  *CustomerModel
}

func (schema *CustomerSchema) ID(id int) {
  schema.Id = int32(id)
}

func (model *CustomerModel) Id(id int32) {
  model.ID = int(id)
}

func (c *Customer) copySchema() error {
  if c.Schema == nil {
    return errors.New("Failed to copy schema: Empty Schema")
  }

  if c.Model == nil {
    c.Model = &CustomerModel{}
  }

  copier.Copy(c.Schema, c.Model)
  return nil
}

func (c *Customer) copyModel() error {
  if c.Schema == nil {
    return errors.New("Failed to copy model: Empty Model")
  }

  if c.Schema == nil {
    c.Schema = &CustomerSchema{}
  }

  copier.Copy(c.Model, c.Schema)
  return nil
}
