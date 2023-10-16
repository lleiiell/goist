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
func TestFileExists(t *testing.T) {
	y, err := FileExists("/root")
	fmt.Println("y, err: ", y, err)

	y2, err2 := FileExists("/root/Music")
	fmt.Println("y2, err2: ", y2, err2)

	y3, err3 := FileExists("/lost+found")
	fmt.Println("y3, err3: ", y3, err3)
}