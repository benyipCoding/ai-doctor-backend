package helpers

import (
	"encoding/base64"
	"strconv"
	"strings"
)

// Base64ToBytes 将 Base64 编码的字符串转换为字节切片
func Base64ToBytes(b64 string) ([]byte, error) {
	if idx := strings.Index(b64, ","); idx != -1 {
		b64 = b64[idx+1:]
	}
	return base64.StdEncoding.DecodeString(b64)
}

// ParseLimitOffset 从两个字符串解析 limit 与 offset，返回默认值当解析失败
func ParseLimitOffset(limitStr, offsetStr string, defaultLimit int) (int, int) {
	limit := defaultLimit
	offset := 0
	if limitStr != "" {
		if v, err := strconv.Atoi(limitStr); err == nil && v > 0 {
			limit = v
		}
	}
	if offsetStr != "" {
		if v, err := strconv.Atoi(offsetStr); err == nil && v >= 0 {
			offset = v
		}
	}
	return limit, offset
}
