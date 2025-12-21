package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"muivScheduleTgBot/internal/bot"
	"muivScheduleTgBot/internal/repo"
	"muivScheduleTgBot/internal/schedule"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", err)
		panic(err)
	}
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	defer conn.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	err = conn.Ping(context.Background())
	if err != nil {
		slog.Error("unable to ping databasse")
		panic(err)
	}
	userRepo := repo.New(conn)

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

	gs := schedule.NewGlobalSchedules()

	slog.Info("Starting LoadSchedule")
	err = gs.LoadSchedule()
	if err != nil {
		slog.Error("cant parse new schedule", err)
		panic(err)
	}
	slog.Info("LoadSchedule is successfully complete")

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		if update.Message.Document != nil {
			msg := schedule.Updater(&update, TgBot)
			if _, err = TgBot.Send(msg); err != nil {
				slog.Error("TgBot cant send new message", err)
				panic(err)
			}
			err = gs.LoadSchedule()
			continue
		}

		msg := bot.Handler(context.Background(), &update, userRepo, gs)
		if _, err = TgBot.Send(msg); err != nil {
			slog.Error("TgBot cant send new message", err)
			panic(err)
		}
	}
}
