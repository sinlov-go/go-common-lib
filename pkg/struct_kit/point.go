package struct_kit

import (
	"fmt"
	"reflect"
)

func FindStructPointer(src interface{}) string {
	if src == nil {
		return ""
	}
	value := reflect.ValueOf(src)
	if value.IsValid() && value.IsZero() {
		return ""
	}

	typeInfo := value.Type()

	if typeInfo.Kind() == reflect.Ptr {
		return fmt.Sprintf("%p", src)
	}
	if value.CanAddr() {
		return value.Addr().String()
	} else {
		return fmt.Sprintf("%p", &src)
	}
}
