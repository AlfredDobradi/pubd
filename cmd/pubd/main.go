package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alfreddobradi/activitypub/activitypub"
	"github.com/alfreddobradi/activitypub/activitypub/handler"
	"github.com/alfreddobradi/activitypub/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

func main() {
	if errors := config.ParseEnv(); len(errors) > 0 {
		for key, err := range errors {
			fmt.Printf("ERROR: Failed to set config %s: %v", key, err)
		}
		os.Exit(100)
	}

	r := mux.NewRouter()

	r.HandleFunc("/~{user}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(activitypub.Owner()); err != nil {
			log.Printf("Failed to encode owner: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})

	r.HandleFunc("/.well-known/webfinger", handler.WebfingerHandler)

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spew.Dump(r)
	})

	s := &http.Server{
		Addr:         config.BindAddress(),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		s.ListenAndServe()
	}()

	<-sig
	s.Shutdown(context.Background())
	os.Exit(0)
}
