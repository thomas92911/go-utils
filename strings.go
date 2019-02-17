package utils

import (
	"encoding/binary"
	"strings"
)

// Int64ToBytes func
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// BytesToInt64 func
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// Uint64ToBytes func
func Uint64ToBytes(i uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], i)
	return buf[:]
}

// BytesToUint64 func
func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// ElideString func
func ElideString(s string, maxLen int) string {
	s = strings.Replace(s, "\n", "", -1)

	if maxLen <= 0 {
		return s
	}
	if len(s) <= maxLen {
		return s
	}
	if maxLen < 4 {
		return s[0:maxLen]
	}
	s = s[0:maxLen-3] + "..."

	return s
}

// Add function to add two uint64 numbers
func Add(existing, new []byte) []byte {
	return Uint64ToBytes(BytesToUint64(existing) + BytesToUint64(new))
}

// GetLastNameString func
func GetLastNameString(str string, sep string) (name string) {
	idx := strings.LastIndex(str, sep)
	if idx < 0 {
		return ""
	}
	name = str[idx+1:]
	return name
}

// StringListContain func
func StringListContain(sl []string, str string) bool {
	rc := false
	for _, s := range sl {
		if strings.Contains(s, str) {
			rc = true
			break
		}
	}
	return rc
}
