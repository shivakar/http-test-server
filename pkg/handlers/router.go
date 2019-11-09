package handlers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/echo", Echo).Methods("GET", "PUT", "POST", "DELETE")
	r.HandleFunc("/ping", Ping).Methods("GET")
	r.HandleFunc("/req", Req).Methods("GET")
	return r
}
