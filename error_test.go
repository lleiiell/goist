package goist

import (
    "fmt"
)

func ExampleErr_Error() {
    e := func() error {
        return Err{
            ErrCode: 123,
            ErrMsg:  "test err",
        }
    }()

    ee := e.(Err)

    fmt.Println(ee.Error())

    // Output: ErrCode: 123; ErrMsg: test err
}
