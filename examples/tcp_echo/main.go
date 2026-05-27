package main

import (
	"fmt"
	"log"
	"sync/atomic"

	"github.com/pawelgaczynski/gain"
	"github.com/pawelgaczynski/gain/logger"
	"github.com/rs/zerolog"
)

const (
	port            = 8080
	numberOfClients = 2
)

var testData = []byte("echo")

type EventHandler struct {
	server gain.Server

	logger zerolog.Logger

	overallBytesSent atomic.Uint64
}

func (e *EventHandler) OnStart(server gain.Server) { _ = "STUB: not implemented"; return }

func (e *EventHandler) OnAccept(conn gain.Conn) { _ = "STUB: not implemented"; return }

func (e *EventHandler) OnRead(conn gain.Conn, n int) { _ = "STUB: not implemented"; return }

func (e *EventHandler) OnWrite(conn gain.Conn, n int) { _ = "STUB: not implemented"; return }

func (e *EventHandler) OnClose(conn gain.Conn, err error) { _ = "STUB: not implemented"; return }

func runClients() { _ = "STUB: not implemented"; return }

func main() {
	runClients()

	err := gain.ListenAndServe(
		fmt.Sprintf("tcp://127.0.0.1:%d", port), &EventHandler{}, gain.WithLoggerLevel(logger.WarnLevel))
	if err != nil {
		log.Panic(err)
	}
}
