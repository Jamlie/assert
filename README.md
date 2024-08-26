# Assert

A simple asserting library in Go inspired from [ThePrimeagen](https://github.com/ThePrimagen) to help checking values in before using them to check if the logic is correct.

It can also simplify error handling.

## Example

```go
package main

import (
    "log/slog"
    "os"

    "github.com/Jamlie/assert"
)

func main() {
    val := 4

	logsFile, err := os.OpenFile("logger.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    assert.NoError(err, "Cannot open file")
    defer logsFile.Close()

	jsonHandler := slog.NewJSONHandler(logsFile, nil)
	logger := slog.New(jsonHandler)

    logger.Info("Some message",
        "user_id", val,
    )
}
```

## Functions

| Function          | Description                                                                                                 |
| ----------------- | ----------------------------------------------------------------------------------------------------------- |
| AddAssertData     | Add a key-value pair to a map to log out when the program panics                                            |
| RemoveAssertData  | Remove the key from the map                                                                                 |
| Assert            | Checks if the condition is true, if not, panics with a specified message                                    |
| NoError           | Checks if the given error is nil, if not, logs out the error and panics with a specified message            |
| Equals            | Checks if two valeus are equal, if not, panics with a specified message                                     |
| NotEquals         | Checks if two valeus are not equal, if not, panics with a specified message                                 |
| GreaterThan       | Checks if the first value is greater than the second one, if not, panics with a specified message           |
| LessThan          | Checks if the first value is less than the second one, if not, panics with a specified message              |
| GreaterThanEquals | Checks if the first value is greater than or equals the second one, if not, panics with a specified message |
| LessThanEquals    | Checks if the first value is less than or equals the second one, if not, panics with a specified message    |
| NotEmptySlice     | Checks if a slice is not empty, if not, panics with a specified message                                     |
| NotEmptyMap       | Checks if a slice is not empty, if not, panics with a specified message                                     |
