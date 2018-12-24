package datetool

import "time"

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
func DayInterval(t1, t2 time.Time) int64 {
	u1 := t1.Unix()
	u2 := t2.Unix()
	return u2/86400 - u1/86400
}
