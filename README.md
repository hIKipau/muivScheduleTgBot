# Telegram Schedule Bot

**Telegram Schedule Bot** — это бот для Telegram, который показывает расписание занятий в удобном виде прямо в чате.  

Бот написан на **Go** и умеет:  
- Показывать расписание на текущий день.  
- Позволяет выбрать любой день недели и увидеть расписание.  
- Автоматически парсит Excel-файл с расписанием.  

---

## Технологии
- Go (Golang)  
- Telegram Bot API (`github.com/go-telegram-bot-api/telegram-bot-api/v5`)  
- Excel parsing (`github.com/tealeg/xlsx/v3`)  
- `.env` для хранения конфиденциальных данных  

---

## Установка и запуск

1. Клонируйте репозиторий:

git clone https://github.com/hIKipau/muivScheduleTGbot.git
cd muivScheduleTGbot

Создайте файл .env в корне проекта и добавьте ваш токен бота:
BOT_TOKEN=ваш_токен_бота

Убедитесь, что Excel-файл с расписанием находится в папке currentSchedule:
currentSchedule/newSchedule.xlsx

Запустите бота:
go run ./cmd/bot/main.go

Откройте Telegram и найдите вашего бота по username, затем используйте команду /start для начала.

Структура проекта

muivScheduleTgBot/
│
├── cmd/
│    └── bot/            
│         └── main.go         # точка входа, запуск бота
│
├── internal/
│    ├── schedule/       
│    │     ├── schedule.go   # структура Day, методы ToMessage
│    │     ├── parser.go     # функции для парсинга Excel
│    │     └── updater.go      # обновление расписания
│    │
│    ├── bot/           
│    │     ├── handlers.go   # обработка сообщений и команд  
│    │     └── keyboard.go   # клавиатуры для бота
│
├── currentSchedule/         # папка с актуальными Excel-файлами
│     └── newSchedule.xlsx
│
└── go.mod                   # зависимости проекта
