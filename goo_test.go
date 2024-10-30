package goo_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Flyingmn/goo"
)

// func AnyConvert2T[T any](v any, t T) T
func TestAnyConvert2T(t *testing.T) {
	fmt.Println(goo.AnyConvert2T(nil, ""))
	fmt.Println(goo.AnyConvert2T(1, ""))
	fmt.Println(goo.AnyConvert2T(1, 0))
	fmt.Println(goo.AnyConvert2T(1, 0.0))
	fmt.Println(goo.AnyConvert2T(1, true))
	fmt.Println(goo.AnyConvert2T(1, []int{}))
	fmt.Println(goo.AnyConvert2T(1, map[string]int{}))

	fmt.Println(goo.AnyConvert2T("1", int(0)))
	fmt.Println(goo.AnyConvert2T("1", float64(0)))
	fmt.Println(goo.AnyConvert2T(float64(0), ""))
	fmt.Println(goo.AnyConvert2T([]byte{'1'}, float64(0)))
}

// func ArrayChunk[T ~[]V, V any](s T, size int) []T
// func ArrayColumn[T ~[]M, M ~map[K]V, K comparable, V any](arr T, k K) []V
// func ArrayDiff[T comparable](first []T, others ...[]T) []T
// func ArrayKeys[T ~map[K]V, K comparable, V any](arr T) []K
// func ArrayPluck[T ~[]M, M ~map[string]V, K string, V comparable](arr T, kName, vName string) map[V]V

func Test_ArrayChunk(t *testing.T) {
	fmt.Println(goo.ArrayChunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
	fmt.Println(goo.ArrayChunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 100))
}

func Test_ArrayColumn(t *testing.T) {
	fmt.Println(goo.ArrayColumn([]map[string]int{
		{"a": 1},
		{"a": 2},
		{"a": 3},
	}, "a"))
}

func Test_ArrayDiff(t *testing.T) {
	fmt.Println(goo.ArrayDiff(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{1, 2},
	))
}

func Test_ArrayKeys(t *testing.T) {
	fmt.Println(goo.ArrayKeys(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}))
}

func Test_ArrayPluck(t *testing.T) {
	fmt.Println(goo.ArrayPluck([]map[string]int{
		{
			"a": 1,
			"b": 2,
			"c": 3,
		},
		{
			"a": 4,
			"b": 5,
			"c": 6,
		},
		{
			"a": 7,
			"b": 8,
			"c": 9,
		},
	}, "a", "b"))
}

// func ArrayPluckWithType[T ~[]M, M ~map[string]V, K string, V comparable, KD comparable, VD comparable](arr T, kName string, kDef KD, vName string, vDef VD) map[KD]VD
func Test_ArrayPluckWithType(t *testing.T) {
	fmt.Println(goo.ArrayPluckWithType([]map[string]int{
		{
			"a": 1,
			"b": 2,
			"c": 3,
		},
		{
			"a": 4,
			"b": 5,
			"c": 6,
		},
		{
			"a": 7,
			"b": 8,
			"c": 9,
		},
	}, "a", 0, "b", 0))
}

// func ArrayUnique[T comparable](arr []T) []T
func Test_ArrayUnique(t *testing.T) {
	fmt.Println(goo.ArrayUnique([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}))
}

// func ArrayValues[T ~map[K]V, K comparable, V any](arr T) []V
func Test_ArrayValues(t *testing.T) {
	fmt.Println(goo.ArrayValues(map[int]string{
		1:  "a",
		2:  "b",
		3:  "c",
		4:  "d",
		5:  "e",
		6:  "f",
		7:  "g",
		8:  "h",
		9:  "i",
		10: "j",
		11: "k",
	}))
}

// func DurationToChinese(d time.Duration) string
func Test_DurationToChinese(t *testing.T) {
	fmt.Println(goo.DurationToChinese(0))
	fmt.Println(goo.DurationToChinese(time.Second))
	fmt.Println(goo.DurationToChinese(time.Second * 2))
	fmt.Println(goo.DurationToChinese(time.Second * 60))
	fmt.Println(goo.DurationToChinese(time.Second * 61))
	fmt.Println(goo.DurationToChinese(time.Second * 60 * 60))
	fmt.Println(goo.DurationToChinese(time.Second*3600*24 + 60))
}

// func Empty(v any) bool

func Test_Empty(t *testing.T) {
	fmt.Println(goo.Empty(nil))
	fmt.Println(goo.Empty(0))
	fmt.Println(goo.Empty(0.0))
	fmt.Println(goo.Empty(""))
	fmt.Println(goo.Empty(false))
	fmt.Println(goo.Empty([]int{}))
	fmt.Println(goo.Empty(map[int]int{}))
	fmt.Println(goo.Empty(struct{}{}))

	var i *int

	fmt.Println(goo.Empty(reflect.ValueOf(i)))

	var s *struct{}
	fmt.Println(goo.Empty(reflect.ValueOf(s)))
}

// func GetMapWsDef[C comparable, V any, DV any](m map[C]V, key C, def DV) (DV, bool)
func Test_GetMapWsDef(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	v, ok := goo.GetMapWsDef(m, 1, "")
	fmt.Println(v, ok)
	v, ok = goo.GetMapWsDef(m, 3, "")
	fmt.Println(v, ok)
}

// func IsFloat(data any) bool
func Test_IsFloat(t *testing.T) {
	fmt.Println(goo.IsFloat(1))
	fmt.Println(goo.IsFloat(1.1))
	fmt.Println(goo.IsFloat("1.1"))
	fmt.Println(goo.IsFloat(true))
	fmt.Println(goo.IsFloat(nil))
}

// func IsInteger(data any) bool
func Test_IsInteger(t *testing.T) {
	fmt.Println(goo.IsInteger(1))
	fmt.Println(goo.IsInteger(1.1))
	fmt.Println(goo.IsInteger("1.1"))
	fmt.Println(goo.IsInteger(nil))
}

// func IsMap(data any) bool
func Test_IsMap(t *testing.T) {
	fmt.Println(goo.IsMap(map[string]int{}))
	fmt.Println(goo.IsMap(map[string]int{}))
}

// func IsNumZero(v any) bool
func Test_IsNumZero(t *testing.T) {
	fmt.Println(goo.IsNumZero(0))
	fmt.Println(goo.IsNumZero(uint(0)))
	fmt.Println(goo.IsNumZero(0.0))
	fmt.Println(goo.IsNumZero("0"))
	fmt.Println(goo.IsNumZero("0.0"))
}

// func IsNumeric(data any) bool
func Test_IsNumeric(t *testing.T) {
	fmt.Println(goo.IsNumeric(0))
	fmt.Println(goo.IsNumeric(nil))
	fmt.Println(goo.IsNumeric(0.0))
	fmt.Println(goo.IsNumeric("0"))
	fmt.Println(goo.IsNumeric("0.0"))
}

// func IsSet[C comparable, V any](m map[C]V, key C) bool
func Test_IsSet(t *testing.T) {
	m := map[int]string{1: "1", 2: "2"}
	fmt.Println(goo.IsSet(m, 1))
	fmt.Println(goo.IsSet(m, 3))
}

// func IsStruct(data any) bool
func Test_IsStruct(t *testing.T) {
	fmt.Println(goo.IsStruct(map[string]string{}))
	fmt.Println(goo.IsStruct(struct{}{}))
	fmt.Println(goo.IsStruct(&struct{}{}))
}

// func MapMerge[K comparable, V any](maps ...map[K]V) map[K]V
func Test_MapMerge(t *testing.T) {
	m1 := map[string]string{"a": "1", "b": "2"}
	m2 := map[string]string{"c": "3", "d": "4"}
	m3 := map[string]string{"e": "5", "f": "6"}
	fmt.Println(goo.MapMerge(m1, m2, m3))
}

// func MarshalJson(v any) string
func Test_MarshalJson(t *testing.T) {
	fmt.Println(goo.MarshalJson(map[string]string{"a": "1", "b": "2"}))
	fmt.Println(goo.MarshalJson("111"))
	fmt.Println(goo.MarshalJson(nil))

	fmt.Println(goo.MarshalJson(map[string]func(){
		"a": func() {
			fmt.Println("a")
		},
	}))
}

// func Md5(input string) string
func Test_Md5(t *testing.T) {
	fmt.Println(goo.Md5("123456"))
}

// func SafeDivide[T Number](numerator, denominator T) (T, error)
func Test_SafeDivide(t *testing.T) {
	fmt.Println(goo.SafeDivide(1, 2))
	fmt.Println(goo.SafeDivide(1, 0))
}

// func SliceShuffle[T any](arr []T) []T
func Test_SliceShuffle(t *testing.T) {
	fmt.Println(goo.SliceShuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

// func StructsColumn[T any, V any](structs []T, name string, defValue V) ([]V, error)
func Test_StructsColumn(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	var users = []User{
		{Name: "张三", Age: 18},
		{Name: "李四", Age: 19},
		{Name: "王五", Age: 20},
	}
	var names, err = goo.StructsColumn(users, "Name", "")

	fmt.Println(names, err)

	var i *int
	var s = ""
	var sb = []byte{}

	var st struct{}
	var stp *struct{}

	fmt.Println(goo.StructsColumn([]any{i, s, sb, st, stp, reflect.ValueOf(i), reflect.ValueOf(i).Elem()}, "Name", ""))
}

// func TimeString2Time(t string) time.Time
func Test_TimeString2Unix(t *testing.T) {
	fmt.Println(goo.TimeString2Unix("2022-01-01 00:00:00"))
	fmt.Println(goo.TimeString2Unix("2022-01-01"))
	fmt.Println(goo.TimeString2Unix("2022-01-01 00:00:00.000000"))

	// RFC3339
	fmt.Println(goo.TimeString2Unix("2022-01-01T12:00:00Z"))
	// RFC3339Nano
	fmt.Println(goo.TimeString2Unix("2022-01-01T12:00:00.000000Z"))

	//RFC1123
	fmt.Println(goo.TimeString2Unix("Mon, 01 Jan 2022 12:00:00 MST"))
	//RFC1123Z
	fmt.Println(goo.TimeString2Unix("Mon, 01 Jan 2022 12:00:00 -0000"))
	//RFC822  02 Jan 06 15:04 MST
	fmt.Println(goo.TimeString2Unix("01 Jan 22 12:00 MST"))
	//RFC822Z
	fmt.Println(goo.TimeString2Unix("01 Jan 22 12:00 -0700"))
	//RFC850 Monday, 02-Jan-06 15:04:05 MST
	fmt.Println(goo.TimeString2Unix("Monday, 01-Jan-22 12:00:01 MST"))
}

// func TimeString2Unix(t string) int64
func Test_TimeString2Time(t *testing.T) {
	fmt.Println(goo.TimeString2Time("2022-01-01").Format("2006-01-02 15:04:05"))
	fmt.Println(goo.TimeString2Time("2022-01-01 12:00:00").Format("2006-01-02 15:04:05"))

	fmt.Println(goo.TimeString2Time("2022-01-01 12:00:00.000000").Format("2006-01-02 15:04:05"))
	// RFC3339
	fmt.Println(goo.TimeString2Time("2022-01-01T12:00:00Z").Format("2006-01-02 15:04:05"))
	// RFC3339Nano
	fmt.Println(goo.TimeString2Time("2022-01-01T12:00:00.000000Z").Format("2006-01-02 15:04:05"))

	//RFC1123 Mon, 02 Jan 2006 15:04:05 MST
	fmt.Println(goo.TimeString2Time("Mon, 01 Jan 2022 12:00:00 MST").Format("2006-01-02 15:04:05"))
	//RFC1123Z
	fmt.Println(goo.TimeString2Time("Mon, 01 Jan 2022 12:00:00 -0700").Format("2006-01-02 15:04:05"))
	//RFC822  02 Jan 06 15:04 MST
	fmt.Println(goo.TimeString2Time("01 Jan 22 12:00 MST").Format("2006-01-02 15:04:05"))
	//RFC822Z
	fmt.Println(goo.TimeString2Time("01 Jan 22 12:00 -0700").Format("2006-01-02 15:04:05"))
	//RFC850 Monday, 02-Jan-06 15:04:05 MST
	fmt.Println(goo.TimeString2Time("Monday, 01-Jan-22 12:00:01 MST").Format("2006-01-02 15:04:05"))
}
