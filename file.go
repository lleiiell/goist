package goist

import "os"

func FileExists(filename string) (y bool, err error) {
	_, errStat := os.Stat(filename)
	if errStat != nil {
		if os.IsNotExist(errStat) {
			return
		}
		err = errStat
		return
	}

	y = true
	return
}
