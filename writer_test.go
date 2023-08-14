package goist_test

import (
	"github.com/lleiiell/goist"
	"testing"
	"time"
)

func TestWriteBin(t *testing.T) {
	err := goist.WriteBin("/tmp/TestWriteBin", time.Now().Unix())
	if err != nil {
		t.Error(err)
		return
	}
}
