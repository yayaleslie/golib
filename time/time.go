package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const (
	DateLayout       = "20060102"
	DateTimeLayout   = "2006-01-02 15:04:05"
	DateTimeDate     = "2006-01-02"
	DateTimeDateSpot = "2006.01.02"
	DateTimeNothing  = "20060102150405"

	Day   = time.Hour * 24
	Week  = Day * 7
	Month = Day * 30
	Year  = Day * 365

	DaySeconds  = 86400
	HourSeconds = 3600
	MinSeconds  = 60
)

var (
	DayStrFormatDef = DayStrFormatEn
	DayStrFormatEn  = []string{"day", "h", "m", "s"}
	DayStrFormatZh  = []string{"天", "时", "分", "秒"}
)

// Time JSONTime format json time field by myself
type Time struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// UnixMilli 毫秒
func (t Time) UnixMilli() int64 {
	return t.Time.UnixNano() / 1e6
}

// UnixMicro 微秒
func (t Time) UnixMicro() int64 {
	return t.Time.UnixNano() / 1e3
}

func (t Time) Add(d time.Duration) Time {
	t.Time = t.Time.Add(d)
	return t
}

func (t Time) DayStart() Time {
	locDate := t.Time.Format(DateLayout)
	locDayStart, _ := Parse(locDate, DateLayout)
	return locDayStart
}

func (t Time) DayEnd() Time {
	locDate := t.Time.Format(DateLayout)
	locDayStart, _ := Parse(locDate, DateLayout)
	return locDayStart.Add(Day - time.Duration(1))
}

func (t Time) Weekday(isEn ...bool) int {
	w := int(t.Time.Weekday())
	if w == 0 && (len(isEn) == 0 || !isEn[0]) {
		w = 7
	}

	return w
}

func (t Time) WeekStart() Time {
	w := int(t.Weekday())
	if w == 0 {
		w = 7
	}

	locDate := t.Time.Format(DateLayout)
	weekStart, _ := Parse(locDate, DateLayout)
	return weekStart.Add(-time.Duration(w-1) * Day)
}

func (t Time) WeekEnd() Time {
	w := int(t.Weekday())
	if w == 0 {
		w = 7
	}

	locDate := t.Time.Format(DateLayout)
	weekEnd, _ := Parse(locDate, DateLayout)
	return weekEnd.Add(time.Duration(7-w)*Day + Day - time.Duration(1))
}

func (t Time) String(formats ...string) string {
	var format = DateTimeLayout
	if len(formats) > 0 {
		format = formats[0]
	}

	return t.Format(format)
}

func Now() Time {
	return Time{time.Now()}
}

func Parse(v string, layout string, location ...*time.Location) (Time, error) {
	loc := time.Local
	if len(location) > 0 && location[0] != nil {
		loc = location[0]
	}
	t, err := time.ParseInLocation(layout, v, loc)
	return Time{t}, err
}

// sec 秒
func ParseSec(sec int64) Time {
	t := time.Unix(sec, 0)
	return Time{t}
}

// milli 毫秒
func ParseMilli(milli int64) Time {
	t := time.Unix(milli/1000, (milli%1000)*1e6)
	return Time{t}
}

func SecToStr(v int64, format ...[]string) (string, []int64) {
	fm := DayStrFormatDef
	if len(format) > 0 && len(format[0]) > 0 {
		fm = format[0]
	}

	var dStr, hStr, mStr, sStr string
	if len(fm) > 0 {
		dStr = fm[0]
	}
	if len(fm) > 1 {
		hStr = fm[1]
	}
	if len(fm) > 2 {
		mStr = fm[2]
	}
	if len(fm) > 3 {
		sStr = fm[0]
	}

	var str string
	d := v / DaySeconds
	h := (v - d*DaySeconds) / HourSeconds
	m := (v - d*DaySeconds - h*HourSeconds) / MinSeconds
	s := v - d*DaySeconds - h*HourSeconds - m*MinSeconds

	timeSlice := []int64{d, h, m, s}

	if d > 0 && dStr != "" {
		str += strconv.Itoa(int(d)) + dStr
	}
	if h > 0 && hStr != "" {
		str += strconv.Itoa(int(h)) + hStr
	}
	if m > 0 && mStr != "" {
		str += strconv.Itoa(int(m)) + mStr

	}
	if s > 0 && sStr != "" {
		str += strconv.Itoa(int(s)) + sStr
	}

	return str, timeSlice[:len(fm)]
}
