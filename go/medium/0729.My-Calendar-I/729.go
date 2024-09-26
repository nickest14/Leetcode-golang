// https://leetcode.com/problems/my-calendar-i/
// Level: Medium

package leetcode

type MyCalendar struct {
	bookings [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		bookings: make([][2]int, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, booking := range this.bookings {
		s, e := booking[0], booking[1]
		if start < e && end > s {
			return false
		}
	}
	this.bookings = append(this.bookings, [2]int{start, end})
	return true

}
