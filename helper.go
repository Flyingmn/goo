package goo

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func Md5(input string) string {
	hasher := md5.New()         // 创建一个md5 Hash对象
	hasher.Write([]byte(input)) // 将字符串写入hasher
	digest := hasher.Sum(nil)   // 计算哈希并获取16字节的结果

	// 将16字节的哈希转换为16进制字符串
	MD5String := fmt.Sprintf("%x", digest)

	return MD5String
}

func ArrayColumn[T ~[]M, M ~map[K]V, K comparable, V any](arr T, k K) []V {
	result := make([]V, 0, len(arr))

	for _, v := range arr {
		if v, ok := v[k]; ok {
			result = append(result, v)
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

// 判断变量是否为0，只有数字类型才可能返回true
func IsNumZero(v any) bool {
	switch v := v.(type) {
	case int, int8, int16, int32, int64:
		return v == 0
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return v == 0
	case float32, float64:
		return v == 0.0
	default:
		return false
	}
}

// isNumeric 使用 reflect 判断给定的 interface{} 类型是否为数字类型
func IsNumeric(data any) bool {
	value := reflect.ValueOf(data)
	kind := value.Kind()

	// 判断是否为数字类型
	return kind >= reflect.Int && kind <= reflect.Float64
}

func IsInteger(data any) bool {
	value := reflect.ValueOf(data)
	kind := value.Kind()

	// 判断是否为数字类型
	return kind >= reflect.Int && kind <= reflect.Int64
}

func IsFloat(data any) bool {
	value := reflect.ValueOf(data)
	kind := value.Kind()

	// 判断是否为数字类型
	return kind == reflect.Float32 || kind == reflect.Float64
}

// 判断变量是否为空
func Empty(v any) bool {
	if v == nil {
		return true
	}

	val := reflect.ValueOf(v)

	return val.IsZero()
}

// 判断map类型的key是否存在
func IsSet[C comparable, V any](m map[C]V, key C) bool {
	_, ok := m[key]

	return ok
}

// 判断map类型的key是否存在, 不存在时返回
func GetMapWsDef[C comparable, V any, DV any](m map[C]V, key C, def DV) (DV, bool) {
	v, ok := m[key]

	if !ok {
		return def, ok
	}

	return AnyConvert2T(v, def), ok
}

/*
断言 any 类型是否能转换为指定类型，如果是，返回断言后的结果，否则返回指定的值
*/
func AnyConvert2T[T any](v any, t T) T {
	vVal := reflect.ValueOf(v)
	tVal := reflect.ValueOf(t)

	//检查zero val
	if vVal.IsZero() || tVal.IsZero() {
		return t
	}

	//如果原类型string, 并且目标类型是int,尝试转换
	if vVal.Kind() == reflect.String && tVal.Kind() >= reflect.Int && tVal.Kind() <= reflect.Int64 {
		if num, err := strconv.ParseInt(v.(string), 10, 64); err == nil {
			return AnyConvert2T(num, t)
		}
	}

	//如果原类型是int, 并且目标类型是string,尝试转换
	if vVal.Kind() >= reflect.Int && vVal.Kind() <= reflect.Int64 && tVal.Kind() == reflect.String {
		vint64 := AnyConvert2T(v, int64(0))

		return AnyConvert2T(fmt.Sprintf("%d", vint64), t)
	}

	//如果原类型是float, 并且目标类型是string,尝试转换
	if (vVal.Kind() == reflect.Float32 || vVal.Kind() == reflect.Float64) && tVal.Kind() == reflect.String {
		vfloat64 := AnyConvert2T(v, float64(0))

		return AnyConvert2T(fmt.Sprintf("%f", vfloat64), t)
	}

	//如果原类型是string, 并且目标类型是float,尝试转换
	if vVal.Kind() == reflect.String && (tVal.Kind() == reflect.Float32 || tVal.Kind() == reflect.Float64) {
		if num, err := strconv.ParseFloat(v.(string), 64); err == nil {
			return AnyConvert2T(num, t)
		}
	}

	//其他情况
	if vVal.Type().ConvertibleTo(tVal.Type()) {
		return vVal.Convert(tVal.Type()).Interface().(T)
	}

	return t
}

func MarshalJson(v any) string {
	jon, err := json.Marshal(v)

	if err != nil {
		fmt.Printf("utilities.MarshalJson param %v error %v", v, err)
	}

	return string(jon)
}

func IsMap(data any) bool {
	return reflect.TypeOf(data).Kind() == reflect.Map
}

func IsStruct(data any) bool {
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Ptr:
		// 如果是指针类型，则检查它指向的对象是否为 struct
		return val.Elem().Kind() == reflect.Struct
	case reflect.Struct:
		// 如果本身就是 struct 类型
		return true
	default:
		// 其他情况返回 false
		return false
	}
}

// YYYY-MM-DD -> unix
// YYYY-MM-DD hh:mm:ss -> unix
func TimeString2Unix(t string) int64 {
	loc, _ := time.LoadLocation("Local") // 获取时区

	if len(t) <= 10 {
		t = fmt.Sprintf("%s 00:00:00", t)
	}
	timer, _ := time.ParseInLocation(time.DateTime, t, loc)

	return timer.Unix()
}

func TimeString2Time(t string) time.Time {
	loc, _ := time.LoadLocation("Local") // 获取时区

	if len(t) <= 10 {
		t = fmt.Sprintf("%s 00:00:00", t)
	}
	timer, _ := time.ParseInLocation(time.DateTime, t, loc)

	return timer
}
