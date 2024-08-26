package assert

import (
	"cmp"
	"log"
	"log/slog"
	"reflect"
)

var assertData = make(map[string]any)

func AddAssertData(key string, value any) {
	assertData[key] = value
}

func RemoveAssertData(key string) {
	delete(assertData, key)
}

func runAssert(msg string) {
	for k, v := range assertData {
		slog.Error("Context", "key", k, "value", v)
	}

	log.Panic(msg)
}

func Assert(truth bool, msg string) {
	if !truth {
		runAssert(msg)
	}
}

func NoError(err error, msg string) {
	if err != nil {
		slog.Error("NoError#error encountered", "error", err)
		runAssert(msg)
	}
}

func Equals(a, b any, msg string) {
	Assert(reflect.DeepEqual(a, b), msg)
}

func NotEquals(a, b any, msg string) {
	Assert(!reflect.DeepEqual(a, b), msg)
}

func GreaterThan[T cmp.Ordered](a, b T, msg string) {
	Assert(a > b, msg)
}

func LessThan[T cmp.Ordered](a, b T, msg string) {
	Assert(a < b, msg)
}

func GreaterThanEquals[T cmp.Ordered](a, b T, msg string) {
	Assert(a >= b, msg)
}

func LessThanEquals[T cmp.Ordered](a, b T, msg string) {
	Assert(a <= b, msg)
}

func NotEmptySlice[T ~[]E, E any](s T, msg string) {
	Assert(len(s) != 0, msg)
}

func NotEmptyMap[T ~map[K]V, K comparable, V any](m T, msg string) {
	Assert(len(m) != 0, msg)
}
