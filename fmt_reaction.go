package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/nlopes/slack"
)

func fmtReactions(emojiCountPerUser map[string]map[string]int, token string) (text string) {

	api := slack.New(token)

	for userID, emojiCount := range emojiCountPerUser {
		user, err := api.GetUserInfo(userID)
		if err != nil {
			fmt.Printf("[userID: %d] %s\n", userID, err)
			continue
		}
		if user.Deleted == true {
			continue
		}

		realName := user.RealName
		fmt.Printf("[%s]\t%s\n", userID, realName)
		text += "*" + realName + "*\n\t\t\t"

		sortedPairList := rankByWordCount(emojiCount)

		for _, pair := range sortedPairList {
			emoji := pair.Key
			count := pair.Value

			text += " :" + emoji + ": x" + strconv.Itoa(count)
		}
		text += "\n"

	}

	return text
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
