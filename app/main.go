// main.go

package main

import (
  "fmt"
  "os"
)

func main() {
  a := App{}
  connectionString :=
    fmt.Sprintf("sslmode=disable host=postgres user=%s password=%s dbname=%s",
      os.Getenv("APP_DB_USERNAME"),
      os.Getenv("APP_DB_PASSWORD"),
      os.Getenv("APP_DB_NAME"))
  a.Initialize(connectionString)

  a.Run(":8080")
}
