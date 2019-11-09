package handlers

import (
	"net/http"
	"net/http/httputil"
)

func Req(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("RemoteAddr", r.RemoteAddr)
	c, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(c)
}
