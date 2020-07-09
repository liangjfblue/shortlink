/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 结果结构体
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功返回
func (r *Result) Success(c *gin.Context, data interface{}) {
	r.Code = 1
	r.Data = data
	r.Msg = "ok"

	c.JSON(http.StatusOK, Result{
		Code: 1,
		Data: data,
		Msg:  "ok",
	})
}

// Failure 失败返回
func (r *Result) Failure(c *gin.Context, errno *Errno) {
	c.JSON(http.StatusOK, Result{
		Code: 0,
		Data: map[string]interface{}{
			"code": errno.Code,
			"msg":  errno.Msg,
		},
		Msg: "error",
	})
}

var (
	_codes = map[int]struct{}{}
)

func New(e int) int {
	if e <= 0 {
		panic("code must greater than zero")
	}
	return add(e)
}

func add(e int) int {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return e
}

type Errno struct {
	Code int         `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data,omitempty"`
}

func (e Errno) Error() string {
	return fmt.Sprintf("code:%d, msg:%s", e.Code, e.Msg)
}

var (
	Success = &Errno{Code: New(1), Msg: "ok"}

	ErrParams          = &Errno{Code: New(100), Msg: "param error"}
	ErrLimitMiddleware = &Errno{Code: New(101), Msg: "ip over limit"}

	ErrCreateShortLink     = &Errno{Code: New(1000), Msg: "create short link error"}
	ErrShortCodeNoLongLink = &Errno{Code: New(1001), Msg: "shortCode no longLink"}
)
