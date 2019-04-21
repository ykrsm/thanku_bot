package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func postMsg(text string, channelID string, token string) {

	api := slack.New(token)

	channelID, timestamp, err := api.PostMessage(channelID,
		slack.MsgOptionText(text, false),
		slack.MsgOptionAttachments())

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}
