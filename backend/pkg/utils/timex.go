package utils

import "time"

// NormalizeDate ทำเวลาให้เป็นเวลาเที่ยงคืนของโซนที่กำหนด (ตัดผลกระทบจาก TZ/hh:mm:ss)
func NormalizeDate(t time.Time, loc *time.Location) time.Time {
	if t.IsZero() {
		return t
	}
	t = t.In(loc)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
}

// EqualDate เทียบความเท่ากันแบบ "เท่าในระดับวัน" (ตัด TZ/hh:mm:ss)
func EqualDate(a, b time.Time, loc *time.Location) bool {
	na := NormalizeDate(a, loc)
	nb := NormalizeDate(b, loc)
	return na.Equal(nb)
}
