package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseRFC3339ToUnix parses RFC3339 datetime string to unix timestamp (seconds)
func ParseRFC3339ToUnix(s string) int64 {
	t, _ := time.Parse(time.RFC3339, s)
	return t.Unix()
}

// FormatToIndonesiaTime formats RFC3339 time to "WIB/WITA/WIT"
func FormatToIndonesiaTime(rfc3339 string) (string, error) {
	t, err := time.Parse(time.RFC3339, rfc3339)
	if err != nil {
		return "", err
	}

	// Determine timezone label from offset
	_, offset := t.Zone()

	tz := "UTC"
	switch offset {
	case 7 * 3600:
		tz = "WIB"
	case 8 * 3600:
		tz = "WITA"
	case 9 * 3600:
		tz = "WIT"
	}

	formatted := fmt.Sprintf(
		"%s %s",
		t.Format("02 Jan 2006, 15:04"),
		tz,
	)

	return formatted, nil
}

// FormatIDR formats integer to Indonesian Rupiah format
func FormatIDR(amount int) string {
	s := strconv.Itoa(amount)

	n := len(s)
	if n <= 3 {
		return "Rp " + s
	}

	var parts []string
	for n > 3 {
		parts = append([]string{s[n-3:]}, parts...)
		s = s[:n-3]
		n = len(s)
	}
	if n > 0 {
		parts = append([]string{s}, parts...)
	}

	return "Rp " + strings.Join(parts, ".")
}

func FormatDuration(minutes int) string {
	h := minutes / 60
	m := minutes % 60

	if h > 0 && m > 0 {
		return fmt.Sprintf("%dh %dm", h, m)
	}
	if h > 0 {
		return fmt.Sprintf("%dh", h)
	}
	return fmt.Sprintf("%dm", m)
}
