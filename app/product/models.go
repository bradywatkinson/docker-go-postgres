// product/models.go

package product

import (
  "time"

  "github.com/jinzhu/gorm"
)

// ProductModel is used to load/dump from the
// database product table
type ProductModel struct {
  ID         int       `gorm:"primary_key"`
  CreatedAt  time.Time
  UpdatedAt  time.Time
  DeletedAt  *time.Time `sql:"index"`
  Name       string
  Price      float64
  MerchantID int
}

func (p *ProductModel) createProduct(db *gorm.DB) error {
  return db.Create(&p).Error
}

func (p *ProductModel) readProduct(db *gorm.DB) error {
  return db.First(&p).Error
}

func (p *ProductModel) updateProduct(db *gorm.DB) error {
  return db.Model(&p).Updates(&p).Error
}

func (p *ProductModel) deleteProduct(db *gorm.DB) error {
  return db.Delete(&p).Error
}

func readProducts(db *gorm.DB, start, count int) ([]ProductModel, error) {
  products := []ProductModel{}
  if err := db.Limit(count).Offset(start).Find(&products).Error; err != nil {
    return nil, err
  }

  return products, nil
}
