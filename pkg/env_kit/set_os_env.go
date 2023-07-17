package env_kit

import (
	"fmt"
	"os"
	"strconv"
)

// SetEnvStr
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvStr(key string, val string) error {
	err := os.Setenv(key, val)
	if err != nil {
		return fmt.Errorf("set env key [%v] string err: %v", key, err)
	}
	return nil
}

// SetEnvBool
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvBool(key string, val bool) error {
	var err error
	if val {
		err = os.Setenv(key, "true")
	} else {
		err = os.Setenv(key, "false")
	}
	if err != nil {
		return fmt.Errorf("set env key [%v] bool err: %v", key, err)
	}
	return nil
}

// SetEnvU64
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvU64(key string, val uint64) error {
	err := os.Setenv(key, strconv.FormatUint(val, 10))
	if err != nil {
		return fmt.Errorf("set env key [%v] uint64 err: %v", key, err)
	}
	return nil
}

// SetEnvInt64
//
//	set env by key and val
//
//nolint:golint,unused
func SetEnvInt64(key string, val int64) error {
	err := os.Setenv(key, strconv.FormatInt(val, 10))
	if err != nil {
		return fmt.Errorf("set env key [%v] int64 err: %v", key, err)
	}
	return nil
}
