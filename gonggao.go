package myvin

import (
	"strings"

	"github.com/yudeguang17/stringsx"
)

// 获取公告型号,去除汉字等字符
func FmtGonggaoNo(gonggaoNo string) string {
	lastHAN := 0
	r := []rune(gonggaoNo)
	for i := range r {
		if stringsx.RuneIsHan(r[i]) || r[i] == 65533 {
			lastHAN = i
		}
	}
	if lastHAN != 0 {
		lastHAN++
	}
	return string(r[lastHAN:])
}

// 获取公告型号,去除汉字等字符
func FmtGonggaoNoPlus(gonggaoNo string) string {
	a := strings.ToUpper(gonggaoNo)
	a = strings.Replace(a, ` `, ``, -1)
	a = strings.Replace(a, `~`, ``, -1)
	a = strings.Replace(a, `!`, ``, -1)
	a = strings.Replace(a, `@`, ``, -1)
	a = strings.Replace(a, `#`, ``, -1)
	a = strings.Replace(a, `$`, ``, -1)
	a = strings.Replace(a, `%`, ``, -1)
	a = strings.Replace(a, `^`, ``, -1)
	a = strings.Replace(a, `&`, ``, -1)
	a = strings.Replace(a, `*`, ``, -1)
	a = strings.Replace(a, `(`, ``, -1)
	a = strings.Replace(a, `)`, ``, -1)
	a = strings.Replace(a, `（`, ``, -1)
	a = strings.Replace(a, `）`, ``, -1)
	a = strings.Replace(a, `-`, ``, -1)
	a = strings.Replace(a, `_`, ``, -1)
	a = strings.Replace(a, `=`, ``, -1)
	a = strings.Replace(a, `+`, ``, -1)
	a = strings.Replace(a, `{`, ``, -1)
	a = strings.Replace(a, `}`, ``, -1)
	a = strings.Replace(a, `[`, ``, -1)
	a = strings.Replace(a, `]`, ``, -1)
	a = strings.Replace(a, `\`, ``, -1)
	a = strings.Replace(a, `|`, ``, -1)
	a = strings.Replace(a, `/`, ``, -1)
	a = strings.Replace(a, `;`, ``, -1)
	a = strings.Replace(a, `:`, ``, -1)
	a = strings.Replace(a, `"`, ``, -1)
	a = strings.Replace(a, `“`, ``, -1)
	a = strings.Replace(a, `?`, ``, -1)
	a = strings.Replace(a, `<`, ``, -1)
	a = strings.Replace(a, `>`, ``, -1)
	// 对于罗马字母，还需要小心替换一下,很容易用户端口输入不一致，统一替换做到一致
	a = strings.Replace(a, `Ⅰ`, `I`, -1)
	a = strings.Replace(a, `Ⅱ`, `II`, -1)
	a = strings.Replace(a, `Ⅲ`, `III`, -1)
	a = strings.Replace(a, `Ⅳ`, `IV`, -1)
	a = strings.Replace(a, `Ⅴ`, `V`, -1)
	a = strings.Replace(a, `Ⅵ`, `VI`, -1)
	a = strings.Replace(a, `Ⅶ`, `VII`, -1)
	a = strings.Replace(a, `Ⅷ`, `VIII`, -1)
	a = strings.Replace(a, `Ⅸ`, `IX`, -1)
	//逗号不能替，因为多个公告号是用逗号分隔的
	//a = strings.Replace(a, `,`, ``, -1)
	a = strings.Replace(a, `.`, ``, -1)
	a = strings.Replace(a, `。`, ``, -1)
	// a = strings.Replace(a, `，`, ``, -1)
	return a
}
