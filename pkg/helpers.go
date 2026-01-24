package helpers

import (
	"encoding/base64"
	"strings"
)

// Base64ToBytes 将 Base64 编码的字符串转换为字节切片
func Base64ToBytes(b64 string) ([]byte, error) {
	if idx := strings.Index(b64, ","); idx != -1 {
		b64 = b64[idx+1:]
	}
	return base64.StdEncoding.DecodeString(b64)
}
