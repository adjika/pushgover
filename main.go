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
var msgLimit int = 1024

type RequestBody struct {
	AppToken string `json:"token"`
	UserKey  string `json:"user"`
	Message  string `json:"message"`

	// Optional
	Attachment string `json:"attachment,omitempty"`
	Device     string `json:"device,omitempty"`
	Priority   string `json:"priority,omitempty"`
	Title      string `json:"title,omitempty"`
	URL        string `json:"url,omitempty"`
}

type Response struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
}

func main() {
	var reqbody RequestBody
	var response Response
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
	if len(reqbody.Message) > msgLimit {
		fmt.Println("Warning: Message too long, truncating")
		reqbody.Message = reqbody.Message[:1023]
	}

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
	var buf bytes.Buffer
	_, err = buf.ReadFrom(res.Body)

	if err != nil {
		fmt.Printf("Error: Couldn't read response body")
		return
	}
	err = json.Unmarshal(buf.Bytes(), &response)
	if response.Status != 1 {
		fmt.Println("Error: Post to API unsuccessful because of: \n", response.Errors)
		return
	}
}
