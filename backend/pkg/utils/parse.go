package utils

import (
	"strconv"
	"strings"
)

// ParseUintCSV แปลง query string แบบ "2,5,7" เป็น []uint
// ข้ามค่าที่ว่างหรือไม่ใช่ตัวเลข; ถ้า input ว่างจะคืน nil
func ParseUintCSV(s string) []uint {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	ids := make([]uint, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.ParseUint(p, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, uint(n))
	}
	return ids
}
