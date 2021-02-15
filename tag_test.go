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
		name  string
		args  args
		want  *Tag
		want1 bool
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
			want1: true,
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
			want1: true,
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
			want1: true,
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
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseTag(tt.args.st)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTag() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseTag() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
