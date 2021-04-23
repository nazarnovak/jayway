package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	pkgRobot "github.com/nazarnovak/jayway/backend/pkg/robot"
	pkgRoom "github.com/nazarnovak/jayway/backend/pkg/room"
)

var (
	room  pkgRoom.Room
	robot *pkgRobot.Robot
)

func main() {
	cliPtr := flag.Bool("cli", false, "Set this to true to run the application in cli mode.")
	serverPtr := flag.Bool("server", false, "Set this to true to run the application in server mode.")
	flag.Parse()

	cliMode := *cliPtr
	serverMode := *serverPtr

	if (!cliMode && !serverMode) || (cliMode && serverMode) {
		fmt.Printf("Error: Please specify one mode to run the application\n\n")
		flag.Usage()
		return
	}

	if cliMode {
		if err := handleCLIMode(); err != nil {
			fmt.Printf("\nError: %s\n", err)
			return
		}

		return
	}

	if serverMode {
		port := ":8080"

		srv := &http.Server{
			Addr:    port,
			Handler: router(),
		}

		fmt.Println("Running server on port", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(fmt.Errorf("Error starting server: %s", err.Error()))
		}
	}
}
