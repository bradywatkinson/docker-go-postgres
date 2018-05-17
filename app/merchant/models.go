// merchant/model.go

package merchant

import (
  "time"

  "github.com/jinzhu/gorm"
)

// MerchantModel is used to load/dump from the
// database merchant table
type MerchantModel struct {
  ID        int       `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time `sql:"index"`
  Name      string
}

func (m *MerchantModel) createMerchant(db *gorm.DB) error {
  return db.Create(&m).Error
}

func (m *MerchantModel) readMerchant(db *gorm.DB) error {
  return db.First(&m, m.ID).Error
}

func (m *MerchantModel) updateMerchant(db *gorm.DB) error {
  return db.Model(&m).Updates(&m).Error
}

func (m *MerchantModel) deleteMerchant(db *gorm.DB) error {
  return db.Delete(&m).Error
}

func readMerchants(db *gorm.DB, start, count int) ([]MerchantModel, error) {
  merchants := []MerchantModel{}
  if err := db.Limit(count).Offset(start).Find(&merchants).Error; err != nil {
    return nil, err
  }

  return merchants, nil
}
