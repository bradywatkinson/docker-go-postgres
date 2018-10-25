// review/review.go

package review

import (
  "errors"
  "net/http"

  "github.com/jinzhu/copier"
  "github.com/go-ozzo/ozzo-validation"
  "github.com/go-ozzo/ozzo-validation/is"
  "github.com/mholt/binding"
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


func (r *Review) copySchema() {
  if r.Schema == nil {
    panic(errors.New("Failed to copy schema: Empty Schema"))
  }

  if r.Model == nil {
    r.Model = &ReviewModel{}
  }

  copier.Copy(r.Model, r.Schema)
}

func (r *Review) copyModel() {
  if r.Schema == nil {
    r.Schema = &ReviewSchema{}
  }
  copyModel(r.Model, r.Schema)
}

func copyModel(model *ReviewModel, schema *ReviewSchema) {
 if model == nil {
    panic(errors.New("Failed to copy model: Empty Model"))
  }

  copier.Copy(schema, model)
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (r *ReviewSchema) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &r.Id:         "id",
    &r.Rating:     "rating",
    &r.Review:     "review",
    &r.CustomerId: "customer_id",
    &r.ProductId:  "product_id",
  }
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (r *ReviewsQuery) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &r.Start: "start",
    &r.Count: "count",
  }
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (r *ReviewQuery) Validate() error {
  return validation.ValidateStruct(r,
    validation.Field(&r.Id, validation.Required),
    validation.Field(&r.Review, validation.NilOrNotEmpty),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (r *ReviewSchema) Validate() error {
  return validation.ValidateStruct(r,
    validation.Field(&r.Id, validation.In(nil).Error("Cannot update id")),
    validation.Field(&r.Rating, validation.Min(0), validation.Max(5)),
    validation.Field(&r.Review, is.PrintableASCII),
    validation.Field(&r.CustomerId, validation.Required),
    validation.Field(&r.ProductId, validation.Required),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (r *ReviewsQuery) Validate() error {
  if r.Count > 10 || r.Count < 1 {
    r.Count = 10
  }
  if r.Start < 0 {
    r.Start = 0
  }
  return nil
}
