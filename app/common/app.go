// common/app.go

package common

import (
  "github.com/jinzhu/gorm"

  "github.com/gorilla/mux"
  "google.golang.org/grpc"
)


// App serves as the structure to hold the state of the app
type App struct {
  Router *mux.Router
  DB     *gorm.DB
  GRPC   *grpc.Server
}
