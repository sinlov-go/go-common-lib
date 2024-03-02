package string_tools

import (
	"errors"
	"github.com/sinlov-go/go-common-lib/pkg/reflect_kit"
	"reflect"
	"strings"
)

func String2LineRaw(target string) string {
	newStr := strings.Replace(target, "\r\n", `\n`, -1)
	newStr = strings.Replace(newStr, "\n", `\n`, -1)
	newStr = strings.Replace(newStr, "\r", `\n`, -1)
	return newStr
}

const StructMemberString2LineRawTag = "string_line_2_raw"

// StructMemberString2LineRaw
// convert struct member string to line raw with tag string_tools.StructMemberString2LineRawTag
// just support as
//
//	type Foo struct {
//		Name string `string_line_2_raw:"name"`
//		Age  int
//	}
//
//	type Bar struct {
//		Foo  Foo
//		Name string `string_line_2_raw:"name"`
//		Age  int
//	}
//
//	bar := Bar{
//			Foo: Foo{
//				Name: "Foo\r\n",
//				Age:  18,
//			},
//			Name: "bob\r",
//		}
//	// change to by struct member tag
//	err := string_tools.StructMemberString2LineRaw(&bar)
//	if err != nil {
//		t.Error(err)
//	}
//
//	// some with struct
//	newBar := Bar{
//		Foo: Foo{
//			Name:  `Foo\n`,
//			Age:  18,
//		},
//		Name: `bob\n`,
//	}
func StructMemberString2LineRaw(src interface{}) error {
	if src == nil {
		return errors.New("obj must not be nil")
	}

	fields, errFieldDeep := reflect_kit.SelectFieldsDeep(src, func(s string, field reflect.StructField, value reflect.Value) bool {
		return field.Tag.Get(StructMemberString2LineRawTag) != ""
	})
	if errFieldDeep != nil {
		return errFieldDeep
	}

	for fName, field := range fields {
		if field.Kind() == reflect.String {
			newValue := String2LineRaw(field.String())
			errFiledChange := reflect_kit.SetEmbedField(src, fName, newValue)
			if errFiledChange != nil {
				return errFiledChange
			}
		}
	}
	return nil
}
