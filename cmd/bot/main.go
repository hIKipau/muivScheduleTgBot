package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"tgBotAshley/internal/schedule"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Сегодня"),
		tgbotapi.NewKeyboardButton("Выбрать день"),
	),
)
var numericKeyboardChoose = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Понедельник"),
		tgbotapi.NewKeyboardButton("Вторник"),
		tgbotapi.NewKeyboardButton("Среда"),
		tgbotapi.NewKeyboardButton("Четверг"),
		tgbotapi.NewKeyboardButton("Пятница"),
		tgbotapi.NewKeyboardButton("Суббота"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Вернуться"),
	),
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	schedule := schedule.New()
	err = schedule.ParseSchedule()
	if err != nil {
		fmt.Println("cant parse new schedule", err)
		panic(err)
	}
	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg.ReplyMarkup = numericKeyboard
			}
		} else {

			switch update.Message.Text {
			case "Сегодня":
				msg.Text = (schedule.WeekDay[getWeekDay()]).ToMessage()
			case "Выбрать день":
				msg.ReplyMarkup = numericKeyboardChoose

			case "Понедельник":
				msg.Text = (schedule.WeekDay[0]).ToMessage()
			case "Вторник":
				msg.Text = (schedule.WeekDay[1]).ToMessage()
			case "Среда":
				msg.Text = (schedule.WeekDay[2]).ToMessage()
			case "Четверг":
				msg.Text = (schedule.WeekDay[3]).ToMessage()
			case "Пятница":
				msg.Text = (schedule.WeekDay[4]).ToMessage()
			case "Суббота":
				msg.Text = (schedule.WeekDay[5]).ToMessage()
			case "Вернуться":
				msg.ReplyMarkup = numericKeyboard

			default:
				msg.Text = "Не знаю такого додик"
			}

		}

		// Send the message.
		if _, err = bot.Send(msg); err != nil {
			panic(err)
		}

	}
}

func getWeekDay() int {
	weekDay := time.Now().Weekday() - 1
	if weekDay == -1 {
		weekDay = 0
	}
	return int(weekDay)
}
