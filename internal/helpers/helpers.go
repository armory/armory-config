// This package contains functions that are use on different parts of the code
package helpers

import "strconv"

func IntToString(value int32) string {
	return strconv.FormatInt(int64(value), 10)
}
