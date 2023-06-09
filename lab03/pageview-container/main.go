package main

import(
    "fmt"
    "log"
    "net/http"
)

var pageView int64

func main() {
      http.HandleFunc("/", handler)
      log.Printf("Server v2.0 starting and listening on 0.0.0.0:8080...")
      log.Fatal(http.ListenAndServe("0.0.0.0:8080",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
      log.Printf("Ping from %s", r.RemoteAddr )
      pageview++
      fmt.Fprintf(w, "[v2.0] Hello, you're visitor #%d!",pageView)
}
