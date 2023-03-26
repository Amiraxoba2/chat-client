package internal

import (
	"bufio"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Send(config Config) {
	print("Chat ID you want to connect to (Ask your Admin if you don't know what this is) >> ")
	chatID, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	for true {
		print("\u001B[H\u001B[2JMessage >> ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		sendForm := url.Values{
			"message": {strings.Split(message, "\n")[0]},
			"token":   {config.Token},
		}
		req, _ := http.NewRequest(http.MethodPost, config.Server+"/send?chat="+strings.Split(chatID, "\n")[0], strings.NewReader(sendForm.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := http.Client{}
		client.Do(req)
	}
}
