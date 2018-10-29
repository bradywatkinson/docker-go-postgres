// test_utils/log_utils.go

package common

import (
  "os"
  "io/ioutil"

  "github.com/sirupsen/logrus"
)

var logger = logrus.New()

func InitializeLogger() {
  // Only log the warning severity or above.
  switch os.Getenv("REALM") {
  case "localdev":
    logrus.SetLevel(logrus.DebugLevel)
    logger.Level = logrus.DebugLevel
  case "test":
    logrus.SetOutput(ioutil.Discard)
    logger.Out = ioutil.Discard
  }
}
