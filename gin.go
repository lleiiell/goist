package goist

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"time"
)

// GinReverseProxy gin reverse proxy
func GinReverseProxy(c *gin.Context, remoteUrl string) error {
	remote, err := url.Parse(remoteUrl)
	if err != nil {
		return err
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host

		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path
		req.URL.RawQuery = c.Request.URL.RawQuery
	}

	proxy.ServeHTTP(c.Writer, c.Request)

	return nil
}

func GinRun(host, port string) error {

	engine := gin.New()

	engine.Use(logger(), recovery())

	route(engine)

	addr := fmt.Sprintf("%s:%s", host, port)

	log.Println("ginRunAddr", "addr", addr)

	return engine.Run(addr)

}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("traceUuid", Uuid32())
		c.Set("traceTimeStart", t)

		bs, _ := httputil.DumpRequest(c.Request, true)

		log.Println("http requests: ", "req", string(bs), "uuid", c.GetString("traceUuid"))

		c.Next()
	}
}

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				log.Printf("[Recovery from panic] error: %v, request: %v, stack: %v\n",
					err,
					string(httpRequest),
					string(debug.Stack()))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func route(r *gin.Engine) {
	r.Any("/listen", func(c *gin.Context) {
		c.String(200, "Success")
	})
}
