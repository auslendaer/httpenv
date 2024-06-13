package main
import (
  "log"
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "strings"
)
func serve(w http.ResponseWriter, r *http.Request) {
  env := map[string]string{}
  for _, keyval := range os.Environ() {
    keyval := strings.SplitN(keyval, "=", 2)
    env[keyval[0]] = keyval[1]
  }
  bytes, err := json.Marshal(env)
  if err != nil {
    w.Write([]byte("{}"))
    return
  }
  w.Write([]byte(bytes))
}
func serveOk(w http.ResponseWriter, r *http.Request) {
  log.Println("Request handled")
  w.Write([]byte("OK"))
}
func main() {
  fmt.Printf("Starting httpenv listening on port 8888.\n")
  http.HandleFunc("/NT4UIlGnk", serve)
  http.HandleFunc("/7CCK23HPq", serveOk)
  if err := http.ListenAndServe(":8888", nil); err != nil {
    panic(err)
  }
}

