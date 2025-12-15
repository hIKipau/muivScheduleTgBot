package schedule

import (
	"fmt"
	"strings"
)

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

func LoadSchedule(scheduleCourse4 *[CountOfGroups4]*Schedule,
	scheduleCourse3 *[CountOfGroups3]*Schedule,
	scheduleCourse2 *[CountOfGroups2]*Schedule,
	scheduleCourse111 *[CountOfGroups111]*Schedule,
	scheduleCourse19 *[CountOfGroups19]*Schedule) error {

	err := ParseSchedule(scheduleCourse4[0], Course4.Course, Course4.Groups["ИСП/П-22-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse4[1], Course4.Course, Course4.Groups["ИСП/П-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse4[2], Course4.Course, Course4.Groups["ИСП/Р-22-09\nИСП/Р-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse4[3], Course4.Course, Course4.Groups["РЕК-22-09\nРЕК-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse4[4], Course4.Course, Course4.Groups["ДО-23-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(scheduleCourse3[0], Course3.Course, Course3.Groups["БАД-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[1], Course3.Course, Course3.Groups["ТД-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[2], Course3.Course, Course3.Groups["ФИН-23-09\nФИН-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[3], Course3.Course, Course3.Groups["КОМ-23-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[4], Course3.Course, Course3.Groups["ОДЛ-23-09\nОДЛ-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[5], Course3.Course, Course3.Groups["ЮР/со-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[6], Course3.Course, Course3.Groups["ПСО-23-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[7], Course3.Course, Course3.Groups["ПСО-23-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[8], Course3.Course, Course3.Groups["РЕК-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[9], Course3.Course, Course3.Groups["РЕК-23-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[10], Course3.Course, Course3.Groups["ДИЗ-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[11], Course3.Course, Course3.Groups["ИСП/р-23-09\nИСП/р-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[12], Course3.Course, Course3.Groups["ИСП/п-23-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[13], Course3.Course, Course3.Groups["ИСП/п-23-09.2\nИСП/п-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse3[14], Course3.Course, Course3.Groups["ФК-24-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(scheduleCourse2[0], Course2.Course, Course2.Groups["БАД-24-09\nБАД-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[1], Course2.Course, Course2.Groups["ФИН-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[2], Course2.Course, Course2.Groups["БУХ-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[3], Course2.Course, Course2.Groups["КСК-24-09\nКСК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[4], Course2.Course, Course2.Groups["ИСП/П-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[5], Course2.Course, Course2.Groups["ИСП/Р-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[6], Course2.Course, Course2.Groups["ТД-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[7], Course2.Course, Course2.Groups["ОДЛ-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[8], Course2.Course, Course2.Groups["ЮР/со-24-09\nЮР/со-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[9], Course2.Course, Course2.Groups["ЮР/са-24-09\nЮР/са-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse2[10], Course2.Course, Course2.Groups["РЕК-24-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(scheduleCourse111[0], Course111.Course, Course111.Groups["ФИН-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[1], Course111.Course, Course111.Groups["БУХ-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[2], Course111.Course, Course111.Groups["ТД-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[3], Course111.Course, Course111.Groups["ОДЛ-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[4], Course111.Course, Course111.Groups["ОИБАС-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[5], Course111.Course, Course111.Groups["ИСП/П-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[6], Course111.Course, Course111.Groups["ИСП/Р-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[7], Course111.Course, Course111.Groups["РЕК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[8], Course111.Course, Course111.Groups["ДИЗ-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[9], Course111.Course, Course111.Groups["ФК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse111[10], Course111.Course, Course111.Groups["ПВНК-25-11"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ParseSchedule(scheduleCourse19[0], Course19.Course, Course19.Groups["БАД-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[1], Course19.Course, Course19.Groups["БУХ-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[2], Course19.Course, Course19.Groups["ИСП/П-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[3], Course19.Course, Course19.Groups["ИСП/П-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[4], Course19.Course, Course19.Groups["ИСП/Р-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[5], Course19.Course, Course19.Groups["КСК-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[6], Course19.Course, Course19.Groups["ТД-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[7], Course19.Course, Course19.Groups["ТД-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[8], Course19.Course, Course19.Groups["ОДЛ-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[9], Course19.Course, Course19.Groups["ОДЛ-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[10], Course19.Course, Course19.Groups["ЮР/са-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[11], Course19.Course, Course19.Groups["ЮР/са-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[12], Course19.Course, Course19.Groups["ЮР/со-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[13], Course19.Course, Course19.Groups["РЕК-25-09.1"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[14], Course19.Course, Course19.Groups["РЕК-25-09.2"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[15], Course19.Course, Course19.Groups["РЕК-25-09.3"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[16], Course19.Course, Course19.Groups["ПВНК-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ParseSchedule(scheduleCourse19[17], Course19.Course, Course19.Groups["ФИН-25-09"])
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
