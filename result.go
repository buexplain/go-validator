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

//校验中产生的错误
type ErrorBag map[string][]error

func (this ErrorBag) Error() string {
	var buf strings.Builder
	i := 1
	for _, errs := range this {
		for _, err := range errs {
			buf.WriteString(strconv.Itoa(i) + "、" + err.Error() + "\n")
			i++
		}
	}
	return buf.String()
}
