package base62

import (
	"math"
	"strings"
)

// 62进制转换的模块

// 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

// 0-9: 0-9
// a-z: 10-35
// A-Z: 36-61

// 10进制数   转换    62进制数
//    0                 0
//    1                 1
//    10                a
//    11                b
//    61                Z
//    62                10
//    63                11
//    6347              ?

// 如何实现62进制转换

// const base62Str = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
// const base62Str = `J0rs12O5TUV8IW7D9aBdXeCfghiMQj3klmop6qtuvbcwx4zAEFGHKLNnPRYSZy`

// 为了避免被人恶意请求，我们可以将上面的字符串打乱

var (
	base62Str string
)

// MustInit 要使用base62这包必须要调用该函数完成初始化
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string!")
	}
	base62Str = bs
}

// Int2String 十进制数转为62进制字符串
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(base62Str[0])
	}
	bl := []byte{} // 23 40 1
	for seq > 0 {
		mod := seq % 62
		div := seq / 62
		bl = append(bl, base62Str[mod])
		seq = div
	}
	// 最后把的到的数反转一下
	return string(reverse(bl))
}

// String2Int 62进制字符串转为10进制数
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)
	// 从右往左遍历
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		seq += uint64(strings.Index(base62Str, string(b))) * uint64(base)
	}
	return seq
}

// [1 2 3 4 5]
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
