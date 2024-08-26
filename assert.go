package assert

import (
	"cmp"
	"log"
	"log/slog"
	"reflect"
)

// A map for data to log out when the program panics
var assertData = make(map[string]any)

// Add a key-value pair to a map to log out when the program panics
func AddAssertData(key string, value any) {
	assertData[key] = value
}

// Remove the key from the map
func RemoveAssertData(key string) {
	delete(assertData, key)
}

// If a condition is false, logs out the data in the map, if exists, and panics with the message
func runAssert(msg string) {
	for k, v := range assertData {
		slog.Error("Context", "key", k, "value", v)
	}

	log.Panic(msg)
}

// Checks if the condition is true, if not, panics with a specified message
func Assert(truth bool, msg string) {
	if !truth {
		runAssert(msg)
	}
}

// Checks if the given error is nil, if not, logs out the error and panics with a specified message
func NoError(err error, msg string) {
	if err != nil {
		slog.Error("NoError#error encountered", "error", err)
		runAssert(msg)
	}
}

// Checks if two valeus are equal, if not, panics with a specified message
func Equals(a, b any, msg string) {
	Assert(reflect.DeepEqual(a, b), msg)
}

// Checks if two valeus are not equal, if not, panics with a specified message
func NotEquals(a, b any, msg string) {
	Assert(!reflect.DeepEqual(a, b), msg)
}

// Checks if the first value is greater than the second one, if not, panics with a specified message
func GreaterThan[T cmp.Ordered](a, b T, msg string) {
	Assert(a > b, msg)
}

// Checks if the first value is less than the second one, if not, panics with a specified message
func LessThan[T cmp.Ordered](a, b T, msg string) {
	Assert(a < b, msg)
}

// Checks if the first value is greater than or equals the second one, if not, panics with a specified message
func GreaterThanEquals[T cmp.Ordered](a, b T, msg string) {
	Assert(a >= b, msg)
}

// Checks if the first value is less than or equals the second one, if not, panics with a specified message
func LessThanEquals[T cmp.Ordered](a, b T, msg string) {
	Assert(a <= b, msg)
}

// Checks if a slice is not empty, if not, panics with a specified message
func NotEmptySlice[T ~[]E, E any](s T, msg string) {
	Assert(len(s) != 0, msg)
}

// Checks if a slice is not empty, if not, panics with a specified message
func NotEmptyMap[T ~map[K]V, K comparable, V any](m T, msg string) {
	Assert(len(m) != 0, msg)
}
