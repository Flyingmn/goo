package goo

import (
	"errors"
	"math/rand"
	"reflect"
)

// 接收一个map的切片和一个key，返回所有map中key对应的value所组成的切片
func ArrayColumn[T ~[]M, M ~map[K]V, K comparable, V any](arr T, k K) []V {
	result := make([]V, 0, len(arr))

	for _, v := range arr {
		if v1, ok := v[k]; ok {
			result = append(result, v1)
		}
	}

	return result
}

// 接收一个结构体切片和一个以结构体为参数的函数，返回所有结构体执行函数后的返回值所组成的切片
func StructsColumn[T any, V any](structs []T, kefFunc func(T) V) []V {
	var result []V

	for _, v := range structs {
		result = append(result, kefFunc(v))
	}

	return result
}

// 将一个切片分割成指定size长度的快，返回一个由块组成的切片
func ArrayChunk[T ~[]V, V any](s T, size int) []T {
	if len(s) <= size {
		return []T{s}
	}

	chunks := int(len(s) / size)

	lastChunkSize := len(s) % size

	var result []T

	for i := 0; i < chunks; i++ {
		start := i * size
		end := start + size

		result = append(result, s[start:end])
	}

	if lastChunkSize > 0 {
		start := len(s) - lastChunkSize
		result = append(result, s[start:])
	}

	return result
}

// 接收一个map,返回map的所有key组成的切片
func ArrayKeys[T ~map[K]V, K comparable, V any](arr T) []K {
	result := make([]K, 0, len(arr))

	for k := range arr {
		result = append(result, k)
	}

	return result
}

// 接收一个结构体,返回结构体的字段名所组成的切片； 优先取tag中gorm的column标签，如果gorm明示-，则跳过， 如果没有则取json， 如果json明示-， 则跳过, 如果没有，则跳过; 如果没有gorm和json，并且也没有标示 - 则取字段名 可以接受结构体或结构体指针
func StructKeys[T any](s T) ([]string, error) {
	if !IsStruct(s) {
		return nil, errors.New("data is't struct")
	}

	val := reflect.ValueOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	numFields := val.NumField()

	result := make([]string, 0, numFields)

	for i := 0; i < numFields; i++ {
		field := val.Type().Field(i)
		tag := field.Tag

		if tag.Get("gorm") == "-" {
			continue
		}

		columnStr, _ := ParseGormColumnTag(tag)

		if columnStr != "" {
			result = append(result, columnStr)
			continue
		}

		jsonStr := tag.Get("json")

		if jsonStr == "-" {
			continue
		}

		if jsonStr != "" {
			result = append(result, jsonStr)
			continue
		}

		result = append(result, field.Name)
	}

	return result, nil
}

// 接收一个map,返回map的所有val组成的切片
func ArrayValues[T ~map[K]V, K comparable, V any](arr T) []V {
	result := make([]V, 0, len(arr))

	for _, v := range arr {
		result = append(result, v)
	}

	return result
}

// 接收一个map组成的切片，返回一个以指定key作为键，指定key作为值的新map
func ArrayPluck[T ~[]M, M ~map[string]V, K string, V comparable](arr T, kName, vName string) map[V]V {
	res := make(map[V]V)

	for _, item := range arr {
		res[item[kName]] = item[vName]
	}

	return res
}

// 接收一个struct组成的切片和一个以struct为入参的func，func返回两个参数；返回一个以func返回的第一个值为key，以func返回的第二个值为value的map
func StructsPluck[T any, K comparable, V any](slice []T, kvFunc func(T) (K, V)) map[K]V {
	res := make(map[K]V)

	for _, item := range slice {
		k, v := kvFunc(item)
		res[k] = v
	}

	return res
}

// 接收一个map的切片 从map中提取指定列的值重排切片的索引，返回一个以指定列的值为键，以原map为值的新map
func ArrayReIndex[T ~[]M, M ~map[K]V, K comparable, V comparable](arr T, idx K) map[V]M {
	result := make(map[V]M)

	for _, v := range arr {
		if v1, ok := v[idx]; ok {
			result[v1] = v
		}
	}

	return result
}

// ArrayReIndex的struct版本
func StructsReIndex[T any, K comparable](slice []T, keyFunc func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range slice {
		key := keyFunc(item)
		result[key] = item
	}
	return result
}

// 切片去重
func ArrayUnique[T comparable](arr []T) []T {
	result := make([]T, 0, len(arr))

	seen := make(map[T]bool)

	for _, item := range arr {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// 遍历切片
func Each[T any, R any](arr []T, callback func(T, int) R) []R {
	result := make([]R, 0, len(arr))
	for i, item := range arr {
		result = append(result, callback(item, i))
	}
	return result
}

// 切片随机打乱
func SliceShuffle[T any](arr []T) []T {
	for i := len(arr) - 1; i > 0; i-- {
		randIndex := rand.Intn(i + 1)

		arr[i], arr[randIndex] = arr[randIndex], arr[i]
	}

	return arr
}

// 合并多个map
func MapMerge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// 两个切片取交集
func Intersection[T comparable](a, b []T) []T {
	set := make(map[T]bool)
	result := []T{}

	// 将第一个slice放入set中
	for _, item := range a {
		set[item] = true
	}

	// 检查第二个slice中的元素是否在set中
	for _, item := range b {
		if _, found := set[item]; found {
			result = append(result, item)
			delete(set, item) // 避免重复添加
		}
	}

	return result
}

// 获取两个切片的差集，切片A中存在但在切片B中不存在的元素所组成的切片
func Difference[T comparable](a, b []T) []T {
	// 创建 map 记录 b 中的元素
	setB := make(map[T]bool)
	for _, item := range b {
		setB[item] = true
	}

	result := []T{}
	added := make(map[T]bool) // 避免重复添加

	// 找出在 a 中但不在 b 中的元素
	for _, item := range a {
		if !setB[item] && !added[item] {
			result = append(result, item)
			added[item] = true
		}
	}

	return result
}

// 两个切片的对称差，返回两个切片中不重复的元素集合，即会移除两个切片中都存在的元素。
func SymmetricDifference[T comparable](a, b []T) []T {
	// A - B
	diffAB := Difference(a, b)
	// B - A
	diffBA := Difference(b, a)

	// 合并结果
	return append(diffAB, diffBA...)
}

// 多个切片的差集，支持传入多个切片，返回第一个切片中不包含在任意一个切片中的元素。
func ArrayDiff[T comparable](first []T, others ...[]T) []T {
	seen := make(map[T]bool)

	// 遍历第一个切片，记录所有元素
	for _, item := range first {
		seen[item] = true
	}

	// 遍历其他切片，移除已经存在的元素
	for _, other := range others {
		for _, item := range other {
			if seen[item] {
				delete(seen, item)
			}
		}
	}

	// 构建最终结果
	result := make([]T, 0, len(seen))
	for item := range seen {
		result = append(result, item)
	}

	return result
}
