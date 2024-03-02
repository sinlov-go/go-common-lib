package struct_kit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeepCopyByGob(t *testing.T) {
	// mock DeepCopyByGob
	type data struct {
		Foo int
		Bar string
	}

	type args struct {
		src interface{}
		dis interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes interface{}
		wantErr error
	}{
		{
			name: "dis non-pointer",
			args: args{
				src: data{
					Foo: 1,
					Bar: "sample",
				},
				dis: data{},
			},
			wantErr: fmt.Errorf("gob: attempt to decode into a non-pointer"),
		},
		{
			name: "src non-pointer",
			args: args{
				src: &data{
					Foo: 1,
					Bar: "sample",
				},
				dis: &data{},
			},
			wantRes: &data{
				Foo: 1,
				Bar: "sample",
			},
		},
		{
			name: "sample",
			args: args{
				src: data{
					Foo: 1,
					Bar: "sample",
				},
				dis: &data{},
			},
			wantRes: &data{
				Foo: 1,
				Bar: "sample",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do DeepCopyByGob
			gotErr := DeepCopyByGob(tc.args.src, tc.args.dis)

			// verify DeepCopyByGob
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			assert.Equal(t, tc.wantRes, tc.args.dis)

			srcPoint := FindStructPointer(tc.args.src)
			disPoint := FindStructPointer(tc.args.dis)
			assert.NotEqual(t, srcPoint, disPoint)
		})
	}
}
