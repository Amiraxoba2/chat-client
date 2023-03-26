package internal

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Read(config Config) {
	print("Chat ID you want to connect to (Ask your Admin if you don't know what this is) >> ")
	chatID, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	req, _ := http.NewRequest(http.MethodGet, config.Server+"/messages?chat="+strings.Split(chatID, "\n")[0], nil)
	client := http.Client{}
	for true {
		resp, _ := client.Do(req)
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		print("\u001B[H\u001B[2J" + string(body))
		time.Sleep(time.Second)
	}
}
