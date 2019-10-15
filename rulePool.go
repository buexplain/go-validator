package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//校验函数
type FN func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error)

//规则池
var rulePool map[string]FN = map[string]FN{}

//从规则池中取出一条规则
func Pool(ruleName string) FN {
	if r, ok := rulePool[ruleName]; ok {
		return r
	}
	return nil
}

//校验非零值
func init() {
	rulePool["required"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		if value == nil {
			return rule.Message(0), nil
		}
		if isEmpty(value) {
			return rule.Message(0), nil
		}
		return "", nil
	}
}

//校验范围 min与max是一段连续的区间
func init() {
	rulePool["between"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		if value == nil {
			return rule.Message(0), nil
		}
		minFloat := rule.GetFloat("min")
		maxFloat := rule.GetFloat("max")
		min := int(minFloat)
		max := int(maxFloat)

		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Array, reflect.Map, reflect.Slice:
			inLen := rv.Len()
			if !(inLen >= min && inLen <= max) {
				return rule.Message(0), nil
			}
		case reflect.String:
			inLen := len([]rune(value.(string)))
			if !(inLen >= min && inLen <= max) {
				return rule.Message(0), nil
			}
		case reflect.Int:
			in := value.(int)
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Int8:
			in := int(value.(int8))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Int16:
			in := int(value.(int16))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Int32:
			in := int(value.(int32))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Int64:
			in := int(value.(int64))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Uint:
			in := int(value.(uint))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Uint8:
			in := int(value.(uint8))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Uint16:
			in := int(value.(uint16))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Uint32:
			in := int(value.(uint32))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Uint64:
			in := int(value.(uint64))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Uintptr:
			in := int(value.(uintptr))
			if !(in >= min && in <= max) {
				return rule.Message(0), nil
			}
		case reflect.Float32:
			in := float64(value.(float32))
			if !(in >= minFloat && in <= maxFloat) {
				return rule.Message(0), nil
			}
		case reflect.Float64:
			in := value.(float64)
			if !(in >= minFloat && in <= maxFloat) {
				return rule.Message(0), nil
			}
		default:
			return "", errors.New("invalid type for between rule")

		}
		return "", nil
	}
}

//字母
func init() {
	rulePool["alpha"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		str := toString(value)
		if str == "" {
			return rule.Message(0), nil
		}
		if !regexAlpha.MatchString(str) {
			return rule.Message(1), nil
		}
		return "", nil
	}
}

//数字
func init() {
	rulePool["numeric"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		str := toString(value)
		if str == "" {
			return rule.Message(0), nil
		}
		if !regexNumeric.MatchString(str) {
			return rule.Message(1), nil
		}
		return "", nil
	}
}

//正整数
func init() {
	rulePool["positive_numeric"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		str := toString(value)
		if str == "" {
			return rule.Message(0), nil
		}
		if !regexPositiveNumeric.MatchString(str) {
			return rule.Message(1), nil
		}
		return "", nil
	}
}

//字母与数字
func init() {
	rulePool["alpha_numeric"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		str := toString(value)
		if str == "" {
			return rule.Message(0), nil
		}
		if !regexAlphaNumeric.MatchString(str) {
			return rule.Message(1), nil
		}
		return "", nil
	}
}

//字母与数字，以及破折号和下划线
func init() {
	rulePool["alpha_numeric_dash"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		str := toString(value)
		if str == "" {
			return rule.Message(0), nil
		}
		if !regexAlphaNumericDash.MatchString(str) {
			return rule.Message(1), nil
		}
		return "", nil
	}
}

//邮箱
func init() {
	rulePool["email"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		str := toString(value)
		if str == "" {
			return rule.Message(0), nil
		}
		if !regexEmail.MatchString(str) {
			return rule.Message(1), nil
		}
		return "", nil
	}
}

//数字、字母、特殊符号至少两种
func init() {
	rulePool["password"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		min := rule.GetInt("min", 8)
		max := rule.GetInt("max", 32)
		str := toString(value)
		if str == "" {
			return rule.Message(0), nil
		}

		runeStr := []rune(str)

		inLen := len(runeStr)
		if !(inLen >= min && inLen <= max) {
			return rule.Message(1), nil //密码格式有误，请输入8位以上32位以下的密码
		}

		//字母
		letter := false
		//数字
		number := false
		//特殊符号
		char := false
		//键盘上的ASCII以外的字符，比如 ￥
		charArr := [...]rune{65281, 65509, 8230, 8230, 65288, 65289,
			8212, 8212, 65292, 12290, 65307, 8216, 12289, 12305, 12304}
		in := func(char rune) bool {
			for _, v := range charArr {
				if char == v {
					return true
				}
			}
			return false
		}

		for _, v := range runeStr {
			if v >= 48 && v <= 57 {
				number = true
			} else if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
				letter = true
			} else if (v >= 33 && v <= 126) || in(v) {
				char = true
			} else {
				return rule.Message(2), nil //密码格式有误，请输入数字、字母、符号
			}
		}

		result := 0

		if letter {
			result += 1
		}
		if number {
			result += 1
		}
		if char {
			result += 1
		}

		if result >= 2 {
			return "", nil
		} else {
			return rule.Message(3), nil //密码格式有误，数字、字母、符号至少两种
		}
	}
}

//校验范围，这个范围不是一段连续的区间，类似白名单
func init() {
	rulePool["in"] = func(field string, value interface{}, rule *Rule, structVar interface{}) (string, error) {
		if value == nil {
			return rule.Message(0), nil
		}
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Slice:
			if rv.Len() == 0 {
				return rule.Message(0), nil
			}
			if arr, ok := value.([]string); ok {
				for _, v := range arr {
					if !rule.HasIn("in", v) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]int); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(v)) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]float64); ok {
				for _, v := range arr {
					if !rule.HasIn("in", fmt.Sprintf("%v", v)) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]int8); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]int16); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]int32); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]int64); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]uint); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]uint8); ok {

				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]uint16); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]uint32); ok {
				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]uint64); ok {

				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]uintptr); ok {

				for _, v := range arr {
					if !rule.HasIn("in", strconv.Itoa(int(v))) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]float32); ok {

				for _, v := range arr {
					if !rule.HasIn("in", fmt.Sprintf("%v", v)) {
						return rule.Message(1), nil
					}
				}
			} else if arr, ok := value.([]interface{}); ok {
				for _, v := range arr {
					if !rule.HasIn("in", fmt.Sprintf("%v", v)) {
						return rule.Message(1), nil
					}
				}
			} else {
				return "", errors.New("invalid type for in rule")
			}
		default:
			str := toString(value)
			if str == "" {
				return rule.Message(0), nil
			}
			//支持切割value后再进行校验
			split := rule.GetString("split")
			if split != "" {
				strArr := strings.Split(str, split)
				for _, v := range strArr {
					if !rule.HasIn("in", v) {
						return rule.Message(1), nil
					}
				}
			} else {
				if !rule.HasIn("in", str) {
					return rule.Message(1), nil
				}
			}
			return "", nil
		}
		return "", nil
	}
}
