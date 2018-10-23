// product/product.go

package product

import (
  "errors"

  "github.com/jinzhu/copier"
)

// Product holds all information about a product
type Product struct {
  Schema *ProductSchema
  Model  *ProductModel
}

// ID is read by copier and used to populate the ProductSchema's
// Id field. ID is enforced by golint but protoc generates the field
// name as Id, hence the need for conversion
func (schema *ProductSchema) ID(id int) {
  schema.Id = int32(id)
}

// MerchantID same as above
func (schema *ProductSchema) MerchantID(merchant_id int) {
  schema.MerchantId = int32(merchant_id)
}

// Id (see above)
func (model *ProductModel) Id(id int32) {
  if id != 0 {
    model.ID = int(id)
  }
}

func (model *ProductModel) MerchantId(merchant_id int32) {
  if merchant_id != 0 {
    model.MerchantID = int(merchant_id)
  }
}


func (c *Product) copySchema() {
  if c.Schema == nil {
    panic(errors.New("Failed to copy schema: Empty Schema"))
  }

  if c.Model == nil {
    c.Model = &ProductModel{}
  }

  copier.Copy(c.Model, c.Schema)
}

func (c *Product) copyModel() {
  if c.Schema == nil {
    c.Schema = &ProductSchema{}
  }
  copyModel(c.Model, c.Schema)
}

func copyModel(model *ProductModel, schema *ProductSchema) {
 if model == nil {
    panic(errors.New("Failed to copy model: Empty Model"))
  }

  copier.Copy(schema, model)
}
