package stringUtils

import (
	"strconv"
	"strings"
)

// 将 escaped 的 uft8 转化为真正的 unicode
func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}