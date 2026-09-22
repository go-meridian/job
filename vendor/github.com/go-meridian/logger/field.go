package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field 日志字段类型，别名 zapcore.Field
type Field = zapcore.Field

// ObjectMarshaler 对象序列化接口
type ObjectMarshaler = zapcore.ObjectMarshaler

// ArrayMarshaler 数组序列化接口
type ArrayMarshaler = zapcore.ArrayMarshaler

// ==================== 基础类型字段 ====================

// String 构造字符串字段
func String(key string, val string) Field {
	return zap.String(key, val)
}

// Bool 构造布尔字段
func Bool(key string, val bool) Field {
	return zap.Bool(key, val)
}

// Int 构造整数字段
func Int(key string, val int) Field {
	return zap.Int(key, val)
}

// Int64 构造 int64 字段
func Int64(key string, val int64) Field {
	return zap.Int64(key, val)
}

// Int32 构造 int32 字段
func Int32(key string, val int32) Field {
	return zap.Int32(key, val)
}

// Int16 构造 int16 字段
func Int16(key string, val int16) Field {
	return zap.Int16(key, val)
}

// Int8 构造 int8 字段
func Int8(key string, val int8) Field {
	return zap.Int8(key, val)
}

// Uint 构造 uint 字段
func Uint(key string, val uint) Field {
	return zap.Uint(key, val)
}

// Uint64 构造 uint64 字段
func Uint64(key string, val uint64) Field {
	return zap.Uint64(key, val)
}

// Uint32 构造 uint32 字段
func Uint32(key string, val uint32) Field {
	return zap.Uint32(key, val)
}

// Uint16 构造 uint16 字段
func Uint16(key string, val uint16) Field {
	return zap.Uint16(key, val)
}

// Uint8 构造 uint8 字段
func Uint8(key string, val uint8) Field {
	return zap.Uint8(key, val)
}

// Float64 构造 float64 字段
func Float64(key string, val float64) Field {
	return zap.Float64(key, val)
}

// Float32 构造 float32 字段
func Float32(key string, val float32) Field {
	return zap.Float32(key, val)
}

// ==================== 错误字段 ====================

// Error 构造 error 字段（key 固定为 "error"）
func Error(err error) Field {
	return zap.Error(err)
}

// NamedError 构造指定 key 的 error 字段
func NamedError(key string, err error) Field {
	return zap.NamedError(key, err)
}

// ==================== 时间与持续时间 ====================

// Time 构造时间字段
func Time(key string, val time.Time) Field {
	return zap.Time(key, val)
}

// Duration 构造持续时间字段
func Duration(key string, val time.Duration) Field {
	return zap.Duration(key, val)
}

// ==================== 复合类型字段 ====================

// Any 根据值类型自动选择最佳字段构造方式
func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}

// Reflect 使用反射构造字段，适用于任意类型
func Reflect(key string, val interface{}) Field {
	return zap.Reflect(key, val)
}

// Binary 构造二进制数据字段
func Binary(key string, val []byte) Field {
	return zap.Binary(key, val)
}

// ByteString 构造 UTF-8 编码的字节切片字段
func ByteString(key string, val []byte) Field {
	return zap.ByteString(key, val)
}

// Stringer 构造实现了 fmt.Stringer 接口的字段
func Stringer(key string, val fmt.Stringer) Field {
	return zap.Stringer(key, val)
}

// Object 构造对象字段
func Object(key string, val ObjectMarshaler) Field {
	return zap.Object(key, val)
}

// Array 构造数组字段
func Array(key string, val ArrayMarshaler) Field {
	return zap.Array(key, val)
}

// Dict 构造嵌套字典字段
func Dict(key string, val ...Field) Field {
	return zap.Dict(key, val...)
}

// Namespace 创建命名空间
func Namespace(key string) Field {
	return zap.Namespace(key)
}

// Skip 构造空操作字段
func Skip() Field {
	return zap.Skip()
}

// Stack 构造堆栈信息字段
func Stack(key string) Field {
	return zap.Stack(key)
}

// ==================== 指针类型字段 ====================

// Stringp 构造 *string 字段
func Stringp(key string, val *string) Field {
	return zap.Stringp(key, val)
}

// Boolp 构造 *bool 字段
func Boolp(key string, val *bool) Field {
	return zap.Boolp(key, val)
}

// Intp 构造 *int 字段
func Intp(key string, val *int) Field {
	return zap.Intp(key, val)
}

// Int64p 构造 *int64 字段
func Int64p(key string, val *int64) Field {
	return zap.Int64p(key, val)
}

// Float64p 构造 *float64 字段
func Float64p(key string, val *float64) Field {
	return zap.Float64p(key, val)
}

// Durationp 构造 *time.Duration 字段
func Durationp(key string, val *time.Duration) Field {
	return zap.Durationp(key, val)
}

// Timep 构造 *time.Time 字段
func Timep(key string, val *time.Time) Field {
	return zap.Timep(key, val)
}

// ==================== 切片类型字段 ====================

// Strings 构造字符串切片字段
func Strings(key string, val []string) Field {
	return zap.Strings(key, val)
}

// Bools 构造布尔切片字段
func Bools(key string, val []bool) Field {
	return zap.Bools(key, val)
}

// Ints 构造整数切片字段
func Ints(key string, val []int) Field {
	return zap.Ints(key, val)
}

// Int64s 构造 int64 切片字段
func Int64s(key string, val []int64) Field {
	return zap.Int64s(key, val)
}

// Float64s 构造 float64 切片字段
func Float64s(key string, val []float64) Field {
	return zap.Float64s(key, val)
}

// Durations 构造持续时间切片字段
func Durations(key string, val []time.Duration) Field {
	return zap.Durations(key, val)
}

// Times 构造时间切片字段
func Times(key string, val []time.Time) Field {
	return zap.Times(key, val)
}

// Errors 构造错误切片字段
func Errors(key string, val []error) Field {
	return zap.Errors(key, val)
}

// Uints 构造 uint 切片字段
func Uints(key string, val []uint) Field {
	return zap.Uints(key, val)
}

// Uint64s 构造 uint64 切片字段
func Uint64s(key string, val []uint64) Field {
	return zap.Uint64s(key, val)
}
