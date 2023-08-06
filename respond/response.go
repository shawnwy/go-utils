// Package respond is a helper for gin response
// It encourages to use package go-utils/errors for error code definition.
// All method from this package is potentially use the GetCode from package go-utils/errors
package respond

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shawnwy/go-utils/v5/errors"
	"net/http"
)

var (
	textCodecs = map[Lang]map[int]string{}
	lang       Lang
)

type Lang string

const (
	EN_US Lang = "en-us"
	ZH_CN Lang = "zh-cn"
)

func isLangSet() bool {
	return lang != ""
}

// AddLanguageTranslation add translation for lang language
func AddLanguageTranslation(lang Lang, messages map[int]string) {
	codecs := textCodecs[lang]
	if codecs == nil {
		codecs = make(map[int]string)
	}
	for code, msg := range messages {
		codecs[code] = msg
	}
}

// SetLang set the language of respond error msg
func SetLang(l Lang) {
	lang = l
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"messages,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type builder struct {
	c          *gin.Context
	statusCode int
}

func New(c *gin.Context) *builder {
	return &builder{
		c:          c,
		statusCode: http.StatusOK,
	}
}

func (b *builder) Succeed() *builder {
	b.statusCode = http.StatusOK
	return b
}

func (b *builder) NotFound(err error) {
	b.c.JSON(http.StatusNotFound, Response{
		Code: errors.GetCode(err),
		Msg:  err.Error(),
	})
}

// StatusCreated is 201 in http status code
func (b *builder) StatusCreated() {
	b.c.JSON(http.StatusCreated, Response{
		Code: http.StatusCreated,
	})
}

func (b *builder) WithData(data interface{}, msg string) {
	b.c.JSON(b.statusCode, Response{
		Code: b.statusCode,
		Msg:  msg,
		Data: data,
	})
}

func (b *builder) WithError(err error, msg string) {
	b.c.JSON(b.statusCode, Response{
		Code: errors.GetCode(err),
		Msg:  fmt.Sprintf("%s: %s", err.Error(), msg),
	})
}
