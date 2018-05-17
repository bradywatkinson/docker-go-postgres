// customer/models.go

package customer

import (
  "time"

  "github.com/jinzhu/gorm"
)

// CustomerModel is used to load/dump from the
// database customer table
type CustomerModel struct {
  ID        int       `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time `sql:"index"`
  FirstName string
  LastName  string
}

func (c *CustomerModel) createCustomer(db *gorm.DB) error {
  return db.Create(&c).Error
}

func (c *CustomerModel) readCustomer(db *gorm.DB) error {
  return db.First(&c, c.ID).Error
}

func (c *CustomerModel) updateCustomer(db *gorm.DB) error {
  return db.Model(&c).Updates(&c).Error
}

func (c *CustomerModel) deleteCustomer(db *gorm.DB) error {
  return db.Delete(&c).Error
}

func readCustomers(db *gorm.DB, start, count int) ([]CustomerModel, error) {
  customers := []CustomerModel{}
  if err := db.Limit(count).Offset(start).Find(&customers).Error; err != nil {
    return nil, err
  }

  return customers, nil
}
