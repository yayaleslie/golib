package time2

import (
	"time"
)

const (
	RangeTypeHour  = 1
	RangeTypeDay   = 2
	RangeTypeWeek  = 3
	RangeTypeMonth = 4
)

// RangeItem 每周
type RangeItem struct {
	Start Time
	End   Time
	Type  int
}

// separator 符号
// layout 格式
func (it RangeItem) String(separator string, layout ...string) string {
	format := DateTimeDate
	if len(layout) > 0 {
		format = layout[0]
	}

	start := it.Start.String(format)
	end := it.End.String(format)
	if it.Type == RangeTypeHour || start == end {
		return start
	}

	return start + separator + end
}

func (it RangeItem) StartMilli() int64 {
	return it.Start.UnixMilli()
}

func (it RangeItem) EndMilli() int64 {
	return it.End.UnixMilli()
}

func DateToRange(start, end string, rangeType int, layout string, location ...*time.Location) []RangeItem {
	switch rangeType {
	case RangeTypeHour:
		return DateToHour(start, end, layout, location...)
	case RangeTypeDay:
		return DateToDay(start, end, layout, location...)
	case RangeTypeWeek:
		return DateToWeek(start, end, layout, location...)
	case RangeTypeMonth:
		return DateToMouth(start, end, layout, location...)
	}
	return nil
}

// DateToHour start 开始时间
// end 结束时间
// layout 日期格式
// location 时区
func DateToHour(start, end, layout string, location ...*time.Location) []RangeItem {
	dayItems := DateToDay(start, end, layout, location...)
	items := make([]RangeItem, len(dayItems)*24)
	for k, v := range dayItems {
		for i := 0; i < 24; i++ {
			items[24*k+i] = RangeItem{
				Start: v.Start.Add(time.Duration(i) * time.Hour),
				End:   v.Start.Add(time.Duration(i+1)*time.Hour - 1),
				Type:  RangeTypeHour,
			}
		}
	}

	return items
}

func DateToDay(start, end, layout string, location ...*time.Location) []RangeItem {
	t1, _ := Parse(start, layout, location...)
	t2, _ := Parse(end, layout, location...)
	l := int(int64(t2.Sub(t1.Time)) / int64(Day))

	if l == 0 {
		return []RangeItem{{Start: t1.DayStart(), End: t2.DayEnd(), Type: RangeTypeDay}}
	}

	items := make([]RangeItem, 0)
	for step := 0; step <= l; step++ {
		items = append(items, RangeItem{
			Start: t1.Add(time.Duration(step) * Day).DayStart(),
			End:   t1.Add(time.Duration(step) * Day).DayEnd(),
			Type:  RangeTypeDay,
		})
	}

	return items
}

func DateToWeek(start, end, layout string, location ...*time.Location) []RangeItem {
	t1, _ := Parse(start, layout, location...)
	t2, _ := Parse(end, layout, location...)

	w1 := t1.Weekday()
	w2 := t2.Weekday()
	l := int(int64(t2.Sub(t1.Time)) / int64(Day))

	if w2-w1 == l {
		return []RangeItem{{Start: t1.DayStart(), End: t2.DayEnd(), Type: RangeTypeWeek}}
	}

	items := make([]RangeItem, 0)
	lStep, rStep := 0, 7-w1
	for rStep < l {
		items = append(items, RangeItem{
			Start: t1.Add(time.Duration(lStep) * Day).DayStart(),
			End:   t1.Add(time.Duration(rStep) * Day).DayEnd(),
			Type:  RangeTypeWeek,
		})

		lStep, rStep = rStep+1, rStep+7
		if rStep >= l {
			items = append(items, RangeItem{
				Start: t1.Add(time.Duration(lStep) * Day).DayStart(),
				End:   t1.Add(time.Duration(lStep+w2-1) * Day).DayEnd(),
				Type:  RangeTypeWeek,
			})
			break
		}
	}

	return items
}

func DateToMouth(start, end string, layout string, location ...*time.Location) []RangeItem {
	t1, _ := Parse(start, layout, location...)
	t2, _ := Parse(end, layout, location...)

	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	l := 12*(y2-y1) + int(m2-m1)
	if l == 0 {
		return []RangeItem{{Start: t1.DayStart(), End: t2.DayEnd(), Type: RangeTypeMonth}}
	}

	t3 := t1.MonthStart()
	items := make([]RangeItem, 0)
	mountStep, dayStep := 0, d1-1
	for mountStep < l {
		items = append(items, RangeItem{
			Start: t3.AddDate(0, mountStep, dayStep),
			End:   t3.AddDate(0, mountStep+1, -1).DayEnd(),
			Type:  RangeTypeMonth,
		})

		mountStep, dayStep = mountStep+1, 0
		if mountStep >= l {
			items = append(items, RangeItem{
				Start: t3.AddDate(0, mountStep, 0),
				End:   t3.AddDate(0, mountStep, d2-1).DayEnd(),
				Type:  RangeTypeMonth,
			})
			break
		}
	}

	return items
}
