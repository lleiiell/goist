package goist

import (
	"fmt"
	"testing"
)

/*
y, err:  true <nil>
y2, err2:  false stat /root/Music: permission denied
y3, err3:  true <nil>
*/
func TestPathExists(t *testing.T) {
	y, err := PathExists("/root")
	fmt.Println("y, err: ", y, err)

	y2, err2 := PathExists("/root/Music")
	fmt.Println("y2, err2: ", y2, err2)

	y3, err3 := PathExists("/lost+found")
	fmt.Println("y3, err3: ", y3, err3)
}

func TestCsvWrite(t *testing.T) {
	records := [][]string{
		{"a", "aa", "aaa"},
		{"b", "bb", "bbb"},
	}

	err := CsvWrite("/tmp/csvWrite.csv", records)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestPathClean(t *testing.T) {
	err := PathClean("/tmp/123/")
	if err != nil {
		t.Error(err)
	}
}
