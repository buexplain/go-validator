package validator

import (
	"strconv"
	"strings"
)

//校验结果
type Result struct {
	Message map[int][]string
	Field []string
}

func newResult() *Result {
	tmp := new(Result)
	tmp.Message = map[int][]string{}
	tmp.Field = []string{}
	return tmp
}

func (this *Result) IsEmpty() bool {
	return len(this.Field) == 0
}

func (this *Result) Add(field string, message string) *Result {
	for k, v := range this.Field {
		if v == field {
			this.Message[k] = append(this.Message[k], message)
			return this
		}
	}
	this.Field = append(this.Field, field)
	k := len(this.Field) - 1
	this.Message[k] = []string{message}
	return this
}


func (this Result) ToString(eof ...string) string {
	var buf strings.Builder
	var _eof []byte
	if len(eof) == 0 {
		_eof = []byte("<br>")
	}else {
		_eof = []byte(eof[0])
	}
	for index, field := range this.Field {
		buf.WriteString(field)
		buf.Write(_eof)
		for index, message := range this.Message[index] {
			buf.WriteString("    ")
			buf.WriteString(strconv.Itoa(index+1))
			buf.WriteString(". ")
			buf.WriteString(message)
			buf.Write(_eof)
		}
	}
	return buf.String()
}

func (this Result) ToSimpleString(eof ...string) string {
	var buf strings.Builder
	var _eof []byte
	if len(eof) == 0 {
		_eof = []byte("<br>")
	}else {
		_eof = []byte(eof[0])
	}
	counter := 0
	fieldLen := len(this.Field)
	for index, _ := range this.Field {
		messageLen := len(this.Message[index])
		for _, message := range this.Message[index] {
			if buf.Len() > 0 {
				buf.Write(_eof)
			}
			counter++
			if !(fieldLen == 1 && messageLen == 1) {
				buf.WriteString(strconv.Itoa(counter))
				buf.WriteString(". ")
			}
			buf.WriteString(message)
		}
	}
	return buf.String()
}

func (this Result) String() string {
	return this.ToString("\n")
}