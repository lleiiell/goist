package goist

import (
    "bytes"
    "fmt"
    "github.com/google/uuid"
    "io"
    "log"
    "net/http"
    "net/url"
    "reflect"
    "strings"
    "time"
)

// Get Http Get
// query array： convert map["name1"]["v1", "v2"] to name1[]=v1&name1[]=v2
func Get(url string, params map[string]interface{}, headers map[string]string, timeout time.Duration) (ret []byte, err error) {

    var query string
    query = BuildQuery(params)

    if strings.Contains(url, "?") {
        url += query
    } else {
        url += "?" + query
    }

    return request(url, nil, http.MethodGet, headers, timeout)
}

// BuildQuery query parse
func BuildQuery(m map[string]interface{}) (q string) {

    u, _ := url.Parse("")
    qu := u.Query()

    for k, v := range m {
        item := reflect.ValueOf(v)
        switch item.Kind() {
        case reflect.Slice, reflect.Array:
            for i := 0; i < item.Len(); i++ {
                rf := reflect.ValueOf(item.Index(i))
                qu.Set(fmt.Sprintf("%s[]", k), fmt.Sprintf("%v", rf.Interface()))
            }
        default:
            qu.Set(k, fmt.Sprintf("%v", v))
        }
    }

    u.RawQuery = qu.Encode()

    q = u.RawQuery
    return
}

// Post HTTP POST
func Post(url string, rq []byte, headers map[string]string, timeout time.Duration) (rp []byte, err error) {
    return request(url, rq, http.MethodPost, headers, timeout)
}

// Put HTTP PUT
func Put(url string, rq []byte, headers map[string]string, timeout time.Duration) (rp []byte, err error) {
    return request(url, rq, http.MethodPut, headers, timeout)
}

// Delete HTTP DELETE
func Delete(url string, rq []byte, headers map[string]string, timeout time.Duration) (rp []byte, err error) {
    return request(url, rq, http.MethodDelete, headers, timeout)
}

func request(url string, rq []byte, method string, headers map[string]string, timeout time.Duration) (body []byte, err error) {

    hUuid := uuid.New().String()

    // "API request：", "http-uuid", hUuid, "url", u, "rq", string(rq), "headers", headers
    log.Println(fmt.Sprintf("API request：http-uuid:%s; url:%s; req:%s; headers:%v", hUuid, url, string(rq), headers))
    req, err := http.NewRequest(method, url, bytes.NewBuffer(rq))

    // fix EOF err, disable KeepAlive
    req.Close = true

    if err != nil {
        return
    }
    req.Header.Set("Content-Type", "application/json")
    if headers != nil {
        for k, v := range headers {
            req.Header.Set(k, v)
        }
    }

    client := &http.Client{Timeout: timeout}
    resp, err := client.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close()

    body, _ = io.ReadAll(resp.Body)

    if _, ok := headers["do-not-log"]; !ok {
        // "API response", "http-uuid", hUuid, "body", string(body)
        log.Println(fmt.Sprintf("API response：http-uuid:%s; body:%s;", hUuid, string(body)))
    } else {
        delete(headers, "do-not-log")
    }

    return
}
