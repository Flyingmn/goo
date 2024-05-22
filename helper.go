package goo

import (
	"crypto/md5"
	"fmt"
)

func Md5(input string) string {
	hasher := md5.New()         // 创建一个md5 Hash对象
	hasher.Write([]byte(input)) // 将字符串写入hasher
	digest := hasher.Sum(nil)   // 计算哈希并获取16字节的结果

	// 将16字节的哈希转换为16进制字符串
	MD5String := fmt.Sprintf("%x", digest)

	return MD5String
}
