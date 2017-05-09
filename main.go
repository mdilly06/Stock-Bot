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
	if user.ID == botID || user.Bot {
		help(s, m)
		return
	}
	content := strings.ToLower(m.Content)
	else if len(content) <= len("-") {
		help(s, m)
		return
	}
	else if content[:len("-")] != "-" {
		help(s, m)
		return
	}
	else if content == "-nixeus" {
		summary := ""
		summary += "```"
		summary += "Nixeus charts and provides fundamental analysis for publicly traded companies. To have Nixeus chart a company for you, use the following format:\n\n\t\"-Nixeus chart [ticker name] [time frame]\"\n\nFor example, if you wanted to chart AAPL on a 1 year time frame, you would enter: \n\n\t\"-Nixeus chart AAPL 1y\"\n\nTime frames include: daily (1d, 3d, 7d, 14d), monthly (1m, 2m, 6m, 9m), and yearly (1y, 2y, 5y, 10y)\n\nTo have Nixeus provide fundamental analysis, use the following format:\n\n\t\"-Nixeus FA [ticker name]\"\n\nFor example, if you wanted fundamental info on AAPL, you would enter:\n\n\t\"-Nixeus FA AAPL\""
		summary += "```"
		_, _ = s.ChannelMessageSend(m.ChannelID, summary)
	}
	else if content[1] == "chart"{
		//URL : https://chart.finance.yahoo.com/z?s=AAPL&t=6m&q=l&l=on&z=s&p=m50,m200
		url := "https://chart.finance.yahoo.com/z?s=" + content[2] + "&t=" + content[3] + "&q=l&l=on&z=s&p=m50,m200"
	}
	else if content[1] == "fa" {

	}

}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	help_message := "I did not understand your input. For help type:\n\n```-Nixeus```" 
	s.ChannelMessageSend(m.ChannelID, help_message)

}