package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func SendMessage(message string, sendAsFile bool) {
	if sendAsFile || len(message) > 4095 {
		fileName := WriteTextToFile(message)
		SendFile(fileName)
	} else {
		client := createClient()
		msg := tgbotapi.NewMessage(ReadChatId(), message)
		_, err := client.Send(msg)
		if err != nil {
			log.Fatalf("Couldn't send message to telegram: %s", err)
		}
	}
}

func SendFile(fileName string) {
	client := createClient()
	msg := tgbotapi.NewDocumentUpload(ReadChatId(), fileName)
	_, err := client.Send(msg)
	if err != nil {
		log.Fatalf("Couldn't send file to telegram: %s", err)
	}
}

func createClient() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(ReadBotToken())
	if err != nil {
		log.Fatalf("Couldn't create telegram client: %s", err)
	}
	return bot
}
