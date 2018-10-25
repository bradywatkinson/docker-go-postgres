// review/models.go

package review

import (
  "database/sql"
)

// ReviewModel is used to load/dump from the
// database review table
type ReviewModel struct {
  ID         int
  Rating     int
  Review     string
  CustomerID int
  ProductID  int
}

func (r *ReviewModel) createReview(db *sql.DB) error {
  err := db.QueryRow(
    "INSERT INTO review(rating, review, customer_id, product_id) VALUES($1, $2, $3, $4) RETURNING id",
    r.Rating, r.Review, r.CustomerID, r.ProductID).Scan(&r.ID)

  if err != nil {
    return err
  }

  return nil
}

func (r *ReviewModel) readReview(db *sql.DB) error {
  return db.QueryRow("SELECT rating, review FROM review WHERE id=$1",
    r.ID).Scan(&r.Rating, &r.Review)
}

func (r *ReviewModel) updateReview(db *sql.DB) error {
  _, err :=
    db.Exec("UPDATE review SET rating=$1, review=$2 WHERE id=$3",
      r.Rating, r.Review, r.ID)

  return err
}

func (r *ReviewModel) deleteReview(db *sql.DB) error {
  _, err := db.Exec("DELETE FROM review WHERE id=$1", r.ID)

  return err
}

func readReviews(db *sql.DB, start, count int) ([]ReviewModel, error) {
  rows, err := db.Query(
    "SELECT id, rating, review, customer_id, product_id FROM review LIMIT $1 OFFSET $2",
    count, start)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  reviews := []ReviewModel{}

  for rows.Next() {
    var r ReviewModel
    if err := rows.Scan(&r.ID, &r.Rating, &r.Review, &r.CustomerID, &r.ProductID); err != nil {
      return nil, err
    }
    reviews = append(reviews, r)
  }

  return reviews, nil
}
