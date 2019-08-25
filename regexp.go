package validator

import "regexp"

const (
	//纯字符
	alpha string = "^[a-zA-Z]+$"
	//纯数字
	numeric string = "^[0-9]+$"
	//字符与数字
	alphaNumeric string = "^[a-zA-Z0-9]+$"
	//字母数字字符，以及破折号和下划线
	alphaNumericDash string = "^[a-zA-Z0-9_-]+$"
	//邮箱
	email string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$"
)

var (
	regexAlpha            = regexp.MustCompile(alpha)
	regexNumeric          = regexp.MustCompile(numeric)
	regexAlphaNumeric     = regexp.MustCompile(alphaNumeric)
	regexAlphaNumericDash = regexp.MustCompile(alphaNumericDash)
	regexEmail            = regexp.MustCompile(email)
)
