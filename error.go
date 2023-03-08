package goist

import "fmt"

type Err struct {
    ErrCode interface{}
    ErrMsg  interface{}
}

func (e Err) Error() string {
    return fmt.Sprintf("ErrCode: %v; ErrMsg: %v", e.ErrCode, e.ErrMsg)
}
