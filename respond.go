package goo

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(code int, msg string, data interface{}) Resp {
	return Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewRespOK(data interface{}) Resp {
	return NewResp(0, "success", data)
}

func NewRespError(msg string, data interface{}) Resp {
	return NewResp(50000, msg, data)
}

// 自定义状态码
func RespJson(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, NewResp(code, msg, data))
}

// http200且成功
func RespJsonOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewRespOK(data))
}

// http400且错误
func RespJsonError(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, NewRespError(msg, data))
}

// http200但错误
func RespJsonFail(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, NewRespError(msg, data))
}

/*
从结构体中获取错误信息:

	type RequestStruct struct {
		ID        int64  `form:"id" json:"id" binding:"required" msg:"id不能为空"`
	}

	var req RequestStruct

	if err := CtxBindWithTagError(ctx, &req); err != nil {
		// err.Error() == "id不能为空"
		return
	}
*/
func CtxBindWithTagError(ctx *gin.Context, ptr any) error {

	if err := ctx.ShouldBind(ptr); err != nil {
		return TagError(err, ptr)
	}

	return nil
}

// 自定义错误消息
func TagError(err error, ptr any) error {

	if reflect.TypeOf(err).Elem().String() == "json.UnmarshalTypeError" {
		errs := err.(*json.UnmarshalTypeError)

		return errs
	}

	errs := err.(validator.ValidationErrors)
	errStr := ""

	ptrType := reflect.TypeOf(ptr)

	if ptrType.Kind() == reflect.Ptr {
		ptrType = ptrType.Elem()
	}

	for _, fieldError := range errs {
		filed, _ := ptrType.FieldByName(fieldError.Field())

		errTag := fieldError.Tag() + "_msg"
		// 获取对应binding得错误消息
		errTagText := filed.Tag.Get(errTag)
		// 获取统一错误消息
		errText := filed.Tag.Get("msg")
		if errTagText != "" {
			errStr += errTagText + "; "
		} else if errText != "" {
			errStr += errText + "; "
		} else {
			errStr += fieldError.Field() + ":" + fieldError.Tag() + "; "
		}
	}

	return errors.New(errStr)
}
