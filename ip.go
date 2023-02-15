package go_tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsIpv4IP(ip string) bool {
	ipArray := strings.Split(ip, ".")
	if len(ipArray) != 4 {
		return false
	}
	for _, item := range ipArray {
		num, err := strconv.Atoi(item)
		if err != nil {
			fmt.Printf("err is :%v \n", err.Error())
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
	}
	return true
}

func IsIpByZz(ip string) {
	ips, err := regexp.MatchString("[0-9].{1,}[0-9]", ip)
	if err != nil {
		fmt.Printf("err is: %v \n", err.Error())
	}
	fmt.Printf("val is :%v \n", ips)

}
