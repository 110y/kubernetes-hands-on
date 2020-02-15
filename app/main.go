package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to get hostname")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("I'm %s", hostname)))
	}))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "failed to launch server: %s", err.Error())
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM)
	<-c

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "failed to shutdown server: %s", err.Error())
		os.Exit(1)
	}
}
