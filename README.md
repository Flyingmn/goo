## 简介

封装一些go常用的函数或者操作

## 安装

```bash
# 使用go modules
go get -u github.com/Flyingmn/goo
```

## 使用

### 日志zap
```go

// 默认info级别，如果要自定义配置（注意要在Zap()以及Sap()之前），则执行
// goo.SetZapCfg(goo.ZapLevel("debug"), goo.ZapOutFile("./log/test.log", goo.ZapOutFileMaxSize(128), goo.ZapOutFileMaxAge(7)))

//普通zap *zap.Logger
goo.Zap().Info("hello world", zap.String("name", "zhang"), zap.Any("age", 18))

//带语法糖的zap *zap.SugaredLogger
goo.Sap().Infow("hello world", "name", "zhang", "age", 18)

// {"level":"info","time":"2024-05-21 16:29:20.927","line":"...testing/testing.go:1595","func":"testing.tRunner","msg":"hello world","name":"zhang","age":18}
```

### 数据验证
```go
//从结构体中获取错误信息:

type RequestStruct struct {
	ID  int64 `form:"id" json:"id" binding:"required" msg:"id不能为空"`
	Age int64 `form:"age" json:"age" binding:"required,gt=18" required_msg:"年龄不能为空" gt_msg:"年龄必须大于18"`
}

var req RequestStruct

// ctx： `{"id":0,"age":16}`
if err := goo.CtxBindWithTagError(ctx, &req); err != nil {
    // err.Error() == "id不能为空, 年龄必须大于18"
    return
}
```