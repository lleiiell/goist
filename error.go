package goist

import "fmt"

// code err

const ErrCodeCommon = 1001
const ErrCodeParam = 1002
const ErrCodeAuth = 1003

// component err

const ErrCodeDb = 2001
const ErrCodeRedis = 2002
const ErrCodeMySQL = 2003

// external err

const ErrCodeApi = 3001

type Err struct {
	ErrCode interface{}
	ErrMsg  interface{}
}

func (e Err) Error() string {
	return fmt.Sprintf("ErrCode: %v; ErrMsg: %v", e.ErrCode, e.ErrMsg)
}
