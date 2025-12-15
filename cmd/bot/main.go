package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"muivScheduleTgBot/internal/bot"
	"muivScheduleTgBot/internal/schedule"
	"os"
)

type Course struct {
	GroupCount int
	Schedule   []schedule.Schedule
}

func main() {

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", err)
		panic(err)
	}

	TgBot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		slog.Error("Error loading .env file", err)
		panic(err)
	}

	TgBot.Debug = true

	log.Printf("Authorized on account %s", TgBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := TgBot.GetUpdatesChan(u)
	slog.Info("Get updates channel")

	scheduleCourse4 := &[schedule.CountOfGroups4]*schedule.Schedule{}
	for i := 0; i < schedule.CountOfGroups4; i++ {
		scheduleCourse4[i] = &schedule.Schedule{}
	}
	scheduleCourse3 := &[schedule.CountOfGroups3]*schedule.Schedule{}
	for i := 0; i < schedule.CountOfGroups3; i++ {
		scheduleCourse3[i] = &schedule.Schedule{}
	}
	scheduleCourse2 := &[schedule.CountOfGroups2]*schedule.Schedule{}
	for i := 0; i < schedule.CountOfGroups2; i++ {
		scheduleCourse2[i] = &schedule.Schedule{}
	}
	scheduleCourse111 := &[schedule.CountOfGroups111]*schedule.Schedule{}
	for i := 0; i < schedule.CountOfGroups111; i++ {
		scheduleCourse111[i] = &schedule.Schedule{}
	}
	scheduleCourse19 := &[schedule.CountOfGroups19]*schedule.Schedule{}
	for i := 0; i < schedule.CountOfGroups19; i++ {
		scheduleCourse19[i] = &schedule.Schedule{}
	}
	slog.Info("Starting LoadSchedule")
	err = schedule.LoadSchedule(scheduleCourse4,
		scheduleCourse3,
		scheduleCourse2,
		scheduleCourse111,
		scheduleCourse19)

	if err != nil {
		slog.Error("cant parse new schedule", err)
		panic(err)
	}
	slog.Info("LoadSchedule is successfully complete")

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := bot.Handler(&update, scheduleCourse4,
			scheduleCourse3,
			scheduleCourse2,
			scheduleCourse111,
			scheduleCourse19)

		if _, err = TgBot.Send(msg); err != nil {
			slog.Error("TgBot cant send new message", err)
			panic(err)
		}
	}
}
