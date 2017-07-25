package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Slack struct {
	Text string `json:"text"`
}

func main() {

	var text string
	var err error
	var webHookURL string = os.Getenv("SLACKURL")

	flag.Parse()
	if flag.NArg() > 1 {
		err = fmt.Errorf("too many args")
	} else if flag.NArg() == 0 {
		err = fmt.Errorf("error ex: slack_client メッセージ")
	} else {
		text = strings.TrimSpace(flag.Arg(0))
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	params, _ := json.Marshal(Slack{text})
	data := url.Values{"payload": {string(params)}}

	_, perr := http.PostForm(webHookURL, data)
	if err != nil {
		fmt.Println(perr)
		return
	}
}
