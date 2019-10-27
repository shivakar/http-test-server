package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading response body: %v", err), http.StatusBadRequest)
		return
	}
	w.Write(body)
}
