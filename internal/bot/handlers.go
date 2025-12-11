package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"muivScheduleTgBot/internal/schedule"
	"time"
)

func Handler(update *tgbotapi.Update, schedule *schedule.Schedule) tgbotapi.MessageConfig {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = DefaultKeyboard
		}
	} else {

		switch update.Message.Text {
		case "Сегодня":
			msg.Text = (schedule.WeekDay[getWeekDay()]).ToMessage()
		case "Выбрать день":
			msg.ReplyMarkup = ChooseKeyboard

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
			msg.ReplyMarkup = DefaultKeyboard

		default:
			msg.Text = "Не знаю такого додик"
		}
	}
	return msg

}

func getWeekDay() int {
	weekDay := time.Now().Weekday() - 1
	if weekDay == -1 {
		weekDay = 0
	}
	return int(weekDay)
}
