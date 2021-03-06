package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ssubedir/go-shoot-dashboard/internal/handlers"

	gohandlers "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {

	sh := handlers.NewServices()

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/", sh.Index)
	getR.HandleFunc("/enqueued", sh.Enqueued)
	getR.HandleFunc("/scheduled", sh.Scheduled)
	getR.HandleFunc("/successful", sh.Successful)
	getR.HandleFunc("/sum", sh.TaskSummary)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         ":9000",           // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Println("Starting Api service on port 9000")

		err := s.ListenAndServe()
		if err != nil {
			log.Fatal("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// serve dashboard
	log.Println("Starting Dashboard service on port 9001")
	http.ListenAndServe(":9001", http.FileServer(http.Dir("dashboard/dist/go-shoot")))

	sigs := make(chan os.Signal, 1)
	sigdone := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Signals
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		sigdone <- true
	}()

	<-sigdone
	log.Println("Got signal, Exiting service")
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
