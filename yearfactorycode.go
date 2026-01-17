package myvin

import (
	"fmt"
	"strings"
)

// 通过年份算出年份代码
func GetYearCodeFromYear(year string) string {
	switch year {
	case "1971":
		return "1"
	case "1972":
		return "2"
	case "1973":
		return "3"
	case "1974":
		return "4"
	case "1975":
		return "5"
	case "1976":
		return "6"
	case "1977":
		return "7"
	case "1978":
		return "8"
	case "1979":
		return "9"
	case "1980":
		return "a"
	case "1981":
		return "b"
	case "1982":
		return "c"
	case "1983":
		return "d"
	case "1984":
		return "e"
	case "1985":
		return "f"
	case "1986":
		return "g"
	case "1987":
		return "h"
	case "1988":
		return "j"
	case "1989":
		return "k"
	case "1990":
		return "l"
	case "1991":
		return "m"
	case "1992":
		return "n"
	case "1993":
		return "p"
	case "1994":
		return "r"
	case "1995":
		return "s"
	case "1996":
		return "t"
	case "1997":
		return "v"
	case "1998":
		return "w"
	case "1999":
		return "x"
	case "2000":
		return "y"
	case "2001":
		return "1"
	case "2002":
		return "2"
	case "2003":
		return "3"
	case "2004":
		return "4"
	case "2005":
		return "5"
	case "2006":
		return "6"
	case "2007":
		return "7"
	case "2008":
		return "8"
	case "2009":
		return "9"
	case "2010":
		return "a"
	case "2011":
		return "b"
	case "2012":
		return "c"
	case "2013":
		return "d"
	case "2014":
		return "e"
	case "2015":
		return "f"
	case "2016":
		return "g"
	case "2017":
		return "h"
	case "2018":
		return "j"
	case "2019":
		return "k"
	case "2020":
		return "l"
	case "2021":
		return "m"
	case "2022":
		return "n"
	case "2023":
		return "P"
	case "2024":
		return "r"
	case "2025":
		return "s"
	default:
		//没匹配上直接返回0
		return "-1"
	}
}

// 从VIN中获得年份代码
func GetYearCodeFromVin(vin string) string {
	return vin[9:10]
}

// 从VIN码中获得工厂代码
func GetFactoryCodeFromVin(vin string) string {
	return vin[10:11]
}

func GetVin8YearCodeFactoryCodeVinlast6(vin string) (vin8, yearCode, factoryCode, vinLast6 string, err error) {
	if len(vin) != 17 {
		err = fmt.Errorf("vin必须是17位")
		return 
	}
	vin8 = vin[0:8]
	yearCode = vin[9:10]
	factoryCode = vin[10:11]
	vinLast6 = vin[11:]
	return
}

// 输入年份代码，当前仍然在路上运行的车辆年份
func GetYearFromVinCurrent30Years(vin string) int {
	return GetYearFromYearCodeCurrent30Years(vin[9:10])
}

// 输入年份代码，当前仍然在路上运行的车辆年份
// (汽车一般15年报废,只有最近30年左右的汽车年份有意义，VIN年份信息每30年轮回一次)
func GetYearFromYearCodeCurrent30Years(yearCode string) int {
	yearCode = strings.ToLower(yearCode)
	switch yearCode {
	case "t":
		return 1996
	case "v":
		return 1997
	case "w":
		return 1998
	case "x":
		return 1999
	case "y":
		return 2000
	case "1":
		return 2001
	case "2":
		return 2002
	case "3":
		return 2003
	case "4":
		return 2004
	case "5":
		return 2005
	case "6":
		return 2006
	case "7":
		return 2007
	case "8":
		return 2008
	case "9":
		return 2009
	case "a":
		return 2010
	case "b":
		return 2011
	case "c":
		return 2012
	case "d":
		return 2013
	case "e":
		return 2014
	case "f":
		return 2015
	case "g":
		return 2016
	case "h":
		return 2017
	case "j":
		return 2018
	case "k":
		return 2019
	case "l":
		return 2020
	case "m":
		return 2021
	case "n":
		return 2022
	case "p":
		return 2023
	case "r":
		return 2024
	case "s":
		return 2025
	}
	return -1
}

// 根据VIN码以及年份代码获得VIN年份信息
func GetYearFromVinPlus(vin string, madeYear int) int {
	return GetYearFromYearCodePlus(vin[9:10], madeYear)
}

// 根据当前生产日期以及年份代码获得VIN年份信息
func GetYearFromYearCodePlus(yearCode string, madeYear int) int {
	yearCode = strings.ToLower(yearCode)
	var years []int
	switch yearCode {
	case "1":
		years = []int{1971, 2001}
	case "2":
		years = []int{1972, 2002}
	case "3":
		years = []int{1973, 2003}
	case "4":
		years = []int{1974, 2004}
	case "5":
		years = []int{1975, 2005}
	case "6":
		years = []int{1976, 2006}
	case "7":
		years = []int{1977, 2007}
	case "8":
		years = []int{1978, 2008}
	case "9":
		years = []int{1979, 2009}
	case "a":
		years = []int{1980, 2010}
	case "b":
		years = []int{1981, 2011}
	case "c":
		years = []int{1982, 2012}
	case "d":
		years = []int{1983, 2013}
	case "e":
		years = []int{1984, 2014}
	case "f":
		years = []int{1985, 2015}
	case "g":
		years = []int{1986, 2016}
	case "h":
		years = []int{1987, 2017}
	case "j":
		years = []int{1988, 2018}
	case "k":
		years = []int{1989, 2019}
	case "l":
		years = []int{1990, 2020}
	case "m":
		years = []int{1991, 2021}
	case "n":
		years = []int{1992, 2022}
	case "p":
		years = []int{1993, 2023}
	case "r":
		years = []int{1994, 2024}
	case "s":
		years = []int{1995, 2025}
	case "t":
		years = []int{1996, 2026}
	case "v":
		years = []int{1997, 2027}
	case "w":
		years = []int{1998, 2028}
	case "x":
		years = []int{1999, 2029}
	case "y":
		years = []int{2000, 2030}
	default:
		return madeYear
	}
	return getBestMutchModelYear(years, madeYear)

}

// 根据若干年份，返回最接近的年份
func getBestMutchModelYear(years []int, madeYear int) int {
	if len(years) == 0 {
		return -1
	}
	modelYear := -1
	difference := 10000
	for _, year := range years {
		if x := abs(year - madeYear); x < difference {
			modelYear = year
			difference = x
		}
	}
	return modelYear
}

// 绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
