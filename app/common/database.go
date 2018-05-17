// app.go

package common

import (
  "strings"

  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq" // postgres sql driver

  log "github.com/sirupsen/logrus"
  "github.com/o1egl/gormrus"
)


// InitializeDB setups all the main operating components off app;
// this includes creating a new database connection and setting
// up the router and registering routes.
func (a *App) InitializeDB(connectionString string) {

  var err error
  a.DB, err = gorm.Open("postgres", connectionString)
  if err != nil {
    log.Fatal(err)
  }
  a.DB.LogMode(true)
  a.DB.SetLogger(gormrus.NewWithLogger(logger))
  a.DB.SingularTable(true)
  gorm.DefaultTableNameHandler = func (db *gorm.DB, tableName string) string  {
    return strings.TrimSuffix(tableName, "_model")
  }
}
