package injector

import (
	"reflect"
	"testing"
)

func Test_parseTag(t *testing.T) {
	type args struct {
		st reflect.StructTag
	}
	tests := []struct {
		name string
		args args
		want *Tag
	}{
		{
			name: "Skip",
			args: args{
				st: reflect.StructTag(`inject:"-"`),
			},
			want: &Tag{
				Name:     "",
				Required: false,
				Skip:     true,
			},
		},
		{
			name: "Required",
			args: args{
				st: reflect.StructTag(`inject:"required"`),
			},
			want: &Tag{
				Name:     "",
				Required: true,
				Skip:     false,
			},
		},
		{
			name: "Named required",
			args: args{
				st: reflect.StructTag(`inject:"foo,required"`),
			},
			want: &Tag{
				Name:     "foo",
				Required: true,
				Skip:     false,
			},
		},
		{
			name: "Named",
			args: args{
				st: reflect.StructTag(`inject:"foo"`),
			},
			want: &Tag{
				Name:     "foo",
				Required: false,
				Skip:     false,
			},
		},
		{
			name: "No tags",
			args: args{
				st: reflect.StructTag(``),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseTag(tt.args.st)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}
