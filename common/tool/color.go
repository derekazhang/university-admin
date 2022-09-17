package tool

import (
	"strconv"
	"strings"
)

//"rgba(144,19,254,1)"
func GetRgba(rgba string) (out []int) {
	s1 := strings.Replace(rgba, "rgba(", "", -1)
	s1 = strings.Replace(s1, ")", "", -1)
	for _, v := range strings.Split(s1, ",") {
		rgbaNum, _ := strconv.Atoi(v)
		out = append(out, rgbaNum)
	}
	return
}
