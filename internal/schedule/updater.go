package schedule

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	ScheduleDir  = "./currentSchedule"
	ScheduleFile = "newSchedule.xlsx"
)

func Updater(update *tgbotapi.Update, TgBot *tgbotapi.BotAPI) tgbotapi.MessageConfig {
	caption := update.Message.Caption

	if strings.HasPrefix(caption, "/update") {
		parts := strings.Fields(caption)
		if len(parts) < 2 {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, укажите пароль после команды: /update password")
		}

		password := parts[1]
		if password != os.Getenv("UPDATE_SCHEDULE") {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Неверный пароль!")
		}

		// гарантируем, что папка существует
		if err := os.MkdirAll(ScheduleDir, 0755); err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Проблемы со стороны сервера при создании папки")

		}
		schedulePath := filepath.Join(ScheduleDir, ScheduleFile)

		// получаем файл от Telegram
		file, err := TgBot.GetFile(tgbotapi.FileConfig{
			FileID: update.Message.Document.FileID,
		})
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось получить файл от Telegram")
		}

		// скачиваем файл
		resp, err := http.Get(file.Link(TgBot.Token))
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при скачивании файла")
		}
		defer resp.Body.Close()

		// создаём файл с НУЖНЫМ ИМЕНЕМ
		out, err := os.Create(schedulePath)
		if err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось сохранить файл")

		}
		defer out.Close()

		if _, err := io.Copy(out, resp.Body); err != nil {
			return tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при записи файла")
		}
		return tgbotapi.NewMessage(update.Message.Chat.ID, "✅ Расписание успешно обновлено")

	} else {
		return tgbotapi.NewMessage(update.Message.Chat.ID, "Используйте команду /update в подписи к файлу.")
	}

}
