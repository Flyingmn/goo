package goo

import (
	"errors"
	"math/rand"
	"reflect"
)

func ArrayColumn[T ~[]M, M ~map[K]V, K comparable, V any](arr T, k K) []V {
	result := make([]V, 0, len(arr))

	for _, v := range arr {
		if v, ok := v[k]; ok {
			result = append(result, v)
		}
	}

	return result
}

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

func ArrayKeys[T ~map[K]V, K comparable, V any](arr T) []K {
	result := make([]K, 0, len(arr))

	for k := range arr {
		result = append(result, k)
	}

	return result
}

func ArrayValues[T ~map[K]V, K comparable, V any](arr T) []V {
	result := make([]V, 0, len(arr))

	for _, v := range arr {
		result = append(result, v)
	}

	return result
}

func ArrayPluck[T ~[]M, M ~map[string]V, K string, V comparable](arr T, kName, vName string) map[V]V {
	res := make(map[V]V)

	for _, item := range arr {
		res[item[kName]] = item[vName]
	}

	return res
}

func ArrayPluckWithType[T ~[]M, M ~map[string]V, K string, V comparable, KD comparable, VD comparable](arr T, kName string, kDef KD, vName string, vDef VD) map[KD]VD {
	res := make(map[KD]VD)

	for _, item := range arr {
		k, _ := GetMapWsDef(item, kName, kDef)
		val, _ := GetMapWsDef(item, vName, vDef)

		res[k] = val
	}

	return res
}

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

func SliceShuffle[T any](arr []T) []T {
	for i := len(arr) - 1; i > 0; i-- {
		randIndex := rand.Intn(i + 1)

		arr[i], arr[randIndex] = arr[randIndex], arr[i]
	}

	return arr
}

func StructsColumn[T any, V any](structs []T, name string, defValue V) ([]V, error) {
	result := make([]V, 0, len(structs))

	for _, v := range structs {
		val := reflect.ValueOf(v)

		if !val.IsValid() {
			result = append(result, defValue)
			continue
		}

		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		if val.Kind() != reflect.Struct {
			return nil, errors.New("data is't struct")
		}

		field := val.FieldByName(name)

		if !field.IsValid() {
			return nil, errors.New("field is't valid")
		}

		value := AnyConvert2T(field.Interface(), defValue)

		result = append(result, value)
	}

	return result, nil
}
