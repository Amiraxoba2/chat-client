package internal

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Register(config Config) {
	print("Username >> ")
	username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	print("Password >> ")
	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	loginData := url.Values{
		"username": {strings.Split(username, "\n")[0]},
		"password": {strings.Split(password, "\n")[0]},
	}
	req, _ := http.NewRequest("POST", config.Server+"/register", strings.NewReader(loginData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	println(Success.Render(string(body)))
}
