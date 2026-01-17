package myvin

import (
	"github.com/yudeguang17/stringsx"
	"strings"
)

type country struct {
	From   string
	To     string
	NameEn string
	NameZh string
}

var countries = []country{
	//A–C = Africa
	{From: "AA", To: "AH", NameEn: "South Africa", NameZh: "南非"},
	{From: "AJ", To: "AK", NameEn: "Côte d'Ivoire", NameZh: "科特迪瓦"},
	{From: "AL", To: "AM", NameEn: "Lesotho", NameZh: "莱索托"},
	{From: "AN", To: "AP", NameEn: "Botswana", NameZh: "博茨瓦纳"},
	{From: "AR", To: "AS", NameEn: "Namibia", NameZh: "纳米比亚"},
	{From: "AT", To: "AU", NameEn: "Madagascar", NameZh: "马达加斯加"},
	{From: "AV", To: "AW", NameEn: "Mauritius", NameZh: "毛里求斯"},
	{From: "AX", To: "AY", NameEn: "Tunisia", NameZh: "突尼斯"},
	{From: "AZ", To: "A1", NameEn: "Cyprus", NameZh: "塞浦路斯"},
	{From: "A2", To: "A3", NameEn: "Zimbabwe", NameZh: "津巴布韦"},
	{From: "A4", To: "A5", NameEn: "Mozambique", NameZh: "莫桑比克"},
	{From: "BA", To: "BB", NameEn: "Angola", NameZh: "安哥拉"},
	{From: "BC", To: "BC", NameEn: "Ethiopia", NameZh: "埃塞俄比亚"},
	{From: "BF", To: "BG", NameEn: "Kenya", NameZh: "肯尼亚"},
	{From: "BH", To: "BH", NameEn: "Rwanda", NameZh: "卢旺达"},
	{From: "BL", To: "BL", NameEn: "Nigeria", NameZh: "尼日利亚"},
	{From: "BR", To: "BR", NameEn: "Algeria", NameZh: "阿尔及利亚"},
	{From: "BT", To: "BT", NameEn: "Swaziland", NameZh: "斯威士兰"},
	{From: "BU", To: "BU", NameEn: "Uganda", NameZh: "乌干达"},
	{From: "B3", To: "B4", NameEn: "Libya", NameZh: "利比亚"},
	{From: "CA", To: "CB", NameEn: "Egypt", NameZh: "埃及"},
	{From: "CF", To: "CG", NameEn: "Morocco", NameZh: "摩洛哥"},
	{From: "CL", To: "CM", NameEn: "Zambia", NameZh: "赞比亚"},
	//H–R = Asia
	{From: "HA", To: "H0", NameEn: "China", NameZh: "中国"},
	{From: "JA", To: "J0", NameEn: "Japan", NameZh: "日本"},
	{From: "KF", To: "KH", NameEn: "Israel", NameZh: "以色列"},
	{From: "KL", To: "KR", NameEn: "South Korea", NameZh: "韩国"},
	{From: "KS", To: "KT", NameEn: "Jordan", NameZh: "约旦"},
	{From: "K1", To: "K3", NameEn: "South Korea", NameZh: "韩国"},
	{From: "K5", To: "K5", NameEn: "Kyrgyzstan", NameZh: "吉尔吉斯斯坦"},
	{From: "LA", To: "L0", NameEn: "China", NameZh: "中国"},
	{From: "MA", To: "ME", NameEn: "India", NameZh: "印度"},
	{From: "MF", To: "MK", NameEn: "Indonesia", NameZh: "印度尼西亚"},
	{From: "ML", To: "MR", NameEn: "Thailand", NameZh: "泰国"},
	{From: "MS", To: "MS", NameEn: "Myanmar", NameZh: "缅甸"},
	{From: "MU", To: "MU", NameEn: "Mongolia", NameZh: "蒙古"},
	{From: "MX", To: "MX", NameEn: "Kazakhstan", NameZh: "哈萨克斯坦"},
	{From: "MY", To: "M0", NameEn: "India", NameZh: "印度"},
	{From: "NA", To: "NE", NameEn: "Iran", NameZh: "伊朗"},
	{From: "NF", To: "NG", NameEn: "Pakistan", NameZh: "巴基斯坦"},
	{From: "NJ", To: "NJ", NameEn: "Iraq", NameZh: "伊拉克"},
	{From: "NL", To: "NR", NameEn: "Turkey", NameZh: "土耳其"},
	{From: "NS", To: "NT", NameEn: "Uzbekistan", NameZh: "乌兹别克斯坦"},
	{From: "NV", To: "NV", NameEn: "Azerbaijan", NameZh: "阿塞拜疆"},
	{From: "NX", To: "NX", NameEn: "Tajikistan", NameZh: "塔吉克斯坦"},
	{From: "NY", To: "NY", NameEn: "Armenia", NameZh: "亚美尼亚"},
	{From: "N1", To: "N5", NameEn: "Iran", NameZh: "伊朗"},
	{From: "N7", To: "N8", NameEn: "Turkey", NameZh: "土耳其"},
	{From: "PA", To: "PC", NameEn: "Philippines", NameZh: "菲律宾"},
	{From: "PF", To: "PG", NameEn: "Singapore", NameZh: "新加坡"},
	{From: "PL", To: "PR", NameEn: "Malaysia", NameZh: "马来西亚"},
	{From: "PS", To: "PT", NameEn: "Bangladesh", NameZh: "孟加拉国"},
	{From: "PV", To: "PV", NameEn: "Cambodia", NameZh: "柬埔寨"},
	{From: "P5", To: "P0", NameEn: "India", NameZh: "印度"},
	{From: "RA", To: "RB", NameEn: "United Arab Emirates", NameZh: "阿联酋"},
	{From: "RF", To: "RK", NameEn: "Taiwan", NameZh: "中国台湾"},
	{From: "RL", To: "RN", NameEn: "Vietnam", NameZh: "越南"},
	{From: "RP", To: "RP", NameEn: "Laos", NameZh: "老挝"},
	{From: "RS", To: "RT", NameEn: "Saudi Arabia", NameZh: "沙特阿拉伯"},
	{From: "R1", To: "R7", NameEn: "Hong Kong", NameZh: "中国香港"},
	//E, S–Z = Europe
	{From: "EA", To: "E0", NameEn: "Russia", NameZh: "俄罗斯"},
	{From: "SA", To: "SM", NameEn: "United Kingdom", NameZh: "英国"},
	{From: "SN", To: "ST", NameEn: "Germany", NameZh: "德国"},
	{From: "SU", To: "SZ", NameEn: "Poland", NameZh: "波兰"},
	{From: "S1", To: "S2", NameEn: "Latvia", NameZh: "拉脱维亚"},
	{From: "S3", To: "S3", NameEn: "Georgia", NameZh: "格鲁吉亚"},
	{From: "S4", To: "S4", NameEn: "Iceland", NameZh: "冰岛"},
	{From: "TA", To: "TH", NameEn: "Switzerland", NameZh: "瑞士"},
	{From: "TJ", To: "TP", NameEn: "Czech Republic", NameZh: "捷克"},
	{From: "TR", To: "TV", NameEn: "Hungary", NameZh: "匈牙利"},
	{From: "TW", To: "T2", NameEn: "Portugal", NameZh: "葡萄牙"},
	{From: "T3", To: "T5", NameEn: "Serbia", NameZh: "塞尔维亚"},
	{From: "T6", To: "T6", NameEn: "Andorra", NameZh: "安道尔"},
	{From: "T7", To: "T8", NameEn: "Netherlands", NameZh: "荷兰"},
	{From: "UA", To: "UC", NameEn: "Spain", NameZh: "西班牙"},
	{From: "UH", To: "UM", NameEn: "Denmark", NameZh: "丹麦"},
	{From: "UN", To: "UR", NameEn: "Ireland", NameZh: "爱尔兰"},
	{From: "UU", To: "UX", NameEn: "Romania", NameZh: "罗马尼亚"},
	{From: "U1", To: "U2", NameEn: "North Macedonia", NameZh: "北马其顿"},
	{From: "U5", To: "U7", NameEn: "Slovakia", NameZh: "斯洛伐克"},
	{From: "U8", To: "U0", NameEn: "Bosnia and Herzegovina", NameZh: "波黑"},
	{From: "VA", To: "VE", NameEn: "Austria", NameZh: "奥地利"},
	{From: "VF", To: "VR", NameEn: "France", NameZh: "法国"},
	{From: "VS", To: "VW", NameEn: "Spain", NameZh: "西班牙"},
	{From: "VX", To: "V2", NameEn: "France", NameZh: "法国"},
	{From: "V3", To: "V5", NameEn: "Croatia", NameZh: "克罗地亚"},
	{From: "V6", To: "V8", NameEn: "Estonia", NameZh: "爱沙尼亚"},
	{From: "WA", To: "W0", NameEn: "Germany", NameZh: "德国"},
	{From: "XA", To: "XC", NameEn: "Bulgaria", NameZh: "保加利亚"},
	{From: "XD", To: "XE", NameEn: "Russia", NameZh: "俄罗斯"},
	{From: "XF", To: "XH", NameEn: "Greece", NameZh: "希腊"},
	{From: "XJ", To: "XK", NameEn: "Russia", NameZh: "俄罗斯"},
	{From: "XL", To: "XR", NameEn: "Netherlands", NameZh: "荷兰"},
	{From: "XS", To: "XW", NameEn: "Russia", NameZh: "俄罗斯"},
	{From: "XX", To: "XY", NameEn: "Luxembourg", NameZh: "卢森堡"},
	{From: "XZ", To: "X0", NameEn: "Russia", NameZh: "俄罗斯"},
	{From: "YA", To: "YE", NameEn: "Belgium", NameZh: "比利时"},
	{From: "YF", To: "YK", NameEn: "Finland", NameZh: "芬兰"},
	{From: "YN", To: "YN", NameEn: "Malta", NameZh: "马耳他"},
	{From: "YS", To: "YW", NameEn: "Sweden", NameZh: "瑞典"},
	{From: "YX", To: "Y2", NameEn: "Norway", NameZh: "挪威"},
	{From: "Y3", To: "Y5", NameEn: "Belarus", NameZh: "白俄罗斯"},
	{From: "Y6", To: "Y8", NameEn: "Ukraine", NameZh: "乌克兰"},
	{From: "ZA", To: "ZU", NameEn: "Italy", NameZh: "意大利"},
	{From: "ZX", To: "ZZ", NameEn: "Slovenia", NameZh: "斯洛文尼亚"},
	{From: "Z1", To: "Z1", NameEn: "San Marino", NameZh: "圣马力诺"},
	{From: "Z3", To: "Z5", NameEn: "Lithuania", NameZh: "立陶宛"},
	{From: "Z6", To: "Z0", NameEn: "Russia", NameZh: "俄罗斯"},
	//1–5, 7 = North America
	{From: "1A", To: "10", NameEn: "United States", NameZh: "美国"},
	{From: "2A", To: "20", NameEn: "Canada", NameZh: "加拿大"},
	{From: "3A", To: "3X", NameEn: "Mexico", NameZh: "墨西哥"},
	{From: "34", To: "34", NameEn: "Nicaragua", NameZh: "尼加拉瓜"},
	{From: "35", To: "35", NameEn: "Dominican Republic", NameZh: "多米尼加"},
	{From: "36", To: "36", NameEn: "Honduras", NameZh: "洪都拉斯"},
	{From: "37", To: "37", NameEn: "Panama", NameZh: "巴拿马"},
	{From: "38", To: "39", NameEn: "Puerto Rico", NameZh: "波多黎各"},
	{From: "4A", To: "40", NameEn: "United States", NameZh: "美国"},
	{From: "5A", To: "50", NameEn: "United States", NameZh: "美国"},
	{From: "7A", To: "70", NameEn: "United States", NameZh: "美国"},
	//6 = Oceania
	{From: "6A", To: "60", NameEn: "Australia", NameZh: "澳大利亚"},
	{From: "6Y", To: "61", NameEn: "New Zealand", NameZh: "新西兰"},
	//8–9 = South America
	{From: "8A", To: "8E", NameEn: "Argentina", NameZh: "阿根廷"},
	{From: "8F", To: "8G", NameEn: "Chile", NameZh: "智利"},
	{From: "8L", To: "8N", NameEn: "Ecuador", NameZh: "厄瓜多尔"},
	{From: "8S", To: "8W", NameEn: "Peru", NameZh: "秘鲁"},
	{From: "8X", To: "8Z", NameEn: "Venezuela", NameZh: "委内瑞拉"},
	{From: "82", To: "82", NameEn: "Bolivia", NameZh: "玻利维亚"},
	{From: "84", To: "84", NameEn: "Costa Rica", NameZh: "哥斯达黎加"},
	{From: "9A", To: "9E", NameEn: "Brazil", NameZh: "巴西"},
	{From: "9F", To: "9G", NameEn: "Colombia", NameZh: "哥伦比亚"},
	{From: "9S", To: "9V", NameEn: "Uruguay", NameZh: "乌拉圭"},
	{From: "91", To: "90", NameEn: "Brazil", NameZh: "巴西"},
}

// 获取生产地区英文名
func GetRegionEn(vin string) string {
	region := strings.ToUpper(stringsx.Left(vin, 1))
	if region >= "A" && region <= "H" {
		return "Africa"
	} else if region >= "J" && region <= "R" {
		return "Asia"
	} else if region >= "S" && region <= "Z" {
		return "Europe"
	} else if region >= "1" && region <= "5" {
		return "North America"
	} else if region >= "6" && region <= "7" {
		return "Oceania"
	} else if region >= "8" && region <= "9" {
		return "South America"
	} else {
		return "Unknown Region"
	}
}

// 获取生产地区中文名
func GetRegionZh(vin string) string {
	region := strings.ToUpper(stringsx.Left(vin, 1))
	if region >= "A" && region <= "H" {
		return "非洲"
	} else if region >= "J" && region <= "R" {
		return "亚洲"
	} else if region >= "S" && region <= "Z" {
		return "欧洲"
	} else if region >= "1" && region <= "5" || region == "7" {
		return "北美洲"
	} else if region == "6" {
		return "大洋洲"
	} else if region >= "8" && region <= "9" {
		return "南美洲"
	} else {
		return "未知区域"
	}
}

// 获取生产国家英文名
func GetCountryEn(vin string) string {
	qi := indexOf(strings.ToUpper(stringsx.Left(vin, 2)))
	for _, country := range countries {
		i := indexOf(country.From)
		j := indexOf(country.To)
		if qi >= i && qi <= j {
			return country.NameEn
		}
	}
	return "Unknown country"
}

// 获取生产国家中文名
func GetCountryZh(vin string) string {
	qi := indexOf(strings.ToUpper(stringsx.Left(vin, 2)))
	for _, country := range countries {
		i := indexOf(country.From)
		j := indexOf(country.To)
		if qi >= i && qi <= j {
			return country.NameZh
		}
	}
	return "未知国家"
}

const chars = "ABCDEFGHIJKLMNOPRSTUVWXYZ1234567890"

func indexOf(s string) int {
	return strings.IndexByte(chars, s[0])*len(chars) + strings.IndexByte(chars, s[1])
}
