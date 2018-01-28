package datetool

import "time"

// OffsetDate 偏移日期
// 以当前开始时间至后一天开始时间前一秒为一个时间周期，该周期在形式上记作开始当前开始时间所在日期的周期
type OffsetDate struct {
	dayStartHour   int
	dayStartMinute int
	loc            *time.Location
	monthStartDay  int
}

// NewOffsetDate 创建一个偏移日期计算器
func NewOffsetDate(hour, minute int) *OffsetDate {
	return &OffsetDate{
		dayStartHour:   hour,
		dayStartMinute: minute,
		monthStartDay:  1,
		loc:            time.Local,
	}
}

// SetMonthStartDay 设置location
func (d *OffsetDate) SetMonthStartDay(day int) {
	d.monthStartDay = day
}

// SetLocation 设置location
func (d *OffsetDate) SetLocation(loc *time.Location) {
	d.loc = loc
}

// DayStartTime 返回传入时间的周期的开始时间
func (d *OffsetDate) DayStartTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), ct.Day(), d.dayStartHour, d.dayStartMinute, 0, 0, d.loc)
	if ct.Before(t1) {
		t1 = t1.AddDate(0, 0, -1)
	}
	return t1
}

// MidDayStartTime 返回传入时间的周期中点的开始时间
func (d *OffsetDate) MidDayStartTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), ct.Day(), d.dayStartHour, d.dayStartMinute, 0, 0, d.loc).Add(time.Hour * 12)
	if ct.Before(t1) {
		t1 = t1.AddDate(0, 0, -1)
	}
	return t1
}

// DayStartUnixTime 返回传入时间的周期的开始时间戳
func (d *OffsetDate) DayStartUnixTime(ct int64) int64 {
	offset := int64(d.dayStartHour) * 3600
	t1 := (ct-offset)/86400*86400 + offset
	return t1
}

// HalfDayStartUnixTime 返回传入时间的半周期的开始时间戳
func (d *OffsetDate) HalfDayStartUnixTime(ct int64) int64 {
	offset := int64(d.dayStartHour) * 3600
	t1 := (ct-offset)/43200*43200 + offset
	return t1
}

// CertainDayStartTime 返回传入时间年月日的周期开始时间
func (d *OffsetDate) CertainDayStartTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), ct.Day(), d.dayStartHour, d.dayStartMinute, 0, 0, d.loc)
	return t1
}

// DayEndTime 返回传入时间周期的结束时间
func (d *OffsetDate) DayEndTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), ct.Day(), d.dayStartHour, d.dayStartMinute, 0, 0, d.loc).Add(-1 * time.Second)
	if ct.After(t1) {
		t1 = t1.AddDate(0, 0, 1)
	}
	return t1
}

// MidDayEndTime 返回传入时间的周期中点的结束时间
func (d *OffsetDate) MidDayEndTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), ct.Day(), d.dayStartHour, d.dayStartMinute, 0, 0, d.loc).Add(time.Hour * -12)
	if ct.Before(t1) {
		t1 = t1.AddDate(0, 1, 0)
	}
	return t1
}

// DayEndUnixTime 返回传入时间的周期的结束时间戳
func (d *OffsetDate) DayEndUnixTime(ct int64) int64 {
	offset := int64(d.dayStartHour) * 3600
	t1 := (ct-offset)/86400*86400 + offset
	return t1 + 86399
}

// HalfDayEndUnixTime 返回传入时间的半周期的开始时间戳
func (d *OffsetDate) HalfDayEndUnixTime(ct int64) int64 {
	offset := int64(d.dayStartHour) * 3600
	t1 := (ct-offset)/43200*43200 + offset
	return t1 + 43199
}

// CertainDayEndTime 返回传入时间年月日的周期结束时间
func (d *OffsetDate) CertainDayEndTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), ct.Day(), d.dayStartHour, d.dayStartMinute, 0, 0, d.loc).Add(-1*time.Second).AddDate(0, 1, 0)
	return t1
}

// MonthStartTime 返回传入时间周期当月的开始时间
func (d *OffsetDate) MonthStartTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), d.monthStartDay, d.dayStartHour, d.dayStartMinute, 0, 0, d.loc)
	if ct.Before(t1) {
		t1 = t1.AddDate(0, -1, 0)
	}
	return t1
}

// MonthEndTime 返回传入时间周期当月的结束时间
func (d *OffsetDate) MonthEndTime(ct time.Time) time.Time {
	t1 := time.Date(ct.Year(), ct.Month(), d.monthStartDay, d.dayStartHour, d.dayStartMinute, 0, 0, d.loc).Add(-1 * time.Second)
	if ct.After(t1) {
		t1 = t1.AddDate(0, 1, 0)
	}
	return t1
}

// PreSecondHM 返回前一秒
func (d *OffsetDate) PreSecondHM() (h, m, s int) {
	h = d.dayStartHour
	m = d.dayStartMinute
	s = 59
	if m > 0 {
		m--
		return
	}
	m = 59
	if h != 0 {
		h--
		return
	}
	h = 23
	return
}
