package validator

import (
	"strconv"
	"strings"
)

//规则
type Rule struct {
	name    string
	message []string
	param   map[string][]string
}

func newRule(rule string, message ...string) *Rule {
	if len(message) == 0 {
		panic("NewRule: message is a required parameter")
	}
	tmp := new(Rule)
	tmp.message = make([]string, len(message))
	copy(tmp.message, message)
	tmp.parse(rule)
	return tmp
}

//解析规则名称与规则参数，
//rule的范例 ruleName:param1=1&param2=get,post&param3=a,b,\,,\&,\=,c,:
//参数部分解析后的结果是 map[param1:[1] param2:[get post] param3:[a b , & = c :]]
func (this *Rule) parse(rule string) {
	//从规则字符串中切割出规则的名称
	s := strings.SplitN(rule, ":", 2)
	this.name = strings.Trim(s[0], " ")
	if this.name == "" {
		panic("NewRule: rule is a bad parameter")
	}
	this.param = map[string][]string{}
	if len(s) <= 1 {
		return
	}
	//从规则字符串中解析出规则的参数
	and := "ghm_And_Shadow"
	equal := "ghm_Equal_Shadow"
	comma := "ghm_Comma_Shadow"
	s[1] = strings.Replace(s[1], `\&`, and, -1)
	s[1] = strings.Replace(s[1], `\=`, equal, -1)
	s[1] = strings.Replace(s[1], `\,`, comma, -1)
	param := strings.Split(s[1], "&")
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
		tmpValue := strings.Split(tmp[1], ",")
		value := make([]string, 0, len(tmpValue))
		for _, v := range tmpValue {
			v := strings.Trim(v, " ")
			if v != "" {
				if v == and {
					v = "&"
				} else if v == equal {
					v = "="
				} else if v == comma {
					v = ","
				}
				value = append(value, v)
			}
		}
		if len(value) > 0 {
			this.param[tmp[0]] = value
		}
	}
}

func (this Rule) Clone() *Rule {
	tmp := &Rule{}
	tmp.name = this.name
	tmp.message = make([]string, len(this.message))
	copy(tmp.message, this.message)
	tmp.param = map[string][]string{}
	for k, v := range this.param {
		copyV := make([]string, len(v))
		copy(copyV, v)
		tmp.param[k] = copyV
	}
	return tmp
}

func (this Rule) Message(index int) string {
	if l := len(this.message); index >= 0 && index <= (l-1) {
		return this.message[index]
	}
	return this.message[0]
}

func (this Rule) Name() string {
	return this.name
}

func (this Rule) Has(key string) bool {
	_, ok := this.param[key]
	return ok
}

func (this Rule) HasIn(key string, value string) bool {
	tmp, ok := this.param[key]
	if ok {
		for _, v := range tmp {
			if v == value {
				return true
			}
		}
	}
	return false
}

func (this Rule) Get(key string) []string {
	if tmp, ok := this.param[key]; ok {
		result := make([]string, len(tmp))
		copy(result, tmp)
		return result
	}
	return []string{}
}

func (this Rule) GetString(key string, def ...string) string {
	if tmp, ok := this.param[key]; ok {
		return tmp[0]
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

func (this Rule) GetInt(key string, def ...int) int {
	if tmp, ok := this.param[key]; ok {
		if n, err := strconv.Atoi(tmp[0]); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (this Rule) GetFloat(key string, def ...float64) float64 {
	if tmp, ok := this.param[key]; ok {
		if n, err := strconv.ParseFloat(tmp[0], 64); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func (this Rule) GetBool(key string, def ...bool) bool {
	if tmp, ok := this.param[key]; ok {
		if n, err := strconv.ParseBool(tmp[0]); err == nil {
			return n
		}
	}
	if len(def) > 0 {
		return def[0]
	}
	return false
}
