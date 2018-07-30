package datetool

import (
	"log"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02T15:04:05Z07:00"
	timeFormat  = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {

	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	log.Println("now:", now)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).Add(time.Hour*8).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}
