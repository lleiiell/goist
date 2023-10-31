package goist

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
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
