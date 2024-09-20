package goist

import (
	"encoding/csv"
	"os"
	"path"
)

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FileCreate(filename string) (*os.File, error) {
	flag := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	return os.OpenFile(filename, flag, 0640)
}

func CsvWrite(filename string, records [][]string) (err error) {
	exists, errE := pathExists(filename)
	if errE != nil {
		return errE
	}

	var f *os.File
	if !exists {
		f, err = FileCreate(filename)
		if err != nil {
			return err
		}
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)
	if err != nil {
		return
	}
	w.Flush()
	if err = w.Error(); err != nil {
		return err
	}

	return nil
}

func PathClean(dir string) (err error) {
	dirEntry, errRead := os.ReadDir(dir)
	if errRead != nil {
		if os.IsNotExist(errRead) {
			return
		}
		return errRead
	}

	for _, d := range dirEntry {
		err = os.RemoveAll(path.Join(dir, d.Name()))
		if err != nil {
			return
		}
	}
	return
}
