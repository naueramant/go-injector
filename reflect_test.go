package injector

import "testing"

func Test_isStructPtr(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Nil is not a struct pointer",
			args: args{
				val: nil,
			},
			want: false,
		},
		{
			name: "String is not a struct pointer",
			args: args{
				val: "Hello World",
			},
			want: false,
		},
		{
			name: "Int is not a struct pointer",
			args: args{
				val: 42,
			},
			want: false,
		},
		{
			name: "Valid struct pointer",
			args: args{
				val: &struct{}{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStructPtr(tt.args.val); got != tt.want {
				t.Errorf("isStructPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
