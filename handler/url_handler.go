// url service handlers
package handler

import (
	"fmt"
	"net/http"

	"maki.io/demo/shaver/domain"
)

const (
	urlErrorMessage          = "Url service not implemented for this method"
	urlParameterErrorMessage = "The url parameter is invalid"
)

// url-service handler
func UrlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		urlGetHandler(w, r)
	case http.MethodPost:
		urlSaveHandler(w, r)
	default:
		http.Error(w, urlErrorMessage, http.StatusMethodNotAllowed)
	}
}

// save the real url and return a short one
// indempotent
// only support post method
func urlSaveHandler(w http.ResponseWriter, r *http.Request) {
	// check the input
	realUrl := r.FormValue("realUrl")
	if !validateUrl(realUrl) {
		http.Error(w, urlParameterErrorMessage, http.StatusMethodNotAllowed)
		return
	}
	// get and check the result
	shortUrl, err := domain.ToShortUrl(realUrl)
	if err == nil {
		fmt.Fprintln(w, shortUrl)
	} else {
		http.Error(w, err.Error(), http.StatusFailedDependency)
	}
}

// get the real url from a short one
// only support get method
func urlGetHandler(w http.ResponseWriter, r *http.Request) {
	// check the input
	shortUrl := r.FormValue("shortUrl")
	if !validateUrl(shortUrl) {
		http.Error(w, urlParameterErrorMessage, http.StatusMethodNotAllowed)
		return
	}
	// get and check the result
	realUrl, err := domain.ToRealUrl(shortUrl)
	if err == nil {
		fmt.Fprintln(w, realUrl)
	} else {
		http.Error(w, err.Error(), http.StatusFailedDependency)
	}

}

// TODO: check if the input is a real url
func validateUrl(url string) bool {
	if len(url) > 0 {
		return true
	} else {
		return false
	}
}
