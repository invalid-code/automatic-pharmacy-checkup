package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	PAGE_ID     = 0
	PSID    int = 0
)

func main() {
	godotenv.Load()
	reqBody := strings.NewReader(`
{
	"messaging_type": "RESPONSE",
	"recipient": {
		"id": ""
	},
	"message": {
		"text": "hello"
	}
}
`)
	resp, err := http.Post(fmt.Sprintf("https://graph.facebook.com/v22.0/me/messages?access_token=%v", os.Getenv("ACCESS_TOKEN")), "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		fmt.Println("sent")
	} else {
		fmt.Println("failed to send")
	}
}
