package tools

import (
	"fmt"
	//"github.com/makiuchi-d/gozxing/multi/qrcode"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
	"strings"
)

var (
	bot        *telego.Bot
	botHandler *th.BotHandler
)

func init() {
	var err error

	//load .env file
	//err = godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	//get bot token
	botToken := os.Getenv(BotToken)
	fmt.Println(botToken)

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
	botHandler.HandleMessage(handleBotMessage)
	botHandler.HandleCallbackQuery(handleBotCallback)

	defer botHandler.Stop()
	botHandler.Start()
}
func handleBotMessage(_ *telego.Bot, message telego.Message) {
	//get command
	command, _ := tu.ParseCommand(message.Text)
	if command == "start" {
		//send hello message
		sendMessage(message.Chat.ID, HelloMessage)
	} else {
		configUrl := getConfigUrlFromMessage(message)
		serverInfo, email, err := GetUserEmailFromConfigURL(configUrl)
		if err != nil {
			sendMessage(message.Chat.ID, convertErrorMessage(err.Error()))
		}

		inlineKeyboard := retryInlineButton(email, serverInfo.IP)
		sendMessage(message.Chat.ID, getClientTraffic(configUrl), inlineKeyboard)
	}
}
func handleBotCallback(_ *telego.Bot, query telego.CallbackQuery) {
	queries := strings.Split(query.Data, "###")

	if len(queries) > 1 {
		ip := strings.Split(queries[1], ":")[0]
		email := strings.Split(queries[1], ":")[1]
		switch queries[0] {
		case "retry_callback":
			inlineKeyboard := retryInlineButton(email, ip)
			sendMessage(query.Message.Chat.ID, getClientTrafficByEmail(ip, email), inlineKeyboard)
		}
	}
}
func sendMessage(chatID int64, messageText string, replyMarkup ...telego.ReplyMarkup) {
	_chatID := tu.ID(chatID)
	params := telego.SendMessageParams{
		ChatID:    _chatID,
		Text:      messageText,
		ParseMode: "Markdown",
	}
	if len(replyMarkup) > 0 {
		params.ReplyMarkup = replyMarkup[0]
	}

	_, err := bot.SendMessage(&params)
	if err != nil {
		fmt.Println("Error sending telegram message :", err)
	}
}
func retryInlineButton(email string, ip string) *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(RetryCalcClientUsage).WithCallbackData("retry_callback###" + ip + ":" + email),
		),
	)
}
