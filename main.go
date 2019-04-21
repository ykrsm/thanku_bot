package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Not enough argument\n")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	oauthToken := os.Getenv("OAUTH_TOKEN")

	// os.Args[1] specify if its prod or dev
	arg := os.Args[1]
	var channelID string
	switch arg {
	case "-p":
		fmt.Printf("PRODUCTION MODE\n")
		channelID = os.Getenv("PROD_CHANNEL")

	case "-d":
		fmt.Printf("DEVELOPMENT MODE\n")
		channelID = os.Getenv("DEV_CHANNEL")

	default:
		os.Exit(1)
	}

	channelIDs := getChannelIDs(botToken)

	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	// firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	firstOfMonth := time.Date(currentYear, currentMonth-1, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	startTs := firstOfMonth.Unix()
	endTs := lastOfMonth.Unix()

	emojiCountPerUser := make(map[string]map[string]int)

	for _, channelID := range channelIDs {
		emojiCountPerUser = getReaction(emojiCountPerUser, channelID, oauthToken, startTs, endTs)
	}

	emojiCountPerUserText := fmtReactions(emojiCountPerUser, botToken)

	fmt.Printf("%s", emojiCountPerUserText)

	startDateText := firstOfMonth.Format("1月2日")
	endDateText := lastOfMonth.Format("1月2日")
	bazaruText := startDateText + "から" + endDateText + "までのリアクションでござ～る。\n"
	emojiCountPerUserText = bazaruText + emojiCountPerUserText

	postMsg(emojiCountPerUserText, channelID, botToken)
	fmt.Printf("SUCCESS data collected from %s to %s\n", firstOfMonth, now)
}
