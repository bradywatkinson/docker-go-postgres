// test_utils/log_utils.go

package common

import (
  "os"
  "net/http"
  "encoding/json"
  "time"
  "io"
  "io/ioutil"
  "bytes"

  "github.com/urfave/negroni"
  "github.com/sirupsen/logrus"
)

var logger = logrus.New()

// InitializeLogger sets the configuration for logrus
func InitializeLogger() {
  // Only log the warning severity or above.
  switch os.Getenv("REALM") {
  case "localdev":
    logrus.SetLevel(logrus.DebugLevel)
    logger.Level = logrus.DebugLevel
  case "test":
    logrus.SetOutput(ioutil.Discard)
    logger.Out = ioutil.Discard
  case "prod":
    logrus.SetFormatter(&logrus.JSONFormatter{})
    logger.Formatter = &logrus.JSONFormatter{}
  }
}

type loggingMiddleware struct {
}

// The middleware handler
func (l *loggingMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
      logger.Error("Error reading body: %v", err)
      RespondWithError(w, http.StatusBadRequest, "Failed to read body")
      return
  }

  start := time.Now()

  // Try to get the real IP
  remoteAddr := req.RemoteAddr
  if realIP := req.Header.Get("X-Real-IP"); realIP != "" {
    remoteAddr = realIP
  }

  entry := logger.WithFields(logrus.Fields{
    "request": req.RequestURI,
    "method":  req.Method,
    "remote":  remoteAddr,
  })

  var reqBody interface{}
  json.Unmarshal(body, &reqBody)

  // headers, _ := json.Marshal(req.Header)
  entry.WithFields(logrus.Fields{
    "headers": req.Header,
    "body":    reqBody,
  }).Info("Request Payload")

  // And now set a new body, which will simulate the same data we read:
  req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

  // Create a response wrapper:
  mrw := &myResponseWriter{
      ResponseWriter: w,
      buf:            &bytes.Buffer{},
  }

  // Call next handler, passing the response wrapper:
  next(mrw, req)

  res := w.(negroni.ResponseWriter)
  var resBody interface{}
  json.Unmarshal(mrw.buf.Bytes(), &resBody)
  latency := time.Since(start)
  entry.WithFields(logrus.Fields{
    "status":      res.Status(),
    "text_status": http.StatusText(res.Status()),
    "time_ms":     latency,
    "headers":     res.Header(),
    "body":        resBody,
  }).Info("Response Payload")
  // Now inspect response, and finally send it out:
  // (You can also modify it before sending it out!)
  if _, err := io.Copy(w, mrw.buf); err != nil {
    logger.Error("Failed to send out response: %v", err)
  }
}
