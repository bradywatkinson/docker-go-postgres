// common/app.go

package common

import (
  "database/sql"

  "github.com/gorilla/mux"
  "google.golang.org/grpc"
)


// App serves as the structure to hold the state of the app
type App struct {
  Router *mux.Router
  DB     *sql.DB
  GRPC   *grpc.Server
}
