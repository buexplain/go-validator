package validator

import (
	"strconv"
	"strings"
)

//规则
type Rule struct {
	Name    string
	Message string
	Param   map[string]string
}

func NewRule(rule string, message string) *Rule {
	obj := new(Rule)
	obj.Message = message

	s := strings.SplitN(rule, ":", 2)
	obj.Name = strings.Trim(s[0], " ")
	obj.Param = map[string]string{}
	if len(s) > 1 {
		param := strings.Split(s[1], ",")
		for _, v := range param {
			tmp := strings.SplitN(v, "=", 2)
			if len(tmp) != 2 {
				continue
			}
			tmp[0] = strings.Trim(tmp[0], " ")
			tmp[1] = strings.Trim(tmp[1], " ")
			if tmp[0] == "" || tmp[1] == "" {
				continue
			}
			obj.Param[tmp[0]] = tmp[1]

		}
	}

	return obj
}

func (this Rule) GetString(param string, def ...string) string {
	if tmp, ok := this.Param[param]; ok {
		return tmp
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (this Rule) GetInt(param string, def ...int) int {
	if tmp, ok := this.Param[param]; ok {
		if n, err := strconv.Atoi(tmp); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (this Rule) GetFloat(param string, def ...float64) float64 {
	if tmp, ok := this.Param[param]; ok {
		if n, err := strconv.ParseFloat(tmp, 64); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}
