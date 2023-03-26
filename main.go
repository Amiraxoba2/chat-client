package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Amiraxoba2/chat-client/internal"
	"os"
	"strings"
)

func main() {
	if _, err := os.Open(".config.json"); errors.Is(err, os.ErrNotExist) {
		println("Client is not configured.")
		print("Server address: ")
		server, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		configBytes, _ := json.Marshal(internal.Config{
			Server: strings.Split(server, "\n")[0],
			Token:  "",
		})
		os.WriteFile(".config.json", configBytes, os.ModePerm)
	}

	commands := map[string]func(config internal.Config){
		"login":    internal.Login,
		"register": internal.Register,
		"read":     internal.Read,
		"send":     internal.Send,
		"exit": func(config internal.Config) {
			os.Exit(0)
		},
	}
	for true {
		fmt.Printf("Actions: (%s %s %s %s %s) >> ", internal.Important.Render("exit"), internal.Important.Render("login"), internal.Important.Render("register"), internal.Important.Render("read"), internal.Important.Render("send"))
		commandRaw, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command := strings.Split(commandRaw, "\n")[0]
		configBytes, _ := os.ReadFile(".config.json")
		var config internal.Config
		json.Unmarshal(configBytes, &config)
		if _, exists := commands[command]; !exists {
			println(internal.Error.Render("This command wasn't found"))
		} else {
			commands[command](config)
		}
	}
}
