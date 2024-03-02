package string_tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Str2LineRaw(t *testing.T) {
	// mock _Str2LineRaw
	mockEnvDroneCommitMessage := "mock message commit\nmore line\nand more line\r\n"
	commitMessage := mockEnvDroneCommitMessage
	t.Logf("~> mock _Str2LineRaw")
	// do _Str2LineRaw
	t.Logf("~> do _Str2LineRaw")
	lineRaw := String2LineRaw(commitMessage)
	//commitMessage = strings.Replace(commitMessage, "\n", `\\n`, -1)
	t.Logf("lineRaw: %v", lineRaw)
	assert.Equal(t, "mock message commit\\nmore line\\nand more line\\n", lineRaw)
	// verify _Str2LineRaw
}

func TestStructMemberString2LineRaw(t *testing.T) {
	// mock StructMemberString2LineRaw
	type Foo struct {
		Name string `string_line_2_raw:"name"`
		Age  int
	}

	type Bar struct {
		Foo  Foo
		Name string `string_line_2_raw:"name"`
		Age  int
	}

	type args struct {
		src interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult interface{}
		wantErr    error
	}{
		{
			name: "sample",
			args: args{
				src: &Foo{
					Name: "Foo\r\n",
					Age:  18,
				},
			},
			wantResult: &Foo{
				Name: `Foo\n`,
				Age:  18,
			},
		},
		{
			name: "in struct",
			args: args{
				src: &Bar{
					Foo: Foo{
						Name: "Foo\r\n",
						Age:  18,
					},
					Name: "bob\r",
				},
			},
			wantResult: &Bar{
				Foo: Foo{
					Name: `Foo\n`,
					Age:  18,
				},
				Name: `bob\n`,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do StructMemberString2LineRaw
			gotErr := StructMemberString2LineRaw(tc.args.src)

			// verify StructMemberString2LineRaw
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			assert.Equal(t, tc.wantResult, tc.args.src)
		})
	}
}
