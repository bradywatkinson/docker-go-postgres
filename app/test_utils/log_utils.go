// test_utils/log_utils.go

package testutils

import (
  "log"
  "os"
)

// Log only logs in the local dev environment
func Log(message string) {
  if os.Getenv("REALM") == "localdev" {
    log.Print(message)
  }
}
