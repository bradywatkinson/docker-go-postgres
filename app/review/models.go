// review/models.go

package review

import (
  "time"

  "github.com/jinzhu/gorm"
)

// ReviewModel is used to load/dump from the
// database review table
type ReviewModel struct {
  ID         int        `gorm:"primary_key"`
  CreatedAt  time.Time
  UpdatedAt  time.Time
  DeletedAt  *time.Time `sql:"index"`
  Rating     int
  Review     string
  CustomerID int
  ProductID  int
}

func (r *ReviewModel) createReview(db *gorm.DB) error {
  return db.Create(&r).Error
}

func (r *ReviewModel) readReview(db *gorm.DB) error {
  return db.First(&r).Error
}

func (r *ReviewModel) updateReview(db *gorm.DB) error {
  return db.Model(&r).Updates(&r).Error
}

func (r *ReviewModel) deleteReview(db *gorm.DB) error {
  return db.Delete(&r).Error
}

func readReviews(db *gorm.DB, start, count int) ([]ReviewModel, error) {
  reviews := []ReviewModel{}
  if err := db.Limit(count).Offset(start).Find(&reviews).Error; err != nil {
    return nil, err
  }

  return reviews, nil
}
