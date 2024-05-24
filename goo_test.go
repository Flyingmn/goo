package goo_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Flyingmn/goo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/go-playground/assert.v1"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", goo.Md5("123456"))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f8831", goo.Md5("123456"))
}

func TestZap(t *testing.T) {
	goo.SetZapCfg(goo.ZapLevel("debug"), goo.ZapOutFile("./log/test.log", goo.ZapOutFileMaxSize(128), goo.ZapOutFileMaxAge(7)))
	goo.Zap().Info("hello world", zap.String("name", "zhangsan"), zap.Any("age", 18))
	goo.Sap().Infow("hello world", "name", "zhangsan", "age", 18)
}

func TestCtxBindWithTagError(t *testing.T) {
	type RequestStruct struct {
		ID  int64 `form:"id" json:"id" binding:"required" msg:"id不能为空"`
		Age int64 `form:"age" json:"age" binding:"required,gt=18" required_msg:"id不能为空" gt_msg:"年龄必须大于18"`
	}

	var req RequestStruct

	//模拟一个gin的context
	// 创建一个模拟的HTTP请求

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"id":0,"age":16}`)))
	// 设置请求的内容类型为JSON
	c.Request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, c.Request)

	err := goo.CtxBindWithTagError(c, &req)

	fmt.Printf("%+v\n\n", err.Error())
}
