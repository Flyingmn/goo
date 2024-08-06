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

func ArrayColumn[T []M, M map[K]V, K string, V any](arr T, k K) []V {
	var result []V

	for _, v := range arr {
		if v, ok := v[k]; ok {
			result = append(result, v)
		}
	}

	return result
}

func ArrayChunk[T []V, V any](s T, size int) []T {
	if len(s) <= size {
		return []T{s}
	}

	chunks := int(len(s) / size)

	var result []T

	for i := 0; i < chunks; i++ {
		start := i * size
		end := start + size

		result = append(result, s[start:end])
	}

	return result
}
