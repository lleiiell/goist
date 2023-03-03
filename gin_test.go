package goist

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
)

func ExampleGinReverseProxy() {
    engine := gin.New()

    addr := fmt.Sprintf("0.0.0.0:8080")

    engine.Any("/reverseProxy/:proxyPath", func(c *gin.Context) {
        remoteUrl := fmt.Sprintf("https://remoteServer.com/%s", c.Param("proxyPath"))
        err := GinReverseProxy(c, remoteUrl)

        if err != nil {
            log.Fatal(err)
        }
    })

    err := engine.Run(addr)

    if err != nil {
        log.Fatal(err)
    }

    // // Output:
}
