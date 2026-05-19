package utils

import (
	"fmt"
	"strings"
	"time"
)

// DTO สำหรับรับจาก FE
type CreateListQueueReq struct {
	Title             string   `json:"title"`
	StaffID           uint     `json:"staff_id"`
	FacultyID         uint     `json:"faculty_id"`
	StaffStatusID     uint     `json:"staff_status_id"`
	CourseStatusID    uint     `json:"course_status_id"`
	WordfileSubmit    string   `json:"wordfile_submit"`
	InfoSubmit        string   `json:"info_submit"`
	InfoSubmit14Days  string   `json:"info_submit_14days"`
	TimeRegister      string   `json:"time_register"`
	OnWeb             string   `json:"on_web"`
	AppointmentDateAW string   `json:"appointment_date_aw"`
	Owner             []string `json:"owner"`
	Note              string   `json:"note"`
}

var layouts = []string{
	time.RFC3339,          // "2025-05-21T00:00:00Z" or "+07:00"
	"2006-01-02 15:04:05", // "2025-05-21 00:00:00"
	"2006-01-02",          // "2025-05-21"
}

func ParseTimeFlexible(s string, loc *time.Location) (time.Time, error) {
	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, s)
		if err == nil {
			// ถ้า string ไม่มี timezone → ถือว่าเป็นเวลาใน loc
			if !strings.ContainsAny(s, "Z+-") {
				t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
			}
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("invalid time format: %q", s)
}
