// main.go

package main

import (
  "fmt"
  "os"
  "log"
  "net/http"

  "app/common"
  "app/customer"
  "app/merchant"
  "app/product"
  "app/review"
)

func main() {
  a := common.App{}
  connectionString :=
    fmt.Sprintf("sslmode=disable host=postgres user=%s password=%s dbname=%s",
      os.Getenv("APP_DB_USERNAME"),
      os.Getenv("APP_DB_PASSWORD"),
      os.Getenv("APP_DB_NAME"))
  a.InitializeDB(connectionString)

  a.InitializeRouter()
  customer.InitializeRoutes(&a)
  merchant.InitializeRoutes(&a)
  product.InitializeRoutes(&a)
  review.InitializeRoutes(&a)

  http.Handle("/", a.Router)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
