package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"muivScheduleTgBot/internal/bot"
	"muivScheduleTgBot/internal/schedule"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TgBot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	TgBot.Debug = true

	log.Printf("Authorized on account %s", TgBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := TgBot.GetUpdatesChan(u)

	schdle := schedule.New()
	err = schdle.LoadSchedule()

	if err != nil {
		fmt.Println("cant parse new schedule", err)
		panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		msg := bot.Handler(&update, schdle)
		if _, err = TgBot.Send(msg); err != nil {
			panic(err)
		}
	}
}
