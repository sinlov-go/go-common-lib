package date

import (
	"fmt"
	"time"
)

// DistanceBetweenTimestampSecondHuman
// convert DistanceBetweenTimestampSecond to human-readable
//
//	if from > to, return "N/A"
//	if from == to, return "0s"
//	if from < to, return "daysd hoursh minutesm secondss"
func DistanceBetweenTimestampSecondHuman(from, to int64) string {
	if from > to {
		return "N/A"
	}
	if from == to {
		return "0s"
	}
	days, hours, minutes, seconds := DistanceBetweenTimestampSecond(from, to)
	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}

// DistanceBetweenTimestampSecond
//
//	from and to must be unix timestamp
//	must be from < to
//	return days, hours, minutes, seconds
func DistanceBetweenTimestampSecond(from, to int64) (days, hours, minutes, seconds int) {
	if from > to {
		return 0, 0, 0, 0
	}
	fromT := time.Unix(from, 0)
	toT := time.Unix(to, 0)

	monthDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y1, m1, d1 := fromT.Date()
	y2, m2, d2 := toT.Date()
	h1, min1, s1 := fromT.Clock()
	h2, min2, s2 := toT.Clock()

	totalDays1 := y1*365 + d1
	for i := 0; i < (int)(m1)-1; i++ {
		totalDays1 += monthDays[i]
	}
	totalDays1 += leapYears(fromT)
	totalDays2 := y2*365 + d2
	for i := 0; i < (int)(m2)-1; i++ {
		totalDays2 += monthDays[i]
	}
	totalDays2 += leapYears(toT)
	days = totalDays2 - totalDays1
	hours = h2 - h1
	minutes = min2 - min1
	seconds = s2 - s1
	if seconds < 0 {
		seconds += 60
		minutes--
	}
	if minutes < 0 {
		minutes += 60
		hours--
	}
	if hours < 0 {
		hours += 24
		days--
	}
	return days, hours, minutes, seconds
}

func leapYears(date time.Time) (leaps int) {
	y, m, _ := date.Date()
	if m <= 2 {
		y--
	}
	leaps = y/4 + y/400 - y/100
	return leaps
}
