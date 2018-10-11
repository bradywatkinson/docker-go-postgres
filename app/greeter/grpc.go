// greeter/grpc.go

package greeter

import (
  "app/common"
)


// InitializeService registers the GreeterService
func InitializeService(a *common.App) {
  RegisterGreeterServer(a.GRPC, &GreeterService{})
}
