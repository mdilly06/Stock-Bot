package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var api = "http://chart.whiteurl.com/AppChart/chart.php?c=1&app=Chart&v=1.3&st=5&s=AAPL&chartType=candlestick&ti=RSI,MACD,AccDist,CMF&sh=667&period=1y"

var (
	token string
)

func main() {

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
	}

	discord.AddHandler(returnTech)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	<-make(chan struct{})
}

func returnTech(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.Ready.User.Username == m.Author.Username {
		return
	}
}

func readFile(filename string) {

}
