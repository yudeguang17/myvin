package myvin

import (
	"github.com/yudeguang17/stringsx"
	"strconv"
	"strings"
)

const Vin9strs = `0,1,2,3,4,5,6,7,8,9,x`

func IsVinLegal(vin string) bool {
	//判断长度是否合法
	if len(vin) != 17 {
		return false
	}
	vin = strings.ToLower(vin)
	//判断第9位是否合法
	vin9Int := vin9ToInt(vin[8:9])
	if vin9Int == -1 {
		return false
	}
	//每一位逐一转换为数字
	vin1 := vinToInt(vin[0:1])
	if vin1 == -1 {
		return false
	}

	vin2 := vinToInt(vin[1:2])
	if vin2 == -1 {
		return false
	}

	vin3 := vinToInt(vin[2:3])
	if vin3 == -1 {
		return false
	}

	vin4 := vinToInt(vin[3:4])
	if vin4 == -1 {
		return false
	}

	vin5 := vinToInt(vin[4:5])
	if vin5 == -1 {
		return false
	}

	vin6 := vinToInt(vin[5:6])
	if vin6 == -1 {
		return false
	}

	vin7 := vinToInt(vin[6:7])
	if vin7 == -1 {
		return false
	}

	vin8 := vinToInt(vin[7:8])
	if vin8 == -1 {
		return false
	}

	vin10 := vinToInt(vin[9:10])
	if vin10 == -1 {
		return false
	}

	vin11 := vinToInt(vin[10:11])
	if vin11 == -1 {
		return false
	}

	vin12 := vinToInt(vin[11:12])
	if vin12 == -1 {
		return false
	}

	vin13 := vinToInt(vin[12:13])
	if vin13 == -1 {
		return false
	}

	vin14 := vinToInt(vin[13:14])
	if vin14 == -1 {
		return false
	}

	vin15 := vinToInt(vin[14:15])
	if vin15 == -1 {
		return false
	}

	vin16 := vinToInt(vin[15:16])
	if vin16 == -1 {
		return false
	}

	vin17 := vinToInt(vin[16:17])
	if vin17 == -1 {
		return false
	}

	//全部相加
	total := vin1*8 + vin2*7 + vin3*6 + vin4*5 + vin5*4 + vin6*3 + vin7*2 + vin8*10 + vin10*9 + vin11*8 + vin12*7 + vin13*6 + vin14*5 + vin15*4 + vin16*3 + vin17*2
	return vin9Int == total%11
}

// VIN码第9位转换成数字
func vin9ToInt(vin9 string) int {
	switch vin9 {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "x":
		return 10
	default:
		return -1
	}
}

// 除第9位之外的其它位转换为数字
func vinToInt(vinStr string) int {
	switch vinStr {
	case "a":
		return 1
	case "b":
		return 2
	case "c":
		return 3
	case "d":
		return 4
	case "e":
		return 5
	case "f":
		return 6
	case "g":
		return 7
	case "h":
		return 8
	case "j":
		return 1
	case "k":
		return 2
	case "l":
		return 3
	case "m":
		return 4
	case "n":
		return 5
	case "p":
		return 7
	case "q":
		return 8
	case "r":
		return 9
	case "s":
		return 2
	case "t":
		return 3
	case "u":
		return 4
	case "v":
		return 5
	case "w":
		return 6
	case "x":
		return 7
	case "y":
		return 8
	case "z":
		return 9

	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "0":
		return 0
	default:
		return -1
	}
}

// 判断最后6位是不是数字,是数字就合法
func IsVinLast6Legal(vin string) bool {
	vinLast6 := stringsx.Right(vin, 6)
	if len(vinLast6) != 6 {
		return false
	}
	_, err := strconv.Atoi(vinLast6)
	if err != nil {
		return false
	}
	return true
}

// 判断最后5位是不是数字,是数字就合法
func IsVinLast5Legal(vin string) bool {
	vinLast5 := stringsx.Right(vin, 5)
	if len(vinLast5) != 5 {
		return false
	}
	_, err := strconv.Atoi(vinLast5)
	if err != nil {
		return false
	}
	return true
}

// 判断最后4位是不是数字,是数字就合法
func IsVinLast4Legal(vin string) bool {
	vinLast4 := stringsx.Right(vin, 4)
	if len(vinLast4) != 4 {
		return false
	}
	_, err := strconv.Atoi(vinLast4)
	if err != nil {
		return false
	}
	return true
}
