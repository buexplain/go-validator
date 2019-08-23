package validator

import (
	"strconv"
	"strings"
)

//规则
type Rule struct {
	name    string
	message string
	param   map[string]string
}

func NewRule(rule string, message string) *Rule {
	obj := new(Rule)
	obj.message = message

	s := strings.SplitN(rule, ":", 2)
	obj.name = strings.Trim(s[0], " ")
	obj.param = map[string]string{}
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
			obj.param[tmp[0]] = tmp[1]

		}
	}

	return obj
}

func (this Rule) Clone() *Rule {
	rule := &Rule{}
	rule.name = this.name
	rule.message = this.message
	rule.param = map[string]string{}
	for k,v := range this.param {
		rule.param[k] = v
	}
	return rule
}

func (this Rule) Message () string  {
	return this.message
}

func (this Rule) Name () string {
	return this.name
}

func (this Rule) Has(param string) bool {
	_, ok := this.param[param]
	return ok
}

func (this Rule) GetString(param string, def ...string) string {
	if tmp, ok := this.param[param]; ok {
		return tmp
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (this Rule) GetInt(param string, def ...int) int {
	if tmp, ok := this.param[param]; ok {
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
	if tmp, ok := this.param[param]; ok {
		if n, err := strconv.ParseFloat(tmp, 64); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}
