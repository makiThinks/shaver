// index
package handler

import (
	"fmt"
	"net/http"
)

// test handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to url shaver!")
}
