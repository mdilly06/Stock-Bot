package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
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
		return
	}
	content := strings.ToLower(m.Content)
	if len(content) <= len("-") {
		return
	}
	if content[:len("-")] != "-" {
		return
	}
	if content == "-nixeus" {
		summary := ""
		summary += "```"
		summary += "Nixeus records and keeps track of your active positions. To have Nixeus start recording your position, use the following format:\n\n\t\"-in [ticker name] [amount of shares] [price]\"\n\nFor example, if you started a position in 10 shares of AAPL at $100 per share, you would enter: \n\n\t\"-in AAPL 10 100\"\n\nTo have Nixeus record the exiting of your position, use the following format:\n\n\t\"-out [ticker name] [price]\"\n\nFor example, if you exited your AAPL at $150 per share, you would enter:\n\n\t\"-out AAPL 150\""
		summary += "```"
		_, _ = s.ChannelMessageSend(m.ChannelID, summary)
	}
	if strings.HasPrefix(content, "-in") {

	}
	if strings.HasPrefix(content, "-out") {

	}

}
