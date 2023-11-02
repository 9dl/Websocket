package main

import (
	"WBS/Commands"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pterm/pterm"
	"log"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := string(p)

		parts := strings.Fields(message)
		if len(parts) == 0 {
			continue
		}

		cmdName := parts[0]
		cmdArgs := parts[1:]

		if command, found := Commands.CommandRegistry[cmdName]; found {
			command.Run(cmdArgs)
		}

	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	Clear()
	Title("9yVC0YsH7TUgsAY3UHuB8U5CDZxgq7CM")
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	WriteColoredCenteredV2("Server is {Running}", pterm.NewRGB(33, 3, 255))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
