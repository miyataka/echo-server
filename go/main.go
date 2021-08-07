package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/", echoProcess)
	go http.ListenAndServe(":8888", nil)

	s := <-signals
	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}

func echoProcess(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, "Hello World!")
		break
	}
}
