package struct_kit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindStructPointer(t *testing.T) {
	// mock FindStructPointer
	type data struct {
		Name string
		Foo  int
	}
	type args struct {
		//
		src interface{}
	}
	tests := []struct {
		name                 string
		args                 args
		wantRes              string
		wantNotPointNotEmpty bool
	}{
		{
			name: "nil",
			args: args{
				src: nil,
			},
		},
		{
			name: "foo",
			args: args{
				src: data{
					Name: "foo",
				},
			},
			wantNotPointNotEmpty: true,
		},
		{
			name: "bar",
			args: args{
				src: &data{
					Name: "bar",
				},
			},
			wantNotPointNotEmpty: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do FindStructPointer
			gotResult := FindStructPointer(tc.args.src)

			t.Logf("try get pointer: %v", fmt.Sprintf("%p", tc.args.src))
			if tc.wantNotPointNotEmpty {
				t.Logf("FindStructPointer: %v", gotResult)
				return
			}
			// verify FindStructPointer
			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}
