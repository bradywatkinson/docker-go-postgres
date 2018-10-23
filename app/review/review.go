// review/review.go

package review

import (
  "errors"

  "github.com/jinzhu/copier"
)

// Review holds all information about a review
type Review struct {
  Schema *ReviewSchema
  Model  *ReviewModel
}

// ID is read by copier and used to populate the ReviewSchema's
// Id field. ID is enforced by golint but protoc generates the field
// name as Id, hence the need for conversion
func (schema *ReviewSchema) ID(id int) {
  schema.Id = int32(id)
}

// ProductID (see above)
func (schema *ReviewSchema) ProductID(product_id int) {
  schema.ProductId = int32(product_id)
}

// CustomerID (see above)
func (schema *ReviewSchema) CustomerID(customer_id int) {
  schema.CustomerId = int32(customer_id)
}

// Id (see above)
func (model *ReviewModel) Id(id int32) {
  if id != 0 {
    model.ID = int(id)
  }
}

// ProductId (see above)
func (model *ReviewModel) ProductId(product_id int32) {
  if product_id != 0 {
    model.ProductID = int(product_id)
  }
}

// CustomerId (see above)
func (model *ReviewModel) CustomerId(customer_id int32) {
  if customer_id != 0 {
    model.CustomerID = int(customer_id)
  }
}


func (c *Review) copySchema() {
  if c.Schema == nil {
    panic(errors.New("Failed to copy schema: Empty Schema"))
  }

  if c.Model == nil {
    c.Model = &ReviewModel{}
  }

  copier.Copy(c.Model, c.Schema)
}

func (c *Review) copyModel() {
  if c.Schema == nil {
    c.Schema = &ReviewSchema{}
  }
  copyModel(c.Model, c.Schema)
}

func copyModel(model *ReviewModel, schema *ReviewSchema) {
 if model == nil {
    panic(errors.New("Failed to copy model: Empty Model"))
  }

  copier.Copy(schema, model)
}
