package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main(){
	r := chi.NewRouter()
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	addr := ":8080"
	if port:= os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	fmt.Printf("listening on %s\n", addr)
	http.ListenAndServe(addr, 	r)
}
