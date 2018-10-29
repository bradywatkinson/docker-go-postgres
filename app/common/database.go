// app.go

package common

import (
  "database/sql"

  _ "github.com/lib/pq" // postgres sql driver

  log "github.com/sirupsen/logrus"
)


// InitializeDB setups all the main operating components off app;
// this includes creating a new database connection and setting
// up the router and registering routes.
func (a *App) InitializeDB(connectionString string) {

  var err error
  a.DB, err = sql.Open("postgres", connectionString)
  if err != nil {
    log.Fatal(err)
  }
}
