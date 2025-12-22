package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"muivScheduleTgBot/internal/schedule"
	"strings"
	"time"
)

type userRepository interface {
	UpdateUserCourse(ctx context.Context, userID int64, course string) error
	UpdateUserGroup(ctx context.Context, userID int64, group string) error
	GetUserInfo(ctx context.Context, userID int64) (string, string, error)
}

func Handler(ctx context.Context,
	update *tgbotapi.Update,
	userRepo userRepository,
	gs *schedule.GlobalSchedules) tgbotapi.MessageConfig {

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
			err := userRepo.UpdateUserCourse(ctx,
				update.Message.From.ID,
				schedule.Course4s)
			if err != nil {
				msg.Text = "333: Не получилось сохранить, попробуй /start"
				return msg
			}
			msg.ReplyMarkup = ChooseGroup4CKeyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "3к9кл/2к11кл":
			err := userRepo.UpdateUserCourse(ctx,
				update.Message.From.ID,
				schedule.Course3s)
			if err != nil {
				msg.Text = "333: Не получилось сохранить, попробуй /start"
				return msg
			}
			msg.ReplyMarkup = ChooseGroup3CKeyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "2к9кл/1к11кл":
			err := userRepo.UpdateUserCourse(ctx,
				update.Message.From.ID,
				schedule.Course2s)
			if err != nil {
				msg.Text = "333: Не получилось сохранить, попробуй /start"
				return msg
			}
			msg.ReplyMarkup = ChooseGroup2CKeyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "1к9кл":
			err := userRepo.UpdateUserCourse(ctx,
				update.Message.From.ID,
				schedule.Course19s)
			if err != nil {
				msg.Text = "333: Не получилось сохранить, попробуй /start"
				return msg
			}
			msg.ReplyMarkup = ChooseGroup1C9Keyboard
			msg.Text = "Красава,теперь группу"
			return msg
		case "1к11кл":
			err := userRepo.UpdateUserCourse(ctx,
				update.Message.From.ID,
				schedule.Course111s)
			if err != nil {
				msg.Text = "333: Не получилось сохранить, попробуй /start"
				return msg
			}
			msg.ReplyMarkup = ChooseGroup1C11Keyboard
			msg.Text = "Красава,теперь группу"
			return msg
		default:
			msg.Text = "Попробуй написать /start"
			return msg
		}
	} else if strings.HasSuffix(update.Message.Text, "11") ||
		strings.HasSuffix(update.Message.Text, ".1") ||
		strings.HasSuffix(update.Message.Text, ".2") ||
		strings.HasSuffix(update.Message.Text, ".3") ||
		strings.HasSuffix(update.Message.Text, ".4") ||
		strings.HasSuffix(update.Message.Text, ".5") ||
		strings.HasSuffix(update.Message.Text, "09") {
		err := userRepo.UpdateUserGroup(ctx,
			update.Message.From.ID,
			update.Message.Text)
		if err != nil {
			msg.Text = "444: Не получилось сохранить, попробуй /start"
			return msg
		}
		msg.ReplyMarkup = DefaultKeyboard
		msg.Text = "Нормуль, теперь че те надо"
		return msg
	}

	switch update.Message.Text {
	case "Понедельник":
		course, group, err := userRepo.GetUserInfo(ctx, update.Message.From.ID)
		if err != nil {
			msg.Text = "555: Что то не работает, попробуй /start"
			return msg
		}
		msg.Text = schedule.GetSchedule(course, group, gs).WeekDay[0].ToMessage()

	case "Вторник":
		course, group, err := userRepo.GetUserInfo(ctx, update.Message.From.ID)
		if err != nil {
			msg.Text = "555: Что то не работает, попробуй /start"
			return msg
		}
		msg.Text = schedule.GetSchedule(course, group, gs).WeekDay[1].ToMessage()

	case "Среда":
		course, group, err := userRepo.GetUserInfo(ctx, update.Message.From.ID)
		if err != nil {
			msg.Text = "555: Что то не работает, попробуй /start"
			return msg
		}
		msg.Text = schedule.GetSchedule(course, group, gs).WeekDay[2].ToMessage()

	case "Четверг":
		course, group, err := userRepo.GetUserInfo(ctx, update.Message.From.ID)
		if err != nil {
			msg.Text = "555: Что то не работает, попробуй /start"
			return msg
		}
		msg.Text = schedule.GetSchedule(course, group, gs).WeekDay[3].ToMessage()

	case "Пятница":
		course, group, err := userRepo.GetUserInfo(ctx, update.Message.From.ID)
		if err != nil {
			msg.Text = "555: Что то не работает, попробуй /start"
			return msg
		}
		msg.Text = schedule.GetSchedule(course, group, gs).WeekDay[4].ToMessage()

	case "Суббота":
		course, group, err := userRepo.GetUserInfo(ctx, update.Message.From.ID)
		if err != nil {
			msg.Text = "555: Что то не работает, попробуй /start"
			return msg
		}
		msg.Text = schedule.GetSchedule(course, group, gs).WeekDay[5].ToMessage()

	case "Сегодня":
		course, group, err := userRepo.GetUserInfo(ctx, update.Message.From.ID)
		if err != nil {
			msg.Text = "555: Что то не работает, попробуй /start"
			return msg
		}
		msg.Text = schedule.GetSchedule(course, group, gs).WeekDay[getWeekDay()].ToMessage()

	case "Выбрать день":
		msg.ReplyMarkup = ChooseDayKeyboard
	case "Вернуться":
		msg.ReplyMarkup = DefaultKeyboard
	case "Сбросить":
		msg.ReplyMarkup = ChooseCourseKeyboard
		return msg
	case "Назад":
		msg.ReplyMarkup = ChooseCourseKeyboard
		return msg

	default:
		msg.Text = "666:Напиши /start"
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
