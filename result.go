package validator

import (
	"strconv"
	"strings"
)

//校验结果
type Result map[string][]string

func (this Result) String() string {
	var buf strings.Builder
	i := 1
	for _, rr := range this {
		for _, r := range rr {
			buf.WriteString(strconv.Itoa(i) + "、" + r + "\n")
			i++
		}
	}
	return buf.String()
}
