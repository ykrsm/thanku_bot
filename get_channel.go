package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func getChannelIDs(token string) []string {

	api := slack.New(token)
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	var channelIDs []string

	for _, channel := range channels {
		// channel is of type conversation & groupConversation
		// see all available methods in `conversation.go`

		channelID := channel.ID

		channelIDs = append(channelIDs, channelID)

		fmt.Printf("[%s] %s\n", channel.ID, channel.Name)
	}

	return channelIDs
}
