// merchant/model.go

package merchant

import (
  "database/sql"
)


type merchant struct {
  ID    int     `json:"id"`
  Name  string  `json:"name"`
}

func (m *merchant) createMerchant(db *sql.DB) error {
  err := db.QueryRow(
    "INSERT INTO merchant(name) VALUES($1) RETURNING id",
    m.Name).Scan(&m.ID)

  if err != nil {
    return err
  }

  return nil
}

func (m *merchant) getMerchant(db *sql.DB) error {
  return db.QueryRow("SELECT name FROM merchant WHERE id=$1",
    m.ID).Scan(&m.Name)
}

func (m *merchant) updateMerchant(db *sql.DB) error {
  _, err :=
    db.Exec("UPDATE merchant SET name=$1 WHERE id=$2",
      m.Name, m.ID)

  return err
}

func (m *merchant) deleteMerchant(db *sql.DB) error {
  _, err := db.Exec("DELETE FROM merchant WHERE id=$1", m.ID)

  return err
}

func (m *merchant) getMerchants(db *sql.DB, start, count int) ([]merchant, error) {
  rows, err := db.Query(
    "SELECT id, name FROM merchant LIMIT $1 OFFSET $2",
    count, start)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  merchants := []merchant{}

  for rows.Next() {
    var m merchant
    if err := rows.Scan(&m.ID, &m.Name); err != nil {
      return nil, err
    }
    merchants = append(merchants, m)
  }

  return merchants, nil
}

