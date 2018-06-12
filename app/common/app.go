// common/app.go

package common

import (
  "database/sql"

  "github.com/gorilla/mux"
)


// App serves as the structure to hold the state of the app
type App struct {
  Router *mux.Router
  DB     *sql.DB
}
