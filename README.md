## 🍴 Goo
[![Go](https://github.com/Flyingmn/goo/actions/workflows/main.yml/badge.svg)](https://github.com/Flyingmn/goo/actions/workflows/main.yml) [![codecov](https://codecov.io/github/Flyingmn/goo/graph/badge.svg?token=UL045K7ESR)](https://codecov.io/github/Flyingmn/goo) [![Go Report Card](https://goreportcard.com/badge/github.com/Flyingmn/goo)](https://goreportcard.com/report/github.com/Flyingmn/goo) [![Go Reference](https://pkg.go.dev/badge/github.com/Flyingmn/goo.svg)](https://pkg.go.dev/github.com/Flyingmn/goo) ![Static Badge](https://img.shields.io/badge/License-MIT-blue)

一个实用的 Go 语言工具库，提供了一系列常用的辅助函数和工具类，简化日常开发工作。 <br>
👉️👉️👉️只用golang的内置包👈️👈️👈️

## 📦 安装
```bash
go get github.com/Flyingmn/goo
```

## 🚀 快速开始
```go
package main

import (
    "fmt"
    "github.com/Flyingmn/goo"
)

func main() {
    // 使用goo包中的数组去重
    uniqueSlices := goo.ArrayUnique([]int{1, 2, 2, 8, 2, 9, 10, 10, 10})
    // 输出去重后的结果
    fmt.Println(uniqueSlices)

    // goo_test.go有完整的测试用例    
    fmt.Println("goo_test.go有完整的测试用例，请自助食用！！！")
}
```
## 📖 文档

完整的 API 文档请访问: [pkg.go.dev/github.com/Flyingmn/goo](pkg.go.dev/github.com/Flyingmn/goo)


## 📎 模块信息

当前版本: v1.1.31<br>
<br>
Go 版本要求: 1.18 或更高<br>


## 🏗️ 项目结构

根据 Go 模块标准布局，该项目包含：
```text

github.com/Flyingmn/goo/
├── go.mod          # 模块定义
├── go.sum          # 依赖校验
├── *.go            # 源代码文件
├── goo_test.go     # 👉️👉️👉️测试代码文件👀👀👀,请参考这里的代码使用
└── README.md       # 项目说明
```

## 🤝 贡献
```text
欢迎提交 Issue 和 Pull Request！

Fork 本仓库

创建特性分支 (git checkout -b feature/AmazingFeature)

提交更改 (git commit -m 'Add some AmazingFeature')

推送到分支 (git push origin feature/AmazingFeature)

开启 Pull Request
```

## 📄 许可证

本项目采用 MIT 许可证 - 查看 LICENSE 文件了解详情


## 👤 作者

Flyingmn<br>

如果这个项目对您有帮助，请给个 ⭐️ 支持一下！


