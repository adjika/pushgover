package main

import "os"

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
}
