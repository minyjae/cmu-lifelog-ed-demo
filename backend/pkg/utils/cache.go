package utils

import (
	"context"
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/redis/rueidis"
)

// Cache ห่อ rueidis.Client เพื่อใช้แคชข้อมูลแบบ JSON (cache-aside)
// ปลอดภัยเมื่อ client เป็น nil (เช่น redis ปิดอยู่) — ทุก method จะกลายเป็น no-op
// และเมื่อ redis มีปัญหา จะ log แล้วปล่อยผ่าน ไม่ทำให้ flow หลักล้ม
type Cache struct {
	client rueidis.Client
}

const (
	cacheTTL = 1 * time.Minute
)

// NewCache สร้าง Cache จาก rueidis.Client (ส่ง nil ได้ถ้าไม่ใช้ redis)
func NewCache(client rueidis.Client) *Cache {
	return &Cache{client: client}
}

// GetJSON อ่านค่าจาก cache ตาม key แล้ว unmarshal ลง dst
// คืน true เฉพาะเมื่อ hit และ unmarshal สำเร็จ
func (c *Cache) GetJSON(ctx context.Context, key string, dst any) bool {
	if c == nil || c.client == nil {
		return false
	}
	s, err := c.client.Do(ctx, c.client.B().Get().Key(key).Build()).ToString()
	if err != nil {
		return false // cache miss หรือ redis มีปัญหา → ให้ caller ไปอ่าน source
	}
	return json.Unmarshal([]byte(s), dst) == nil
}

// SetJSON marshal src แล้วเขียนลง cache (TTL = cacheTTL เสมอ, error แค่ log ไม่ return)
func (c *Cache) SetJSON(ctx context.Context, key string, src any) {
	if c == nil || c.client == nil {
		return
	}
	b, err := json.Marshal(src)
	if err != nil {
		return
	}
	if err := c.client.Do(ctx, c.client.B().Set().Key(key).Value(string(b)).Ex(cacheTTL).Build()).Error(); err != nil {
		log.Printf("warning: failed to set cache (%s): %v", key, err)
	}
}

// GetOrLoad: cache-aside แบบ generic — ลองอ่านจาก cache ก่อน ถ้า hit คืนเลย
// ถ้า miss จะเรียก fetch (อ่านจาก source จริง เช่น DB) แล้วเขียนผลกลับ cache (TTL = cacheTTL)
// ส่วน post-processing เฉพาะ domain (เช่น decorate field) ให้ caller ทำกับค่าที่คืนเอง
// เป็น free function เพราะ method ใน Go มี type parameter ไม่ได้
func GetOrLoad[T any](ctx context.Context, c *Cache, key string, fetch func() (T, error)) (T, error) {
	var v T
	if c.GetJSON(ctx, key, &v) {
		return v, nil
	}

	v, err := fetch()
	if err != nil {
		return v, err
	}

	c.SetJSON(ctx, key, v)
	return v, nil
}

// InvalidatePrefix ลบทุก key ที่ match prefix* ด้วย SCAN
// ใช้เมื่อข้อมูลต้นทางเปลี่ยนและกระทบหลาย cache key ที่ขึ้นต้นด้วย prefix เดียวกัน
func (c *Cache) InvalidatePrefix(ctx context.Context, prefix string) {
	if c == nil || c.client == nil {
		return
	}
	var cursor uint64
	for {
		entry, err := c.client.Do(ctx, c.client.B().Scan().Cursor(cursor).Match(prefix+"*").Count(100).Build()).AsScanEntry()
		if err != nil {
			log.Printf("warning: failed to scan cache (%s): %v", prefix, err)
			return
		}
		if len(entry.Elements) > 0 {
			if err := c.client.Do(ctx, c.client.B().Del().Key(entry.Elements...).Build()).Error(); err != nil {
				log.Printf("warning: failed to delete cache (%s): %v", prefix, err)
			}
		}
		cursor = entry.Cursor
		if cursor == 0 {
			break
		}
	}
}

// CacheKeyForIDs สร้าง key จาก prefix + รายการ id โดยเรียง id ก่อน
// เพื่อให้ชุด id เดียวกันได้ key เดียวกันเสมอ ไม่ว่าจะส่งมาลำดับไหน
func CacheKeyForIDs(prefix string, ids []uint) string {
	if len(ids) == 0 {
		return prefix + ":none"
	}
	sorted := append([]uint(nil), ids...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	parts := make([]string, len(sorted))
	for i, id := range sorted {
		parts[i] = strconv.FormatUint(uint64(id), 10)
	}
	return prefix + ":" + strings.Join(parts, "-")
}
