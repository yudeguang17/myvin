package myvin

import (
	"fmt"
	"github.com/yudeguang17/stringsx"
	"sort"
	"strconv"
	"strings"
)

// 根据VIN样本伪随机生成VIN码，规则为从尾数0开始每隔N个数字产生一个VIN码,every必须大于等于1小于999999
func GetVinsByRandFrom(vin string, every int) []string {
	vins := make([]string, 0, 999999/every)
	if every < 1 || every > 999999 {
		return vins
	}
	vin_8 := vin[0:8]
	year_code := GetYearCodeFromVin(vin)
	factory_code := GetFactoryCodeFromVin(vin)
	vin_last_6 := stringsx.Right(vin, 6)

	x, _, cur_serial_number_begin, cur_serial_number_end, err := split_vin_last_6(vin_last_6)

	if err != nil {
		return vins
	}
	for i := cur_serial_number_begin; i <= cur_serial_number_end; i = i + every {
		vin_last_6 := get_vin_last_6_from(x, i)
		vin = getOneVinFrom(vin_8, year_code, factory_code, vin_last_6)
		vins = append(vins, strings.ToUpper(vin))
	}
	return vins
}

// 根据样本VIN码生成与之最为接近的N个VIN码
func GetClosestVinsFrom(vin string, num int) []string {
	vins := make([]string, 0, num)
	if len(vin) != 17 {
		return vins
	}
	vin_8 := vin[0:8]
	year_code := GetYearCodeFromVin(vin)
	factory_code := GetFactoryCodeFromVin(vin)
	vin_last_6 := stringsx.Right(vin, 6)

	x, cur_serial_number, cur_serial_number_begin, cur_serial_number_end, err := split_vin_last_6(vin_last_6)

	if err != nil {
		return vins
	}
	closestSerialNumbers := getClosestSerialNumbersFrom(cur_serial_number, cur_serial_number_begin, cur_serial_number_end, num)
	//log.Println(closestSerialNumbers)
	//判断最后6位是否是数字
	for _, cur_serial_number := range closestSerialNumbers {
		vin_last_6 := get_vin_last_6_from(x, cur_serial_number)
		vin = getOneVinFrom(vin_8, year_code, factory_code, vin_last_6)
		vins = append(vins, strings.ToUpper(vin))
	}
	return vins
}

// 根据样本VIN码生成与之最为接近的N个VIN码,第9位为非校校位的情况
func GetClosestVinsNotFitForCheckRuleFrom(vin string, num int) []string {
	vins := make([]string, 0, num)
	if len(vin) != 17 {
		return vins
	}
	vin_8 := vin[0:8]
	year_code := GetYearCodeFromVin(vin)
	factory_code := GetFactoryCodeFromVin(vin)
	check_code := vin[8:9]
	vin_last_6 := stringsx.Right(vin, 6)

	x, cur_serial_number, cur_serial_number_begin, cur_serial_number_end, err := split_vin_last_6(vin_last_6)

	if err != nil {
		return vins
	}
	closestSerialNumbers := getClosestSerialNumbersFrom(cur_serial_number, cur_serial_number_begin, cur_serial_number_end, num)
	//log.Println(closestSerialNumbers)
	//判断最后6位是否是数字
	for _, cur_serial_number := range closestSerialNumbers {
		vin_last_6 := get_vin_last_6_from(x, cur_serial_number)
		//vin = getOneVinFrom(vin_8, year_code, factory_code, vin_last_6)
		vin = vin_8 + check_code + year_code + factory_code + vin_last_6
		vins = append(vins, strings.ToUpper(vin))
	}
	return vins
}

// 获得最为接近的相关序列号
func getClosestSerialNumbersFrom(cur_serial_number, cur_serial_number_begin, cur_serial_number_end, num int) []int {
	result := make([]int, 0, num)
	m, n := cur_serial_number, cur_serial_number
	for {
		n = n - 1
		if n >= cur_serial_number_begin {
			result = append(result, n)
		}
		if len(result) == num {
			sort.Ints(result)
			return result
		}
		m = m + 1
		//算完之后，没有达到数量，也返回
		if m > cur_serial_number_end {
			sort.Ints(result)
			return result
		}
		result = append(result, m)
		//达到相应的数量，则返回
		if len(result) == num {
			sort.Ints(result)
			return result
		}
	}
}

func getOneVinFrom(vin8, year_code, factory_code, vin_last_6 string) string {
	for _, vin9 := range strings.Split(Vin9strs, ",") {
		vin := vin8 + vin9 + year_code + factory_code + vin_last_6
		if IsVinLegal(vin) {
			return vin
		}
	}
	return ""
}

// 生成最后6位
func get_vin_last_6_from(x string, cur_serial_number int) string {
	//log.Println(x, cur_serial_number)
	s := strconv.Itoa(cur_serial_number)
	return x + strings.Repeat("0", 6-len(s)-len(x)) + s
}

// 把最后6位拆分成两个部分，0x8765 ==> "0x" 8765 0 9999
// 9876544 ==> "" 9876544 0 999999
func split_vin_last_6(vin_last_6 string) (x string, cur_serial_number, serial_number_begin, serial_number_end int, err error) {
	//前6位
	n, err := strconv.Atoi(vin_last_6)
	if err == nil {

		return "", n, 0, 999999, nil
	}
	//前5位
	n, err = strconv.Atoi(vin_last_6[1:])
	if err == nil {
		return vin_last_6[0:1], n, 0, 99999, nil
	}
	//前4位
	n, err = strconv.Atoi(vin_last_6[2:])
	if err == nil {
		return vin_last_6[0:2], n, 0, 9999, nil
	}
	//前3位
	n, err = strconv.Atoi(vin_last_6[3:])
	if err == nil {
		return vin_last_6[0:3], n, 0, 999, nil
	}
	//前2位
	n, err = strconv.Atoi(vin_last_6[4:])
	if err == nil {
		return vin_last_6[0:4], n, 0, 99, nil
	}
	//前1位
	n, err = strconv.Atoi(vin_last_6[5:])
	if err == nil {
		return vin_last_6[0:5], n, 0, 9, nil
	}
	return x, cur_serial_number, 0, 0, fmt.Errorf("此为非法VIN 码")
}
