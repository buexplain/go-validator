package validator

import (
	"strconv"
	"strings"
)

//校验中产生的错误
type ErrorBag map[string][]error

func newErrorBag() ErrorBag {
	return ErrorBag{}
}

func (this ErrorBag) Add(field string, err error) {
	if tmp, ok := this[field]; ok {
		this[field] = append(tmp, err)
	} else {
		this[field] = []error{err}
	}
}

func (this ErrorBag) Error() string {
	var buf strings.Builder
	eof := "\n"
	for field, errors := range this {
		buf.WriteString(field)
		buf.WriteString(eof)
		for index, err := range errors {
			buf.WriteString("    ")
			buf.WriteString(strconv.Itoa(index + 1))
			buf.WriteString(". ")
			buf.WriteString(err.Error())
			buf.WriteString(eof)
		}
	}
	return buf.String()
}
