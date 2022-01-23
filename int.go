package conversion

import (
	"strconv"
)

func IntToString(source int) (target string) {
	target = strconv.FormatInt(int64(source), 10)
	return
}
