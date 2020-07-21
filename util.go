package optional

import (
	"strings"
	"unsafe"
)

// 用于校验uuid
var xvalues = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// 用于校验uuid
func xtob(x1, x2 byte) (byte, bool) {
	b1 := xvalues[x1]
	b2 := xvalues[x2]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}

// []byte2string
func bytesToStr(value []byte) string {
	return *(*string)(unsafe.Pointer(&value)) // nolint
}

// string2[]byte
func strToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s)) // nolint
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h)) // nolint
}

// 中国手机号码前缀
var chinaMobilePrefix = []uint8{
	// 移动
	139, 138, 137, 136, 135, 134, 147, 150, 151, 152, 157, 158, 159, 165, 178, 182, 183, 184, 187, 188, 198,
	// 联通
	130, 131, 132, 155, 156, 166, 167, 185, 186, 145, 175, 176,
	// 电信
	133, 153, 162, 177, 173, 180, 181, 189, 191, 199,
	// 虚拟运营商
	170, 171,
}

func ToCamelCase(s string) string {
	slice := strings.Split(s, "_")
	for i := range slice {
		if i > 0 {
			slice[i] = strings.Title(slice[i])
		}
	}
	return strings.Join(slice, "")
}
func ToPascalCase(s string) string {
	slice := strings.Split(s, "_")
	for i := range slice {
		slice[i] = strings.Title(slice[i])
	}
	return strings.Join(slice, "")
}
func ToSnakeCase(s string) string {
	strLen := len(s)
	result := make([]byte, 0, strLen*2)
	j := false
	for i := 0; i < strLen; i++ {
		char := s[i]
		if i > 0 && char >= 'A' && char <= 'Z' && j {
			result = append(result, '_')
		}
		if char != '_' {
			j = true
		}
		result = append(result, char)
	}
	return strings.ToLower(string(result))
}
