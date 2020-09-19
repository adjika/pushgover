package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var baseURL string = "https://api.pushover.net/1/messages.json"

type RequestBody struct {
	AppToken string `json:"token"`
	UserKey  string `json:"user"`
	Message  string `json:"message"`
}

func main() {
	var reqbody RequestBody
	reqbody.AppToken = os.Getenv("PUSHGOVER_APPTOKEN")
	reqbody.UserKey = os.Getenv("PUSHGOVER_USERKEY")
	if len(reqbody.AppToken) == 0 {
		fmt.Println("Error: No app token has been supplied via PUSHGOVER_APPTOKEN")
		return
	}
	if len(reqbody.UserKey) == 0 {
		fmt.Println("Error: No user key has been supplied via PUSHGOVER_USERKEY")
		return
	}

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Error: No message has been specified on the commandline")
		return
	}

	reqbody.Message = strings.Join(flag.Args(), " ")
	reqbodyJSON, err := json.Marshal(reqbody)
	if err != nil {
		fmt.Println("Error: Couldn't marshal reqbody into JSON")
		return
	}

	res, err := http.Post(baseURL, "application/json", bytes.NewReader(reqbodyJSON))
	if err != nil {
		fmt.Printf("Error: Couldn't post to API because of %s", err)
		return
	}
	defer res.Body.Close()
}
