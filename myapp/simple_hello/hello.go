// myapp
package myapp

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<b><h1>Hello World & Anand!</h1></b>")
}
