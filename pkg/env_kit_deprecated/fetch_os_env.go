package env_kit_deprecated

import (
	"os"
	"strconv"
	"strings"
)

// FetchOsEnvBool
//
//	fetch os env by key.
//	if not found will return defValue.
//	return env not same as true (will be lowercase, so TRUE is same)
//
// Deprecated: use github.com/sinlov-go/unittest-kit/env_kit instead
func FetchOsEnvBool(key string, defValue bool) bool {
	if os.Getenv(key) == "" {
		return defValue
	}
	return strings.ToLower(os.Getenv(key)) == "true"
}

// FetchOsEnvInt
//
//	fetch os env by key.
//	return not found will return devValue.
//	if not parse to int, return defValue
//
// Deprecated: use github.com/sinlov-go/unittest-kit/env_kit instead
func FetchOsEnvInt(key string, defValue int) int {
	if os.Getenv(key) == "" {
		return defValue
	}
	outNum, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defValue
	}

	return outNum
}

// FetchOsEnvStr
//
//	fetch os env by key.
//	return not found will return defValue.
//
// Deprecated: use github.com/sinlov-go/unittest-kit/env_kit instead
func FetchOsEnvStr(key, defValue string) string {
	if os.Getenv(key) == "" {
		return defValue
	}
	return os.Getenv(key)
}

// FetchOsEnvArray
//
//	fetch os env split by `,` and trim space
//	return not found will return []string(nil).
//
// Deprecated: use github.com/sinlov-go/unittest-kit/env_kit instead
func FetchOsEnvArray(key string) []string {
	var defValueStr []string
	if os.Getenv(key) == "" {
		return defValueStr
	}
	envValue := os.Getenv(key)
	splitVal := strings.Split(envValue, ",")
	for _, item := range splitVal {
		defValueStr = append(defValueStr, strings.TrimSpace(item))
	}

	return defValueStr
}
