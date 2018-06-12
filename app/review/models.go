// review/models.go

package review

import (
  "database/sql"
)


type review struct {
  ID         int    `json:"id"`
  Rating     int    `json:"rating"`
  Review     string `json:"review"`
  CustomerID int    `json:"customer_id"`
  ProductID  int    `json:"product_id"`
}

func (r *review) createReview(db *sql.DB) error {
  err := db.QueryRow(
    "INSERT INTO review(rating, review, customer_id, product_id) VALUES($1, $2, $3, $4) RETURNING id",
    r.Rating, r.Review, r.CustomerID, r.ProductID).Scan(&r.ID)

  if err != nil {
    return err
  }

  return nil
}

func (r *review) getReview(db *sql.DB) error {
  return db.QueryRow("SELECT rating, review FROM review WHERE id=$1",
    r.ID).Scan(&r.Rating, &r.Review)
}

func (r *review) updateReview(db *sql.DB) error {
  _, err :=
    db.Exec("UPDATE review SET rating=$1, review=$2 WHERE id=$3",
      r.Rating, r.Review, r.ID)

  return err
}

func (r *review) deleteReview(db *sql.DB) error {
  _, err := db.Exec("DELETE FROM review WHERE id=$1", r.ID)

  return err
}

func (r *review) getReviews(db *sql.DB, start, count int) ([]review, error) {
  rows, err := db.Query(
    "SELECT id, rating, review FROM review LIMIT $1 OFFSET $2",
    count, start)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  reviews := []review{}

  for rows.Next() {
    var r review
    if err := rows.Scan(&r.ID, &r.Rating, &r.Review, &r.CustomerID, &r.ProductID); err != nil {
      return nil, err
    }
    reviews = append(reviews, r)
  }

  return reviews, nil
}
