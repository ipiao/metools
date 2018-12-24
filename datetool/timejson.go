package datetool

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	timeFormat  = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {

	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
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

type UnixTime time.Time

func (j UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(time.Time(j).Unix())), nil
}

func (j *UnixTime) UnmarshalJSON(data []byte) error {
	var err error
	if len(data) == 10 {
		var unix int64
		unix, err = strconv.ParseInt(string(data), 10, 64)
		t := time.Unix(unix, 0)
		*j = UnixTime(t)
	} else {
		t := time.Time{}
		err = json.Unmarshal(data, &t)
		*j = UnixTime(t)
	}
	return err
}
