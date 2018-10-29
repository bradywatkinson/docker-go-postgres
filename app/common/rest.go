// common/router.go

package common

import (
  "fmt"
  "net/http"
  "net/http/httputil"
  "encoding/json"
  "time"
  "io"
  "io/ioutil"
  "bytes"

  "github.com/gorilla/mux"
  "github.com/urfave/negroni"
  "github.com/sirupsen/logrus"
  // "github.com/meatballhat/negroni-logrus"
)


type JsonPanicFormatter struct{}

func (j *JsonPanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *negroni.PanicInformation) {
  e := map[string]string{
    "error": fmt.Sprintf("%s", infos.RecoveredPanic),
  }
  response, _ := json.Marshal(e)
  rw.Write(response)
}

// DefaultBefore is the default func assigned to *Middleware.Before
func LoggingBefore(entry *logrus.Entry, req *http.Request, remoteAddr string) *logrus.Entry {
  // Save a copy of this request for debugging.
  requestDump, err := httputil.DumpRequest(req, true)
  if err != nil {
    entry.Panic(err)
  }
  entry.Info(string(requestDump))
  return entry.WithFields(logrus.Fields{
    "request": req.RequestURI,
    "method":  req.Method,
    "remote":  remoteAddr,
  })
}

// DefaultAfter is the default func assigned to *Middleware.After
func LoggingAfter(entry *logrus.Entry, res negroni.ResponseWriter, latency time.Duration, name string) *logrus.Entry {
  return entry.WithFields(logrus.Fields{
    "status":      res.Status(),
    "text_status": http.StatusText(res.Status()),
    "took":        latency,
  })
}

// func loginmw(handler http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         body, err := ioutil.ReadAll(r.Body)
//         if err != nil {
//             log.Printf("Error reading body: %v", err)
//             http.Error(w, "can't read body", http.StatusBadRequest)
//             return
//         }

//         // Work / inspect body. You may even modify it!

//         // And now set a new body, which will simulate the same data we read:
//         r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

//         // Create a response wrapper:
//         mrw := &MyResponseWriter{
//             ResponseWriter: w,
//             buf:            &bytes.Buffer{},
//         }

//         // Call next handler, passing the response wrapper:
//         handler.ServeHTTP(mrw, r)

//         // Now inspect response, and finally send it out:
//         // (You can also modify it before sending it out!)
//         if _, err := io.Copy(w, mrw.buf); err != nil {
//             log.Printf("Failed to send out response: %v", err)
//         }
//     })
// }

type LoggingMiddleware struct {
}


type MyResponseWriter struct {
  http.ResponseWriter
  buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
  return mrw.buf.Write(p)
}

// The middleware handler
func (l *LoggingMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
      logger.Printf("Error reading body: %v", err)
      http.Error(w, "can't read body", http.StatusBadRequest)
      return
  }

  // And now set a new body, which will simulate the same data we read:
  req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
  logger.Info(string(requestDump))

  // Create a response wrapper:
  mrw := &MyResponseWriter{
      ResponseWriter: w,
      buf:            &bytes.Buffer{},
  }

  // Call next handler, passing the response wrapper:
  next(mrw, req)

  logger.Printf("%#v", mrw.buf.String())
  // Now inspect response, and finally send it out:
  // (You can also modify it before sending it out!)
  if _, err := io.Copy(w, mrw.buf); err != nil {
    logger.Printf("Failed to send out response: %v", err)
  }
}

// InitializeRouter setups the mux router and any additional
// middlewares
func (a *App) InitializeRouter() *negroni.Negroni {
  a.Router = mux.NewRouter()
  n := negroni.New()
  // logging := negronilogrus.NewMiddlewareFromLogger(logger, "web")
  // logging.Before = LoggingBefore
  // logging.After = LoggingAfter
  // n.Use(logging)
  n.Use(&LoggingMiddleware{})
  recovery := negroni.NewRecovery()
  recovery.Formatter = &JsonPanicFormatter{}
  n.Use(recovery)
  return n
}


// RespondWithError is a helper function to respond to a request with an error
func RespondWithError(rw http.ResponseWriter, code int, message string) {
  RespondWithJSON(rw, code, map[string]string{"error": message})
}

// RespondWithJSON is a helper function to respond to a request with a json payload
func RespondWithJSON(rw http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  rw.Header().Set("Content-Type", "application/json")
  rw.WriteHeader(code)
  rw.Write(response)
}
