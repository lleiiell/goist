package goist

import (
	"fmt"
	"strconv"
)

// IdMask Generate irregular, reversible masks
func IdMask(id int64) (str string) {
	return fmt.Sprintf("%s%s", strconv.FormatInt(id, 36), RandStr(3))
}

func IdUnmask(mask string) (id int64) {
	runeSli := []rune(mask)
	if len(runeSli) <= 3 {
		return 0
	}

	id, _ = strconv.ParseInt(string(runeSli[:len(runeSli)-3]), 36, 64)
	return
}
