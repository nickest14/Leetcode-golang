// https://leetcode.com/problems/my-calendar-ii/
// Level: Medium

package leetcode

type MyCalendarTwo struct {
	calendar []Interval
	overlaps []Interval
}
type Interval struct {
	start, end int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		calendar: []Interval{},
		overlaps: []Interval{},
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, date := range this.overlaps {
		if start < date.end && end > date.start {
			return false
		}
	}

	for _, date := range this.calendar {
		if start < date.end && end > date.start {
			this.overlaps = append(this.overlaps, Interval{
				start: max(date.start, start),
				end:   min(date.end, end),
			})
		}
	}
	this.calendar = append(this.calendar, Interval{start, end})
	return true
}
