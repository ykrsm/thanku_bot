package main

import (
	"fmt"
	"strconv"

	"github.com/nlopes/slack"
)

func getReaction(emojiCountPerUser map[string]map[string]int, channelID string, token string, startTs int64, endTs int64) map[string]map[string]int {

	api := slack.New(token)
	params := slack.NewHistoryParameters()
	params.Oldest = strconv.FormatInt(startTs, 10)
	// params.Latest = strconv.FormatInt(endTs, 10)
	params.Count = 2000

	history, err := api.GetChannelHistory(channelID, params)

	if err != nil {
		fmt.Printf("%s\n", err)
		return emojiCountPerUser
	}
	// fmt.Printf("%+v\n", history)

	for _, message := range history.Messages {
		msg := message.Msg
		userID := msg.User
		reactions := msg.Reactions

		if len(reactions) <= 0 {
			// https://www.dotnetperls.com/substring-go
			// runes := []rune(msg.Text)
			// first10Unicode := runes[0:10]
			// fmt.Printf("[SKIPPED] no reactions [%s...]\n", string(first10Unicode))
			continue
		}

		// fmt.Printf("%+v", userID)

		for _, reaction := range reactions {
			emoji := reaction.Name
			count := reaction.Count
			fmt.Printf("[EMOJI] [%s]\n", emoji)

			if _, exist := emojiCountPerUser[userID]; exist == false {
				// if user doest not exist in map
				// emojiCountPerUser[userID] = map[string]int{emoji: 1}
				emojiCountPerUser[userID] = make(map[string]int)
			}

			if _, exist := emojiCountPerUser[userID][emoji]; exist == false {
				// if emoji doest not exist in map
				emojiCountPerUser[userID][emoji] = count
			} else {
				// if emoji exists in map
				emojiCountPerUser[userID][emoji] += count
			}

		}
	}

	return emojiCountPerUser
}
