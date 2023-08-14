package goist

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

func WriteBin(filename string, data interface{}) (err error) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	b := func() []byte {
		bt, _ := json.Marshal(data)
		return bt
	}()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, b)
	if err != nil {
		return
	}
	_, err = fp.Write(buf.Bytes())

	return
}
