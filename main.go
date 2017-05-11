package main

import (
	"fmt"
	"strings"

	"discordgo"
)

var (
	botID string
)

func main() {
	token := "MzAwNDkzODY2MjY2NTI1Njk2.C8wM0A.G4rGn_3Ml8XIH1yXEC-QDGrVZW0"
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	u, err := discord.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	botID = u.ID

	discord.AddHandler(returnTech)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	<-make(chan struct{})
	return
}

func returnTech(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author
	content := strings.ToLower(m.Content)
	if user.ID == botID || user.Bot {
		return
	}
	if len(content) <= len("-") {
		return
	}
	if content[:len("-")] != "-" {
		return
	}
	if content == "-nixeus" {
		summary := "```Nixeus charts and provides fundamental analysis for publicly traded companies. To have Nixeus chart a company for you, use the following format:\n\n\t\"-Nixeus chart [ticker name] [time frame]\"\n\nFor example, if you wanted to chart AAPL on a 1 year time frame, you would enter: \n\n\t\"-Nixeus chart AAPL 1y\"\n\nTime frames include: daily (today, 3d, 7d, 14d), monthly (1m, 2m, 6m, 9m), and yearly (1y, 2y, 5y, 10y)\n\nTo have Nixeus provide fundamental analysis, use the following format:\n\n\t\"-Nixeus FA [ticker name]\"\n\nFor example, if you wanted fundamental info on AAPL, you would enter:\n\n\t\"-Nixeus FA AAPL\"```"
		_, _ = s.ChannelMessageSend(m.ChannelID, summary)
		return
	}
	array := strings.Fields(content)
	if strings.HasPrefix(content, "-nixeus") {
		if array[1] == "chart" {
			// URL : https://chart.finance.yahoo.com/z?s=AAPL&t=6m&q=c&l=off&z=l&p=m50,m200,v&a=m26-12-9,r14
			url := "https://chart.finance.yahoo.com/z?s=" + array[2] + "&t=" + array[3] + "&q=c&l=off&z=l&p=m50,m200,v&a=m26-12-9,r14"
			s.ChannelMessageSend(m.ChannelID, url)
		} else {
			help(s, m)
		}
	}

}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "I did not understand your input. For help type\n\n```-Nixeus```"
	s.ChannelMessageSend(m.ChannelID, helpMessage)

}
