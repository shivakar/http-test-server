package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	response := []byte(fmt.Sprintf("Hello, World!! - %s", time.Now()))
	w.Write(response)
}
