package datetool

import "time"

func TimeMillSec(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

func NowMillSec() int64 {
	return time.Now().UnixNano() / 1e6
}

func IsDayToday(t time.Time) bool {
	now := time.Now()
	y, m, d := t.Date()
	ny, nm, nd := now.Date()
	return y == ny && m == nm && d == nd
}

func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// 后一天比前一天相差几天，同一天相差0
func DayInterval(t1, t2 time.Time) int {
	if t2.Before(t1) {
		return -DayInterval(t2, t1)
	}
	interval := 0
	y1 := t1.Year()
	y2 := t2.Year()
	if y1 < y2 {
		for i := y1; i < y2; i++ {
			interval += GetYearDays(i)
		}
	}
	interval += t2.YearDay() - t1.YearDay()
	return interval
}

func GetMonthDays(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(year) {
			return 29
		} else {
			return 28
		}
	}
	return 0
}

func GetYearDays(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
