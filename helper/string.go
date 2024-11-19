package helper

import "strconv"

func StringToBool(str string) bool {
	convBool, _ := strconv.ParseBool(str)
	return convBool
}
