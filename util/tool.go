package util

import (
  "net/http"
  "net/url"
  "io/ioutil"
  "github.com/gin-gonic/gin"
  "bytes"
  "strings"
  "regexp"
  "vq0599/conf"
)

// 转发请求
func NewRequestForward(c *gin.Context, url string) (*http.Response, error) {
  requestBody, _ := ioutil.ReadAll(c.Request.Body)

  proxyRequest, _ := http.NewRequest(c.Request.Method, url, bytes.NewReader(requestBody))
  proxyRequest.Header = c.Request.Header

  return http.DefaultClient.Do(proxyRequest)
}

// golang 切割空字符串依然会得到包含一个空元素的数组 => [""]
func Split(str, sep string) []string {
  if str != "" {
    return strings.Split(str, sep)
  } else {
    return make([]string, 0)
  }
}

func HtmlToPureText(html string) string {
  reHTML, _ := regexp.Compile("<[^>]*>")
  return reHTML.ReplaceAllString(html, "")
}

func Min(x, y int) int {
  if x < y {
      return x
  }
  return y
}

// 截取文本 支持中文
func SubString(src string, start, end int) string {
  runSrc := []rune(src)
  maxLen := len(runSrc)
  validEnd := Min(maxLen, end)
  return string(runSrc[start: validEnd])
}

// 设置Cookie
func SetCookie(c *gin.Context, name, value string) {
  http.SetCookie(c.Writer, &http.Cookie{
    Name:     name,
    Value:    url.QueryEscape(value),
    MaxAge:   conf.COOKIES_MAXAGE,
    Path:     "/",
    Domain:   conf.COOKIES_DOMAIN,
    Secure:   false,
    HttpOnly: true,
  })
}