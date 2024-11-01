package goo

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
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

/*
	10秒
	1分钟
	1分10秒
	1小时
	1小时10分
	1小时10分10秒
	1小时零10秒
	1天
	1天10小时
	1天10小时10分
	1天10小时10分10秒
	1天零10分
	1天零10分10秒
	1天零10秒
*/

type durChinese struct {
	D    time.Duration
	Name []rune
}

func DurationToChinese(d time.Duration) string {
	days := d / (24 * time.Hour)
	hours := (d % (24 * time.Hour)) / time.Hour
	minutes := (d % time.Hour) / time.Minute
	seconds := (d % time.Minute) / time.Second

	sl := [4]durChinese{
		{days, []rune{'天'}},
		{hours, []rune{'小', '时'}},
		{minutes, []rune{'分'}},
		{seconds, []rune{'秒'}},
	}

	var result []rune

	for _, v := range sl {
		if v.D > 0 {
			result = append(result, []rune(fmt.Sprintf("%d", v.D))...)
			result = append(result, v.Name...)
			continue
		}

		if len(result) > 0 {
			//如果结尾已经是零,直接跳过
			if result[len(result)-1] == '零' {
				continue
			}

			result = append(result, '零')
			continue
		}
	}

	// 如果结尾是零,去掉零
	if len(result) > 0 && result[len(result)-1] == '零' {
		result = result[:len(result)-1]
	}

	// 如果结果为空，返回 "0秒"
	if len(result) == 0 {
		return "0秒"
	}

	return string(result)
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

	if !value.IsValid() {
		return false
	}

	kind := value.Kind()

	// 判断是否为数字类型
	return kind >= reflect.Int && kind <= reflect.Float64
}

func IsInteger(data any) bool {
	value := reflect.ValueOf(data)

	if !value.IsValid() {
		return false
	}

	kind := value.Kind()

	// 判断是否为数字类型
	return kind >= reflect.Int && kind <= reflect.Int64
}

func IsFloat(data any) bool {
	value := reflect.ValueOf(data)

	if !value.IsValid() {
		return false
	}

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

	if !val.IsValid() {
		return true
	}

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

	if !tVal.IsValid() || !vVal.IsValid() {
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

	//如果原类型是int, 并且目标类型是float,尝试转换
	if vVal.Kind() >= reflect.Int && vVal.Kind() <= reflect.Int64 && (tVal.Kind() == reflect.Float32 || tVal.Kind() == reflect.Float64) {
		vint64 := AnyConvert2T(v, int64(0))

		return AnyConvert2T(float64(vint64), t)
	}

	//如果是[]byte类型, 先转为string,再转换
	if vVal.Kind() == reflect.Slice && vVal.Type().Elem().Kind() == reflect.Uint8 {
		return AnyConvert2T(string(v.([]byte)), t)
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
	timer, _ := time.ParseInLocation("2006-01-02 15:04:05", t, loc)

	//尝试其他格式， RFC3339
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC3339, t, loc)
	}

	//尝试其他格式， RFC3339Nano
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC3339Nano, t, loc)
	}

	//尝试其他格式， RFC1123
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC1123, t, loc)
	}

	//尝试其他格式， RFC1123Z
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC1123Z, t, loc)
	}

	//尝试其他格式， RFC822
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC822, t, loc)
	}

	//尝试其他格式， RFC822Z
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC822Z, t, loc)
	}

	//尝试其他格式， RFC850
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC850, t, loc)
	}

	return timer.Unix()
}

func TimeString2Time(t string) time.Time {
	loc, _ := time.LoadLocation("Local") // 获取时区

	if len(t) <= 10 {
		t = fmt.Sprintf("%s 00:00:00", t)
	}
	timer, _ := time.ParseInLocation("2006-01-02 15:04:05", t, loc)

	//尝试其他格式， RFC3339
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC3339, t, loc)
	}

	//尝试其他格式， RFC3339Nano
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC3339Nano, t, loc)
	}

	//尝试其他格式， RFC1123
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC1123, t, loc)
	}

	//尝试其他格式， RFC1123Z
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC1123Z, t, loc)
	}

	//尝试其他格式， RFC822
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC822, t, loc)
	}

	//尝试其他格式， RFC822Z
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC822Z, t, loc)
	}

	//尝试其他格式， RFC850
	if timer.IsZero() {
		timer, _ = time.ParseInLocation(time.RFC850, t, loc)
	}

	return timer
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// 安全的除法， 除数为0返回0和错误
func SafeDivide[T Number](numerator, denominator T) (T, error) {
	// 检查除数是否为 0
	if denominator == 0 {
		return 0, errors.New("division by zero")
	}

	return numerator / denominator, nil
}

func JsonMarshalIndent(jsonData string) string {

	// 解析JSON数据
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return jsonData
	}

	// 将解析后的数据重新编码为带缩进的JSON字符串
	formattedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return jsonData
	}

	return string(formattedJSON)
}
