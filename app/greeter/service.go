// greeter/service.go

package greeter

import (
  "context"
)

// GreeterService implements greeter.GreeterServer
type GreeterService struct{}

// SayHello implements greeter.SayHello
func (s *GreeterService) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
  return &HelloReply{Message: "Hello " + in.Name}, nil
}
