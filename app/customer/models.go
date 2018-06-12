// customer/models.go

package customer

import (
  "database/sql"
)


type customer struct {
  ID        int    `json:"id"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
}

func (c *customer) createCustomer(db *sql.DB) error {
  err := db.QueryRow(
    "INSERT INTO customer(first_name, last_name) VALUES($1, $2) RETURNING id",
    c.FirstName, c.LastName).Scan(&c.ID)

  if err != nil {
    return err
  }

  return nil
}

func (c *customer) getCustomer(db *sql.DB) error {
  return db.QueryRow("SELECT first_name, last_name FROM customer WHERE id=$1",
    c.ID).Scan(&c.FirstName, &c.LastName)
}

func (c *customer) updateCustomer(db *sql.DB) error {
  _, err :=
    db.Exec("UPDATE customer SET first_name=$1, last_name=$2 WHERE id=$3",
      c.FirstName, c.LastName, c.ID)

  return err
}

func (c *customer) deleteCustomer(db *sql.DB) error {
  _, err := db.Exec("DELETE FROM customer WHERE id=$1", c.ID)

  return err
}

func (c *customer) getCustomers(db *sql.DB, start, count int) ([]customer, error) {
  rows, err := db.Query(
    "SELECT id, first_name, last_name FROM customer LIMIT $1 OFFSET $2",
    count, start)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  customers := []customer{}

  for rows.Next() {
    var c customer
    if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName); err != nil {
      return nil, err
    }
    customers = append(customers, c)
  }

  return customers, nil
}
