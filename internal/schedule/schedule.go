package schedule

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
	"log/slog"
	"strings"
)

// Variable store *Schedule for all groups
type GlobalSchedules [5][]*Schedule

// NewGlobalSchedules() - Create and initialize new pointer to GlobalSchedules
func NewGlobalSchedules() *GlobalSchedules {
	var gs GlobalSchedules
	gs.new()
	return &gs
}

func (gs *GlobalSchedules) new() {
	(*gs)[0] = make([]*Schedule, CountOfGroups19)
	for i := 0; i < CountOfGroups19; i++ {
		(*gs)[Course19i][i] = &Schedule{}
	}

	(*gs)[1] = make([]*Schedule, CountOfGroups111)
	for i := 0; i < CountOfGroups111; i++ {
		(*gs)[Course111i][i] = &Schedule{}
	}

	(*gs)[2] = make([]*Schedule, CountOfGroups2)
	for i := 0; i < CountOfGroups2; i++ {
		(*gs)[Course2i][i] = &Schedule{}
	}

	(*gs)[3] = make([]*Schedule, CountOfGroups3)
	for i := 0; i < CountOfGroups3; i++ {
		(*gs)[Course3i][i] = &Schedule{}
	}

	(*gs)[4] = make([]*Schedule, CountOfGroups4)
	for i := 0; i < CountOfGroups4; i++ {
		(*gs)[Course4i][i] = &Schedule{}
	}

}

type Schedule struct {
	WeekDay [6]Day
}
type Day struct {
	Date   string
	Time   [5]string
	Lesson [5]string
}

func (d Day) ToMessage() string {
	var b strings.Builder

	b.WriteString(d.Date + "\n")
	for i := 0; i < len(d.Time); i++ {
		if d.Time[i] == "" && d.Lesson[i] == "" {
			continue // пропуск пустых строк
		}
		b.WriteString(d.Time[i] + " — " + d.Lesson[i] + "\n")
		b.WriteString("------------------------------" + "\n")

	}

	return b.String()
}
func GetSchedule(course string, group string, gs *GlobalSchedules) *Schedule {

	var curCourse int
	curGroup := ScheduleInfo[course][group] - 3

	switch course {
	case Course4s:
		curCourse = 4
	case Course3s:
		curCourse = 3
	case Course2s:
		curCourse = 2
	case Course111s:
		curCourse = 1
	case Course19s:
		curCourse = 0
	default:
		curCourse = 4
	}

	return (*gs)[curCourse][curGroup]
}

// LoadSchedule loads a new schedule from the path:
// "./currentSchedule/newSchedule.xlsx"
func (gs *GlobalSchedules) LoadSchedule() error {
	wb, err := xlsx.OpenFile("./currentSchedule/newSchedule.xlsx")
	if err != nil {
		slog.Error("Cant open newSchedule.xlsx:", err)
		return err
	}

	err = ParseSchedule(wb, (*gs)[Course4i][0], Course4s, ScheduleInfo[Course4s]["ИСП/П-22-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course4i][1], Course4s, ScheduleInfo[Course4s]["ИСП/П-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course4i][2], Course4s, ScheduleInfo[Course4s]["ИСП/Р-22-09\nИСП/Р-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course4i][3], Course4s, ScheduleInfo[Course4s]["РЕК-22-09\nРЕК-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course4i][4], Course4s, ScheduleInfo[Course4s]["ДО-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(wb, (*gs)[Course3i][0], Course3s, ScheduleInfo[Course3s]["БАД-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][1], Course3s, ScheduleInfo[Course3s]["ТД-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][2], Course3s, ScheduleInfo[Course3s]["ФИН-23-09\nФИН-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][3], Course3s, ScheduleInfo[Course3s]["КОМ-23-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][4], Course3s, ScheduleInfo[Course3s]["ОДЛ-23-09\nОДЛ-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][5], Course3s, ScheduleInfo[Course3s]["ЮР/со-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][6], Course3s, ScheduleInfo[Course3s]["ПСО-23-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][7], Course3s, ScheduleInfo[Course3s]["ПСО-23-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][8], Course3s, ScheduleInfo[Course3s]["РЕК-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][9], Course3s, ScheduleInfo[Course3s]["РЕК-23-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][10], Course3s, ScheduleInfo[Course3s]["ДИЗ-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][11], Course3s, ScheduleInfo[Course3s]["ИСП/р-23-09\nИСП/р-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][12], Course3s, ScheduleInfo[Course3s]["ИСП/п-23-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][13], Course3s, ScheduleInfo[Course3s]["ИСП/п-23-09.2\nИСП/п-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course3i][14], Course3s, ScheduleInfo[Course3s]["ФК-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(wb, (*gs)[Course2i][0], Course2s, ScheduleInfo[Course2s]["БАД-24-09\nБАД-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][1], Course2s, ScheduleInfo[Course2s]["ФИН-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][2], Course2s, ScheduleInfo[Course2s]["БУХ-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][3], Course2s, ScheduleInfo[Course2s]["КСК-24-09\nКСК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][4], Course2s, ScheduleInfo[Course2s]["ИСП/П-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][5], Course2s, ScheduleInfo[Course2s]["ИСП/Р-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][6], Course2s, ScheduleInfo[Course2s]["ТД-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][7], Course2s, ScheduleInfo[Course2s]["ОДЛ-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][8], Course2s, ScheduleInfo[Course2s]["ЮР/со-24-09\nЮР/со-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][9], Course2s, ScheduleInfo[Course2s]["ЮР/са-24-09\nЮР/са-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course2i][10], Course2s, ScheduleInfo[Course2s]["РЕК-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(wb, (*gs)[Course111i][0], Course111s, ScheduleInfo[Course111s]["ФИН-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][1], Course111s, ScheduleInfo[Course111s]["БУХ-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][2], Course111s, ScheduleInfo[Course111s]["ТД-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][3], Course111s, ScheduleInfo[Course111s]["ОДЛ-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][4], Course111s, ScheduleInfo[Course111s]["ОИБАС-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][5], Course111s, ScheduleInfo[Course111s]["ИСП/П-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][6], Course111s, ScheduleInfo[Course111s]["ИСП/Р-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][7], Course111s, ScheduleInfo[Course111s]["РЕК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][8], Course111s, ScheduleInfo[Course111s]["ДИЗ-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][9], Course111s, ScheduleInfo[Course111s]["ФК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course111i][10], Course111s, ScheduleInfo[Course111s]["ПВНК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(wb, (*gs)[Course19i][0], Course19s, ScheduleInfo[Course19s]["БАД-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][1], Course19s, ScheduleInfo[Course19s]["БУХ-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][2], Course19s, ScheduleInfo[Course19s]["ИСП/П-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][3], Course19s, ScheduleInfo[Course19s]["ИСП/П-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][4], Course19s, ScheduleInfo[Course19s]["ИСП/Р-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][5], Course19s, ScheduleInfo[Course19s]["КСК-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][6], Course19s, ScheduleInfo[Course19s]["ТД-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][7], Course19s, ScheduleInfo[Course19s]["ТД-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][8], Course19s, ScheduleInfo[Course19s]["ОДЛ-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][9], Course19s, ScheduleInfo[Course19s]["ОДЛ-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][10], Course19s, ScheduleInfo[Course19s]["ЮР/са-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][11], Course19s, ScheduleInfo[Course19s]["ЮР/са-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][12], Course19s, ScheduleInfo[Course19s]["ЮР/со-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][13], Course19s, ScheduleInfo[Course19s]["РЕК-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][14], Course19s, ScheduleInfo[Course19s]["РЕК-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][15], Course19s, ScheduleInfo[Course19s]["РЕК-25-09.3"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][16], Course19s, ScheduleInfo[Course19s]["ПВНК-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(wb, (*gs)[Course19i][17], Course19s, ScheduleInfo[Course19s]["ФИН-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
