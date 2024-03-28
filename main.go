package main

import (
	"log"
	"net/http"
	"os"
	"testNodes/handler"
	"testNodes/schedulerCommunicator/healthPinger"
	"testNodes/setting"
	"testNodes/src/corsController"
	"testNodes/src/logCtrlr"

	"github.com/urfave/negroni"
)

func initialization() {
	logCtrlr.Log("Initialize the agent.")

	if setting.ManagerActive {
		logCtrlr.Log("Use manager.")
		go healthPinger.Enter()
	}

	setting.ServerPort = os.Getenv("ServerPort")
}

func startServer() {
	mux := handler.CreateHandler()
	handler := negroni.Classic()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	// HTTP Server Start
	logCtrlr.Log("HTTP server start.")
	err := http.ListenAndServe(":"+setting.ServerPort, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initialization()

	startServer()
}
