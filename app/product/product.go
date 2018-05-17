// product/product.go

package product

import (
  "errors"
  "net/http"

  "github.com/jinzhu/copier"
  "github.com/go-ozzo/ozzo-validation"
  "github.com/go-ozzo/ozzo-validation/is"
  "github.com/mholt/binding"
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

// ProductID same as above
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


func (p *Product) copySchema() {
  if p.Schema == nil {
    panic(errors.New("Failed to copy schema: Empty Schema"))
  }

  if p.Model == nil {
    p.Model = &ProductModel{}
  }

  copier.Copy(p.Model, p.Schema)
}

func (p *Product) copyModel() {
  if p.Schema == nil {
    p.Schema = &ProductSchema{}
  }
  copyModel(p.Model, p.Schema)
}

func copyModel(model *ProductModel, schema *ProductSchema) {
 if model == nil {
    panic(errors.New("Failed to copy model: Empty Model"))
  }

  copier.Copy(schema, model)
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (p *ProductSchema) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &p.Id:         "id",
    &p.Name:       "name",
    &p.Price:      "price",
    &p.MerchantId: "merchant_id",
  }
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (p *ProductsQuery) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &p.Start: "start",
    &p.Count: "count",
  }
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (p *ProductQuery) Validate() error {
  return validation.ValidateStruct(p,
    validation.Field(&p.Id, validation.Required),
    validation.Field(&p.Product, validation.NilOrNotEmpty),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (p *ProductSchema) Validate() error {
  return validation.ValidateStruct(p,
    validation.Field(&p.Id, validation.In(nil).Error("Cannot update id")),
    validation.Field(&p.Name, is.PrintableASCII),
    validation.Field(&p.Price),
    validation.Field(&p.MerchantId, validation.Required),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (p *ProductsQuery) Validate() error {
  if p.Count > 10 || p.Count < 1 {
    p.Count = 10
  }
  if p.Start < 0 {
    p.Start = 0
  }
  return nil
}
