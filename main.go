package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/tealeg/xlsx"
)

var (
	token string
	botID string
)

func main() {
	token = "MzAwNDYxMTg0Mjg3NTA2NDMz.C8sx2g.zdC9Qig7A5pcqhk54a - bDY4h6a8"
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
	fmt.Println("I have awoken :woke:")

	<-make(chan struct{})
	return
}

func returnTech(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author
	if user.ID == botID || user.Bot {
		return
	}
	content := m.Content
	if len(content) <= len("-") {
		return
	}
	if content[:len("-")] != "-" {
		return
	}
	if content == "-nixeus" {
		printHelp()
	}
	xlsxFile, err := xlsx.OpenFile("C:/Users/Sean/go/src/discord_bot/stats_sheet.xlsx")
	if err != nil {
		fmt.Println("error opening excel file, ", err)
		return
	}
	if strings.HasPrefix(content, "-in") {
		line := strings.Split(content, " ")
		sheet := xlsxFile.Sheets[0]
		row := sheet.AddRow()
		for i := 0; i < 12; i++ {
			row.AddCell()
		}
		row.Cells[0].SetString(user.Username)
		row.Cells[1].SetString("=NOW()")
		row.Cells[2].SetString(line[1])
		row.Cells[3].SetString(line[2])
		row.Cells[3].SetString(line[3])
	}
	if strings.HasPrefix(content, "-out") {
		line := strings.Split(content, " ")
		sheet := xlsxFile.Sheets[0]
		for _, row := range sheet.Rows {
			if row.Cells[0].Value == user.Username && row.Cells[2].Value == line[1] {
				row.Cells[10].SetString("=NOW()")
			}
		}
	}

}

func printHelp() {
	fmt.Println("Nixeus records and keeps track of your active positions")
	fmt.Println("To have Nixeus start recording your position, use the following format:")
	fmt.Println("\t\"-in [ticker name] [amount of shares] [price]\"")
	fmt.Println("For example, if you started a position in 10 shares of AAPL at $100 per share, you would enter:")
	fmt.Println("\t\"-in AAPL 10 100\"")
	fmt.Println("To have Nixeus record the exiting of your position, use the following format:")
	fmt.Println("\t\"-out [ticker name] [amount of shares] [price]\"")
	fmt.Println("For example, if you exited a position of 10 shares of AAPL at $150 per share, you would enter:")
	fmt.Println("\t\"-out AAPL 10 150\"")
}
