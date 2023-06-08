package tools

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"log"
	"os"
)

var (
	bot        *telego.Bot
	botHandler *th.BotHandler
)

func init() {
	var err error

	//load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//get bot token
	botToken := os.Getenv("BOT_TOKEN")

	// create new bot
	bot, err = telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	botHandler, _ = th.NewBotHandler(bot, updates)
}

func Start() {
	botHandler.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		//get command
		command, _ := tu.ParseCommand(message.Text)
		if command == "start" {
			//send hello message
			sendMessage(message, HelloMessage)
		} else {
			log.Println("I got config url")
			sendMessage(message, getClientTraffic(message.Text))
		}
	})

	defer botHandler.Stop()
	botHandler.Start()
}
func sendMessage(message telego.Message, messageText string) {
	chatID := tu.ID(message.Chat.ID)
	params := telego.SendMessageParams{
		ChatID:    chatID,
		Text:      messageText,
		ParseMode: "Markdown",
	}

	_, err := bot.SendMessage(&params)
	if err != nil {
		fmt.Println("Error sending telegram message :", err)
	}
}
