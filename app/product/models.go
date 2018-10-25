// product/models.go

package product

import (
  "database/sql"
)

// ProductModel is used to load/dump from the
// database product table
type ProductModel struct {
  ID         int
  Name       string
  Price      float64
  MerchantID int
}

func (p *ProductModel) createProduct(db *sql.DB) error {
  err := db.QueryRow(
    "INSERT INTO product(name, price, merchant_id) VALUES($1, $2, $3) RETURNING id",
    p.Name, p.Price, p.MerchantID).Scan(&p.ID)

  if err != nil {
    return err
  }

  return nil
}

func (p *ProductModel) readProduct(db *sql.DB) error {
  return db.QueryRow("SELECT name, price, merchant_id FROM product WHERE id=$1",
    p.ID).Scan(&p.Name, &p.Price, &p.MerchantID)
}

func (p *ProductModel) updateProduct(db *sql.DB) error {
  _, err :=
    db.Exec("UPDATE product SET name=$1, price=$2 WHERE id=$3",
      p.Name, p.Price, p.ID)

  return err
}

func (p *ProductModel) deleteProduct(db *sql.DB) error {
  _, err := db.Exec("DELETE FROM product WHERE id=$1", p.ID)

  return err
}

func readProducts(db *sql.DB, start, count int) ([]ProductModel, error) {
  rows, err := db.Query(
    "SELECT id, name, price, merchant_id FROM product LIMIT $1 OFFSET $2",
    count, start)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  products := []ProductModel{}

  for rows.Next() {
    var p ProductModel
    if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.MerchantID); err != nil {
      return nil, err
    }
    products = append(products, p)
  }

  return products, nil
}
