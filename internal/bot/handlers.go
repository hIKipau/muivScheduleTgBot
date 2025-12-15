package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"muivScheduleTgBot/internal/schedule"
	"strings"
	"time"
)

func Handler(update *tgbotapi.Update,
	scheduleCourse4 *[schedule.CountOfGroups4]*schedule.Schedule,
	scheduleCourse3 *[schedule.CountOfGroups3]*schedule.Schedule,
	scheduleCourse2 *[schedule.CountOfGroups2]*schedule.Schedule,
	scheduleCourse111 *[schedule.CountOfGroups111]*schedule.Schedule,
	scheduleCourse19 *[schedule.CountOfGroups19]*schedule.Schedule) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case "start":
			msg.ReplyMarkup = ChooseCourseKeyboard
			return msg
		}
	}

	if strings.HasSuffix(update.Message.Text, "кл") {
		switch update.Message.Text {
		case "4к9кл/3к11кл":
			msg.ReplyMarkup = ChooseGroup4CKeyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "3к9кл/2к11кл":
			msg.ReplyMarkup = ChooseGroup3CKeyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "2к9кл/1к11кл":
			msg.ReplyMarkup = ChooseGroup2CKeyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "1к9кл":
			msg.ReplyMarkup = ChooseGroup1C9Keyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "1к11кл":
			msg.ReplyMarkup = ChooseGroup1C11Keyboard
			msg.Text = "Красава,теперь группу"
			return msg
		default:
			msg.Text = "/start"
			return msg
		}
	} else if strings.HasSuffix(update.Message.Text, "11") ||
		strings.HasSuffix(update.Message.Text, "09") {
		msg.ReplyMarkup = DefaultKeyboard
		msg.Text = "Нормуль, теперь че те надо"
		return msg
	}
	switch update.Message.Text {
	case "Понедельник":
		msg.Text = scheduleCourse4[0].WeekDay[0].ToMessage()
	case "Вторник":
		msg.Text = scheduleCourse4[0].WeekDay[0].ToMessage()
	case "Среда":
		msg.Text = scheduleCourse4[0].WeekDay[0].ToMessage()
	case "Четверг":
		msg.Text = scheduleCourse4[0].WeekDay[0].ToMessage()
	case "Пятница":
		msg.Text = scheduleCourse4[0].WeekDay[0].ToMessage()
	case "Суббота":
		msg.Text = scheduleCourse4[0].WeekDay[0].ToMessage()
	case "Сегодня":
		msg.Text = scheduleCourse4[0].WeekDay[getWeekDay()].ToMessage()
	case "Выбрать день":
		msg.ReplyMarkup = ChooseDayKeyboard
	case "Вернуться":
		msg.ReplyMarkup = DefaultKeyboard
	default:
		msg.Text = "Соси хуй"
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
