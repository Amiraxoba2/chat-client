package internal

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Login(config Config) {
	print("Username >> ")
	username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	print("Password >> ")
	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	loginData := url.Values{
		"username": {strings.Split(username, "\n")[0]},
		"password": {strings.Split(password, "\n")[0]},
	}
	req, _ := http.NewRequest("POST", config.Server+"/login", strings.NewReader(loginData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 404 {
		println(Error.Render("An error occurred"))
	} else {
		config.Token = string(body)
		configBytes, _ := json.Marshal(config)
		os.WriteFile(".config.json", configBytes, os.ModePerm)
		println(Success.Render("Successfully logged in!"))
	}
}
