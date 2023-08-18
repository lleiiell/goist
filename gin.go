package goist

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// GinReverseProxy gin 反向代理
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
	}

	proxy.ServeHTTP(c.Writer, c.Request)

	return nil
}
