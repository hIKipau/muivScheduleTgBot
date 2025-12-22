package schedule

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"log/slog"
	"os"
)

// ParseSchedule - Parse new schedule into s *Schedule
func ParseSchedule(wb *xlsx.File, s *Schedule, course string, group int) error {
	sh, ok := wb.Sheet[course]
	if !ok {
		slog.Error("Cant find Sheet in Table:")
		os.Exit(1)
	}

	slog.Warn("Start parseDate for", "course", course, "group", group)
	parseDate(&s.WeekDay, sh)
	slog.Warn("Start parseTime", "course", course, "group", group)
	parseTime(&s.WeekDay, sh)
	slog.Warn("Start parseLesson", "course", course, "group", group)
	parseLesson(&s.WeekDay, sh, group)
	slog.Warn("Parsing successfully completed", "course", course, "group", group)

	return nil
}

func parseDate(days *[6]Day, sh *xlsx.Sheet) {
	slog.Info("parseDate init")
	day := 0
	for i := 5; i < 33; i += 5 {
		slog.Info("parseDate", day)

		theCell, err := sh.Cell(i, 0)
		if err != nil {
			fmt.Printf("error with parseDate: %v", err)
			panic(err)
		}
		days[day].Date = theCell.String()
		day++
	}
}

func parseTime(line *[6]Day, sh *xlsx.Sheet) {
	slog.Info("parseTime init")

	day := 0
	time := 0
	for i := 5; i < 35; i++ {
		slog.Info("parseTime", day)

		if time > 4 {
			day++
			time = 0
		}
		theCell, err := sh.Cell(i, 1)
		if err != nil {
			fmt.Printf("error with parseTime: %v", err)
			panic(err)
		}
		line[day].Time[time] = theCell.String()
		time++
	}
}

func parseLesson(line *[6]Day, sh *xlsx.Sheet, group int) {
	slog.Info("parseLesson init")
	day := 0
	lesson := 0
	for i := 5; i < 35; i++ {
		slog.Info("parseLesson", "day", day, "group", group)
		if lesson > 4 {
			day++
			lesson = 0
		}
		theCell, err := sh.Cell(i, group)
		if err != nil {
			fmt.Printf("error with parseTime: %v", err)
			panic(err)
		}
		line[day].Lesson[lesson] = theCell.String()
		lesson++
	}
}
