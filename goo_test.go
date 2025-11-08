package goo_test

import (
	"errors"
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

	var vint int = 1
	fmt.Println("vint:", goo.AnyConvert2T(&vint, int(0)))

	type user struct {
		Name string
	}
	u := &user{
		Name: "张三",
	}

	var ua any = u
	fmt.Println("user:", goo.AnyConvert2T(ua, &user{}))
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

	fmt.Println(goo.ArrayDiff(
		[]int{1, 2},
		[]int{2, 3, 4, 5, 6, 7, 8, 9, 10},
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
	data := goo.ArrayPluck([]map[string]int{
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
	}, "a", "b")

	fmt.Println(data)
}

func Test_StructsPluck(t *testing.T) {
	type User struct {
		Name string
		Age  int
		Id   int
	}
	users := []*User{
		{
			Name: "张三",
			Age:  20,
			Id:   1,
		},
		{
			Name: "李四",
			Age:  19,
			Id:   2,
		},
		{
			Name: "王五",
			Age:  18,
			Id:   3,
		},
	}
	fmt.Println(goo.StructsPluck(users, func(user *User) (int, string) {
		return user.Id, user.Name
	}))

	fmt.Println(goo.StructsPluck(users, func(user *User) (int, *User) {
		return user.Id, user
	}))
}

func Test_ArrayReIndex(t *testing.T) {
	fmt.Println(goo.ArrayReIndex([]map[string]int{
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
	}, "a"))
}

func Test_StructsReIndex(t *testing.T) {
	type User struct {
		Name string
		Age  int
		Id   int
	}
	users := []User{
		{
			Name: "张三",
			Age:  20,
			Id:   1,
		},
		{
			Name: "李四",
			Age:  19,
			Id:   2,
		},
		{
			Name: "王五",
			Age:  18,
			Id:   3,
		},
	}
	fmt.Println(goo.StructsReIndex(users, func(user User) int {
		return user.Age
	}))
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
	fmt.Println("nil", goo.Empty(nil))
	fmt.Println("0", goo.Empty(0))
	fmt.Println("0.0", goo.Empty(0.0))
	fmt.Println("''", goo.Empty(""))
	fmt.Println("false", goo.Empty(false))
	fmt.Println("[]int{}", goo.Empty([]int{}))
	fmt.Println("map[int]int{}", goo.Empty(map[int]int{}))
	fmt.Println("struct{}{}", goo.Empty(struct{}{}))

	var m = map[string]any{
		"test": 1,
	}
	//IsValid
	fmt.Println("m[\"not exists\"]", goo.Empty(m["not exists"]))

	var i *int

	fmt.Println("reflect.ValueOf(i)", goo.Empty(reflect.ValueOf(i)))

	var s *struct{}
	fmt.Println("s", goo.Empty(s))
	fmt.Println("reflect.ValueOf(s)", goo.Empty(reflect.ValueOf(s)))
}

// func GetMapWsDef[C comparable, V any, DV any](m map[C]V, key C, def DV) (DV, bool)
func Test_GetMapWsDef(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	v, ok := goo.GetMapWsDef(m, 1, "")
	fmt.Println(v, ok)
	v, ok = goo.GetMapWsDef(m, 3, "")
	fmt.Println(v, ok)
}

func Test_GetMapWsDefWsOutOk(t *testing.T) {
	m := map[int]string{1: "a", 2: "b"}
	v := goo.GetMapWsDefWsOutOk(m, 1, "")
	fmt.Println(v)
	v = goo.GetMapWsDefWsOutOk(m, 3, "")
	fmt.Println(v)
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
	fmt.Println(goo.IsStruct(nil))
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

	mFuncRes := goo.MarshalJson(map[string]func(){
		"a": func() {
			fmt.Println("a")
		},
	})

	fmt.Println("\nmFuncRes:", mFuncRes)
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
	names := goo.StructsColumn(users, func(u User) string {
		return u.Name
	})

	fmt.Println(names)

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

func Test_JsonMarshalIndent(t *testing.T) {
	fmt.Println(goo.JsonMarshalIndent(`{"a":1,"b":2}`))
	fmt.Println(goo.JsonMarshalIndent(`{"a":1,"b":2`))
	fmt.Println(goo.JsonMarshalIndent(`{"a":1,"b:2}`))
	fmt.Println(goo.JsonMarshalIndent(`666`))
}

func Test_StructKeys(t *testing.T) {
	type User struct {
		Name      string    `json:"name" gorm:"column:user_name"`
		Age       int       `json:"age" gorm:"column:user_age"`
		Class     string    `json:"class" gorm:"column:"`
		CreatedAt time.Time `json:"created_at" gorm:"-"`
		UpdatedAt time.Time `json:"-" gorm:"column:updated_at"`
		DeletedAt time.Time `json:"-" gorm:"default:null"`
		IsDeleted bool      `json:"is_deleted" gorm:""`
		IsActive  bool      ``
	}
	var user = User{Name: "张三", Age: 18, Class: "1班", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Now(), IsDeleted: true}
	var user2 User
	fmt.Println(goo.StructKeys(user))

	fmt.Println(goo.StructKeys(&user))

	fmt.Println(goo.StructKeys(&User{}))

	fmt.Println(goo.StructKeys(User{}))

	fmt.Println(goo.StructKeys(user2))

	fmt.Println(goo.StructKeys(1))

	fmt.Println(goo.StructKeys(false))

	fmt.Println(goo.StructKeys("notExistsTag"))

}

func Test_ParseGormColumnTag(t *testing.T) {
	fmt.Println(goo.ParseGormColumnTag("column"))
	fmt.Println(goo.ParseGormColumnTag(""))
}

func Test_ConcurrentWithLimit(t *testing.T) {
	ret := goo.ConcurrentWithLimit([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, func(item int) error {
		fmt.Println(item)
		time.Sleep(time.Second * time.Duration(goo.RandomIntInRange(1, 5)))
		return nil
	})

	fmt.Println(ret)
}

func Test_ConcurrentWithLimitRetErrs(t *testing.T) {
	ret2, err2 := goo.ConcurrentWithLimitRetErrs([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, func(item int) (int, error) {
		//随机返回error
		if goo.RandomIntInRange(1, 10)%2 == 0 {
			return item, errors.New("ConcurrentWithLimitRetErrs test fail")
		}

		fmt.Println(item)
		time.Sleep(time.Second * time.Duration(goo.RandomIntInRange(1, 5)))
		return item, nil
	})

	fmt.Println(ret2, err2)

}

func Test_ChunkExec(t *testing.T) {
	ret, err := goo.ChunkExec([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, func(miniVals []int) ([]int, error) {
		//随机返回error
		if goo.RandomIntInRange(1, 10)%2 == 0 {
			return miniVals, errors.New("ChunkExec test fail")
		}

		fmt.Println(miniVals)
		return miniVals, nil
	})

	fmt.Println(ret, err)
}

func Test_Each(t *testing.T) {
	datas := goo.Each([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(item int, i int) string {
		return fmt.Sprintf("_%d>", item)
	})

	fmt.Println(datas)
}

func Test_Intersection(t *testing.T) {
	fmt.Println(goo.Intersection([]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 6, 7}))
}

func Test_Difference(t *testing.T) {
	fmt.Println(goo.Difference([]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 6, 7}))
}

func Test_SymmetricDifference(t *testing.T) {
	fmt.Println(goo.SymmetricDifference([]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 6, 7}))
}

func Test_ErrJoin(t *testing.T) {
	fmt.Println(goo.ErrJoin(errors.New("1"), nil, errors.New("2")))
	fmt.Println(goo.ErrJoin())
}
