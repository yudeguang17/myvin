package myvin

import (
	"strings"
)

// GetWMI 解析 VIN 中的世界制造厂识别代码（WMI）。
// 返回：
//   - wmi: 标准3位WMI（始终返回）
//   - extension: 若为低产量制造商（WMI第三位为'9'），则返回第12-14位作为扩展标识；否则为空字符串。
func GetWMI(vin string) (wmi, extension string) {
	if len(vin) != 17 {
		return "", ""
	}

	vin = strings.ToUpper(vin)
	wmi = vin[:3]

	if vin[2] == '9' {
		extension = vin[11:14]
	}

	return wmi, extension
}
