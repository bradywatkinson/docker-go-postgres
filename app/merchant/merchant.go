// merchant/merchant.go

package merchant

import (
  "errors"

  "github.com/jinzhu/copier"
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

func (c *Merchant) copySchema() {
  if c.Schema == nil {
    panic(errors.New("Failed to copy schema: Empty Schema"))
  }

  if c.Model == nil {
    c.Model = &MerchantModel{}
  }

  copier.Copy(c.Model, c.Schema)
}

func (c *Merchant) copyModel() {
  if c.Schema == nil {
    c.Schema = &MerchantSchema{}
  }
  copyModel(c.Model, c.Schema)
}

func copyModel(model *MerchantModel, schema *MerchantSchema) {
 if model == nil {
    panic(errors.New("Failed to copy model: Empty Model"))
  }

  copier.Copy(schema, model)
}
