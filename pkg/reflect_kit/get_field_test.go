//go:build !test

package reflect_kit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFieldXMethods(t *testing.T) {
	type args struct {
		obj  interface{}
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get",
			args: args{
				obj:  Person{Name: "John", Age: 30},
				name: "Name",
			},
			want:    "John",
			wantErr: assert.NoError,
		},
		{
			name: "No such field",
			args: args{
				obj:  Person{Name: "John", Age: 30},
				name: "Address",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "no such field: Address")
				return false
			},
		},
		{
			name: "Nil",
			args: args{
				obj:  nil,
				name: "Name",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "obj must not be nil")
				return false
			},
		},
		{
			name: "Not a struct",
			args: args{
				obj:  "test",
				name: "Name",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "obj must be struct")
				return false
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FieldValue(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("FieldValue(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FieldValue(%v, %v)", tt.args.obj, tt.args.name)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := FieldKind(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("FieldKind(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, reflect.TypeOf(tt.want).Kind(), got, "FieldKind(%v, %v)", tt.args.obj, tt.args.name)

		})
		t.Run(tt.name, func(t *testing.T) {
			got, err := FieldType(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("FieldType(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, reflect.TypeOf(tt.want), got, "FieldType(%v, %v)", tt.args.obj, tt.args.name)

		})
		t.Run(tt.name, func(t *testing.T) {
			got, err := FieldTypeStr(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("FieldTypeStr(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, reflect.TypeOf(tt.want).String(), got, "FieldTypeStr(%v, %v)", tt.args.obj, tt.args.name)

		})
	}
}

func TestEmbedFieldXMethods(t *testing.T) {
	to := Town{
		Int:  1,
		Str:  "Town",
		Bool: true,
		Strs: []string{"Str"},
	}
	ci := City{
		PtrTown: &to,
		Town:    to,
	}
	ct := Country{
		ID:      1,
		Name:    "Country",
		City:    ci,
		PtrCity: &ci,
	}

	type args struct {
		obj  interface{}
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "No such field",
			args: args{
				obj:  ct,
				name: "Address",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "no such field: Address")
				return false
			},
		},
		{
			name: "Nil field",
			args: args{
				obj:  &Person{},
				name: "PtrPerson.Name",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "field: PtrPerson is nil")
				return false
			},
		},

		{
			name: "Nil",
			args: args{
				obj:  nil,
				name: "Name",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "obj must not be nil")
				return false
			},
		},
		{
			name: "Not a struct",
			args: args{
				obj:  "test",
				name: "Name",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "obj must be struct")
				return false
			},
		},
		{
			name: "Not a struct",
			args: args{
				obj:  ct,
				name: "ID.Name",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "field: ID is not struct")
				return false
			},
		},
		{
			name: "Get",
			args: args{
				obj:  ct,
				name: "Name",
			},
			want:    "Country",
			wantErr: assert.NoError,
		},
		{
			name: "City.Town.Int",
			args: args{
				obj:  ct,
				name: "",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "field path must not be empty")
				return false
			},
		},
		{
			name: ".Town.Int",
			args: args{
				obj:  ct,
				name: ".Town.Int",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.EqualError(t, err, "field path:.Town.Int is invalid")
				return false
			},
		},
		{
			name: "City.Town.Int",
			args: args{
				obj:  ct,
				name: "City.Town.Int",
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "City.Town.Int",
			args: args{
				obj:  &ct,
				name: "City.Town.Int",
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "PtrCity.PtrTown.Int",
			args: args{
				obj:  &ct,
				name: "PtrCity.PtrTown.Int",
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "City.PtrTown.Int",
			args: args{
				obj:  &ct,
				name: "City.PtrTown.Int",
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "City.Town.Bool",
			args: args{
				obj:  &ct,
				name: "City.Town.Bool",
			},
			want:    true,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EmbedFieldValue(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("EmbedFieldValue(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "FieldValue(%v, %v)", tt.args.obj, tt.args.name)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := EmbedFieldKind(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("EmbedFieldKind(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, reflect.TypeOf(tt.want).Kind(), got, "FieldKind(%v, %v)", tt.args.obj, tt.args.name)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := EmbedFieldType(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("EmbedFieldType(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, reflect.TypeOf(tt.want), got, "FieldType(%v, %v)", tt.args.obj, tt.args.name)
		})

		t.Run(tt.name, func(t *testing.T) {
			got, err := EmbedFieldTypeStr(tt.args.obj, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("EmbedFieldTypeStr(%v, %v)", tt.args.obj, tt.args.name)) {
				return
			}
			assert.Equalf(t, reflect.TypeOf(tt.want).String(), got, "FieldTypeStr(%v, %v)", tt.args.obj, tt.args.name)

		})
	}
}

func TestFieldsXMethods(t *testing.T) {
	m, err := fields(nil, true, "")
	assert.EqualError(t, err, "obj must not be nil")

	m, err = fields("", true, "")
	assert.EqualError(t, err, "obj must be struct")

	m, err = selectFields(nil, nil, true, "")
	assert.EqualError(t, err, "obj must not be nil")

	m, err = selectFields("nil", nil, true, "")
	assert.EqualError(t, err, "obj must be struct")

	err = rangeFields(nil, nil, true, "")
	assert.EqualError(t, err, "obj must not be nil")

	err = rangeFields("nil", nil, true, "")
	assert.EqualError(t, err, "obj must be struct")

	type A struct {
		*A
		inner struct {
			bool
		}
		int    `json:"int"`
		String string
	}

	a := &A{}
	m, err = fields(a, true, "")
	assert.NoError(t, err)
	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)
	}
	assert.Equal(t, 5, len(m))
	assert.Equal(t, false, m["inner.bool"].CanSet())

	m, err = Fields(a)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(m))

	m, err = FieldsDeep(a)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(m))

	m, err = SelectFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return field.Tag != ""
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(m))

	m, err = SelectFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(m))

	m, err = SelectFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return field.Tag != ""
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(m))

	m, err = SelectFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 3, len(m))

	count := 0
	err = RangeFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return true
	})
	assert.NoError(t, err)
	assert.Equal(t, 4, count)

	count = 0
	err = RangeFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	count = 0
	err = RangeFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return true
	})
	assert.NoError(t, err)
	assert.Equal(t, 5, count)

	count = 0
	err = RangeFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	a = &A{
		A: &A{
			A:   nil,
			int: 0,
		},
		inner: struct {
			bool
		}{},
		int:    0,
		String: "",
	}
	m, err = fields(a, true, "")
	assert.NoError(t, err)
	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)
	}
	assert.Equal(t, 10, len(m))

	m, err = FieldsDeep(a)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(m))

	m, err = Fields(a)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(m))

	m, err = SelectFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return field.Tag != ""
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(m))

	m, err = SelectFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(m))

	m, err = SelectFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return field.Tag != ""
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(m))

	m, err = SelectFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 6, len(m))

	count = 0
	err = RangeFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return true
	})
	assert.NoError(t, err)
	assert.Equal(t, 4, count)

	count = 0
	err = RangeFields(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

	count = 0
	err = RangeFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return true
	})
	assert.NoError(t, err)
	assert.Equal(t, 10, count)

	count = 0
	err = RangeFieldsDeep(a, func(s string, field reflect.StructField, value reflect.Value) bool {
		count++
		return !field.IsExported()
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)

}
