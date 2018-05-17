// common/router.go

package common

import (
  "crypto/x509"
  context "golang.org/x/net/context"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
  "github.com/grpc-ecosystem/go-grpc-middleware"
  "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
  "github.com/grpc-ecosystem/go-grpc-middleware/validator"
  "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
  "github.com/grpc-ecosystem/go-grpc-middleware/tags"

  "github.com/sirupsen/logrus"
)

func payloadLogDecider(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
  return true
}

// InitializeGRPC sets up the grpc server
func (a *App) InitializeGRPC(certPool *x509.CertPool, addr string) {

  interceptorLogger := logrus.NewEntry(logger)
  grpc_logrus.ReplaceGrpcLogger(interceptorLogger)

  opts := []grpc.ServerOption{
    grpc.Creds(credentials.NewClientTLSFromCert(certPool, addr)),
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
      grpc_ctxtags.UnaryServerInterceptor(),
      grpc_logrus.UnaryServerInterceptor(interceptorLogger),
      grpc_logrus.PayloadUnaryServerInterceptor(interceptorLogger, payloadLogDecider),
      grpc_validator.UnaryServerInterceptor(),
      grpc_recovery.UnaryServerInterceptor(),
    )),
  }

  a.GRPC = grpc.NewServer(opts...)
}
