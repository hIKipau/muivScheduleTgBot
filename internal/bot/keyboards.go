package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var DefaultKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Сегодня"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Выбрать день"),
		tgbotapi.NewKeyboardButton("Сбросить"),
	),
)
var ChooseDayKeyboard = tgbotapi.NewReplyKeyboard(
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
var ChooseCourseKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4к9кл/3к11кл"),
		tgbotapi.NewKeyboardButton("3к9кл/2к11кл"),
		tgbotapi.NewKeyboardButton("2к9кл/1к11кл"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1к9кл"),
		tgbotapi.NewKeyboardButton("1к11кл"),
	),
)

var ChooseGroup4CKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ИСП/Р-22-09\nИСП/Р-23-11"),
		tgbotapi.NewKeyboardButton("РЕК-22-09\nРЕК-23-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ИСП/П-22-09"),
		tgbotapi.NewKeyboardButton("ИСП/П-23-11"),
		tgbotapi.NewKeyboardButton("ДО-23-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад "),
	),
)

var ChooseGroup3CKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ФИН-23-09\nФИН-24-11"),
		tgbotapi.NewKeyboardButton("БАД-24-11"),
		tgbotapi.NewKeyboardButton("ТД-24-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ОДЛ-23-09\nОДЛ-24-11"),
		tgbotapi.NewKeyboardButton("КОМ-23-09"),
		tgbotapi.NewKeyboardButton("РЕК-23-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ЮР/со-24-11"),
		tgbotapi.NewKeyboardButton("ПСО-23-09.1"),
		tgbotapi.NewKeyboardButton("ПСО-23-09.2"),
		tgbotapi.NewKeyboardButton("ДИЗ-24-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ИСП/р-23-09\nИСП/р-24-11"),
		tgbotapi.NewKeyboardButton("РЕК-24-11"),
		tgbotapi.NewKeyboardButton("ИСП/п-23-09.1"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ИСП/п-23-09.2\nИСП/п-24-11"),
		tgbotapi.NewKeyboardButton("ФК-24-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var ChooseGroup2CKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("БАД-24-09\nБАД-25-11"),
		tgbotapi.NewKeyboardButton("КСК-24-09\nКСК-25-11"),
		tgbotapi.NewKeyboardButton("БУХ-24-09"),
	),
	tgbotapi.NewKeyboardButtonRow(

		tgbotapi.NewKeyboardButton("ЮР/со-24-09\nЮР/со-25-11"),
		tgbotapi.NewKeyboardButton("ЮР/са-24-09\nЮР/са-25-11"),
		tgbotapi.NewKeyboardButton("РЕК-24-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ОДЛ-24-09"),
		tgbotapi.NewKeyboardButton("ФИН-24-09"),
		tgbotapi.NewKeyboardButton("ТД-24-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ЮР/со-24-09\nЮР/со-25-11"),
		tgbotapi.NewKeyboardButton("ЮР/са-24-09\nЮР/са-25-11"),
		tgbotapi.NewKeyboardButton("РЕК-24-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ИСП/П-24-09"),
		tgbotapi.NewKeyboardButton("ИСП/Р-24-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var ChooseGroup1C11Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ФИН-25-11"),
		tgbotapi.NewKeyboardButton("БУХ-25-11"),
		tgbotapi.NewKeyboardButton("ТД-25-11"),
		tgbotapi.NewKeyboardButton("ОДЛ-25-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ОИБАС-25-11"),
		tgbotapi.NewKeyboardButton("ИСП/П-25-11"),
		tgbotapi.NewKeyboardButton("ИСП/Р-25-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("РЕК-25-11"),
		tgbotapi.NewKeyboardButton("ДИЗ-25-11"),
		tgbotapi.NewKeyboardButton("ФК-25-11"),
		tgbotapi.NewKeyboardButton("ПВНК-25-11"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)

var ChooseGroup1C9Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("БАД-25-09"),
		tgbotapi.NewKeyboardButton("БУХ-25-09"),
		tgbotapi.NewKeyboardButton("ФИН-25-09"),
		tgbotapi.NewKeyboardButton("КСК-25-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ИСП/П-25-09.1"),
		tgbotapi.NewKeyboardButton("ИСП/П-25-09.2"),
		tgbotapi.NewKeyboardButton("ИСП/Р-25-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ТД-25-09.1"),
		tgbotapi.NewKeyboardButton("ТД-25-09.2"),
		tgbotapi.NewKeyboardButton("ОДЛ-25-09.1"),
		tgbotapi.NewKeyboardButton("ОДЛ-25-09.2"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ЮР/са-25-09.1"),
		tgbotapi.NewKeyboardButton("ЮР/са-25-09.2"),
		tgbotapi.NewKeyboardButton("ЮР/со-25-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("РЕК-25-09.1"),
		tgbotapi.NewKeyboardButton("РЕК-25-09.2"),
		tgbotapi.NewKeyboardButton("РЕК-25-09.3"),
		tgbotapi.NewKeyboardButton("ПВНК-25-09"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	),
)
