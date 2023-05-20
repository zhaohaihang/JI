package datetime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(TIME_FORMAT, timeStr)
	*t = DateTime(t1)
	return err
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(TIME_FORMAT))
	return []byte(formatted), nil
}

func (t DateTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format(TIME_FORMAT), nil
}

func (t *DateTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = DateTime(vt)
	default:
		return errors.New("time type process error")
	}
	return nil
}

func (t *DateTime) String() string {
	return time.Time(*t).String()
}
