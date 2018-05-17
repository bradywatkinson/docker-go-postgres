// merchant/merchant.go

package merchant

import (
  "errors"
  "net/http"

  "github.com/jinzhu/copier"
  "github.com/go-ozzo/ozzo-validation"
  "github.com/go-ozzo/ozzo-validation/is"
  "github.com/mholt/binding"
)

// Merchant holds all information about a merchant
type Merchant struct {
  Schema *MerchantSchema
  Model  *MerchantModel
}

// ID is read by copier and used to populate the MerchantSchema's
// Id field. ID is enforced by golint but protoc generates the field
// name as Id, hence the need for conversion
func (schema *MerchantSchema) ID(id int) {
  schema.Id = int32(id)
}

// Id (see above)
func (model *MerchantModel) Id(id int32) {
  if id != 0 {
    model.ID = int(id)
  }
}

func (m *Merchant) copySchema() {
  if m.Schema == nil {
    panic(errors.New("Failed to copy schema: Empty Schema"))
  }

  if m.Model == nil {
    m.Model = &MerchantModel{}
  }

  copier.Copy(m.Model, m.Schema)
}

func (m *Merchant) copyModel() {
  if m.Schema == nil {
    m.Schema = &MerchantSchema{}
  }
  copyModel(m.Model, m.Schema)
}

func copyModel(model *MerchantModel, schema *MerchantSchema) {
 if model == nil {
    panic(errors.New("Failed to copy model: Empty Model"))
  }

  copier.Copy(schema, model)
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (m *MerchantSchema) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &m.Id:   "id",
    &m.Name: "name",
  }
}

// FieldMap is used by `github.com/mholt/binding` for data binding
func (m *MerchantsQuery) FieldMap(req *http.Request) binding.FieldMap {
  return binding.FieldMap{
    &m.Start: "start",
    &m.Count: "count",
  }
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (m *MerchantQuery) Validate() error {
  return validation.ValidateStruct(m,
    validation.Field(&m.Id, validation.Required),
    validation.Field(&m.Merchant, validation.NilOrNotEmpty),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (m *MerchantSchema) Validate() error {
  return validation.ValidateStruct(m,
    validation.Field(&m.Id, validation.In(nil).Error("Cannot update id")),
    validation.Field(&m.Name, is.PrintableASCII),
  )
}

// Validate called by:
// - `github.com/mholt/binding` after data binding
// - `github.com/grpc-ecosystem/go-grpc-middleware/validator` after the request is received
func (m *MerchantsQuery) Validate() error {
  if m.Count > 10 || m.Count < 1 {
    m.Count = 10
  }
  if m.Start < 0 {
    m.Start = 0
  }
  return nil
}
