package pkg

import (
	
	"fmt"
    "net/http"
    "appengine"

)

func init() {
    http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    c.Infof("Requested URL: %v", r.URL)
    fmt.Fprintf(w, "<html><body> <h1> Hello Internets! </h1> </body></html>")
}