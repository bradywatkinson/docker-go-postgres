// common/router.go

package common

import (
  "crypto/x509"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
  "github.com/grpc-ecosystem/go-grpc-middleware"
  "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
  "github.com/grpc-ecosystem/go-grpc-middleware/validator"
)

// InitializeGRPC sets up the grpc server
func (a *App) InitializeGRPC(certPool *x509.CertPool, addr string) {
  opts := []grpc.ServerOption{
    grpc.Creds(credentials.NewClientTLSFromCert(certPool, addr)),
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
      grpc_validator.UnaryServerInterceptor(),
      grpc_recovery.UnaryServerInterceptor(),
    )),
  }

  a.GRPC = grpc.NewServer(opts...)
}
