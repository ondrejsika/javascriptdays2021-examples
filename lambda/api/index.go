package handler

import (
	"fmt"
	"time"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(w, "<h1>")
	fmt.Fprintf(w, "Hello from Go! ")
	fmt.Fprintf(w, currentTime.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(w, "</h1>")
}
