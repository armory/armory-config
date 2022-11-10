// This package contains functions that are use on different parts of the code
package helpers

import (
	"strconv"
	"strings"
)

func IntToString64(value int64) string {
	return strconv.FormatInt(value, 10)
}

func PrintFmtInt64(key string, value int64, amountTabs int, addNewline bool) string {
	return PrintFmtStr(key, IntToString64(value), amountTabs, addNewline)
}

func IntToString(value int32) string {
	return strconv.FormatInt(int64(value), 10)
}

func PrintFmtInt(key string, value int32, amountTabs int, addNewline bool) string {
	return PrintFmtStr(key, IntToString(value), amountTabs, addNewline)
}

func PrintFmtBool(key string, value bool, amountTabs int, addNewline bool) string {
	return PrintFmtStr(key, strconv.FormatBool(value), amountTabs, addNewline)
}

func PrintFmtStrApostrophe(key string, value string, amountTabs int, addNewline bool) string {
	str := ""

	if "" != value && 0 != len(value) {
		str = PrintFmtStr(key, "'"+value+"'", amountTabs, addNewline)
	}

	return str
}

func PrintFmtStrDoubleQuote(key string, value string, amountTabs int, addNewline bool) string {
	str := ""

	if "" != value && 0 != len(value) {
		str = PrintFmtStr(key, "\""+value+"\"", amountTabs, addNewline)
	}

	return str
}

// This function only prints key and value if the value is not empty
func PrintFmtStr(key string, value string, amountTabs int, addNewline bool) string {
	str := ""

	//TODO
	// add validation for key to have a space after :
	// remove all spaces from key and make first letter lowercase
	// key = makeFirstLetterLowerCase(key)

	if "" != value && 0 != len(value) {
		if addNewline {
			str += "\n"
		}

		str += GetSpaces(amountTabs*2) + key + value
	}
	return str
}

// This function returns a string with the amount of spaces
func GetSpaces(amountSpaces int) string {
	str := ""
	for i := 0; i < amountSpaces; i++ {
		str += " "
	}
	return str
}

func makeFirstLetterLowerCase(str string) string {
	return strings.ToLower(str[0:1]) + str[1:]
}
