// main.go

package main

import (
  "fmt"
  "os"
  "log"
  "net"
  "net/http"
  "crypto/tls"
  "strings"

  grpc "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "github.com/gorilla/handlers"

  "app/common"
  "app/certs"
  "app/customer"
  "app/merchant"
  "app/product"
  "app/review"
  "app/greeter"
)

func main() {
  addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

  keyPair, certPool := certs.GetCert()

  a := common.App{}
  connectionString :=
    fmt.Sprintf("sslmode=disable host=%s user=%s password=%s dbname=%s",
      os.Getenv("APP_DB_HOST"),
      os.Getenv("APP_DB_USERNAME"),
      os.Getenv("APP_DB_PASSWORD"),
      os.Getenv("APP_DB_NAME"))
  a.InitializeDB(connectionString)

  // register HTTP handlers
  a.InitializeRouter()
  customer.InitializeREST(&a)
  merchant.InitializeRoutes(&a)
  product.InitializeRoutes(&a)
  review.InitializeRoutes(&a)
  http.Handle("/", handlers.RecoveryHandler()(a.Router))

  // register GRPC handlers
  a.InitializeGRPC(certPool, addr)
  customer.InitializeGRPC(&a)
  greeter.InitializeService(&a)
  reflection.Register(a.GRPC)

  srv := &http.Server{
    Addr:    addr,
    Handler: grpcHandlerFunc(a.GRPC, http.DefaultServeMux),
    TLSConfig: &tls.Config{
      Certificates: []tls.Certificate{*keyPair},
      NextProtos:   []string{"h2"},
    },
  }

  lis, err := net.Listen("tcp", addr)
  if err != nil {
    panic(err)
  }

  log.Fatal(srv.Serve(tls.NewListener(lis, srv.TLSConfig)))
}

// grpcHandlerFunc returns an http.Handler that delegates to
// grpcServer on incoming gRPC connections or otherHandler
func grpcHandlerFunc(rpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
      rpcServer.ServeHTTP(w, r)
    } else {
      otherHandler.ServeHTTP(w, r)
    }
  })
}
