// $ go get -u github.com/cweill/gotests/...

package unitTest2

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{

			name: "test Add success",
			args: args{
				x: 1,
				y: 2,
			},
			want: 3,
		},
		{

			name: "test Add success2",
			args: args{
				x: 2,
				y: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindName(t *testing.T) {
	type args struct {
		findName string
		nameList []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "TestIsFindNameSuccess",
			args: args{
				findName: "Chen",
				nameList: []string{"Chen", "Jack", "Daniel", "Sam"},
			},
			want: true,
		},
		{
			name: "TestIsFindNameSuccess2",
			args: args{
				findName: "Chen",
				nameList: []string{"Jack", "Daniel", "Sam"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindName(tt.args.findName, tt.args.nameList); got != tt.want {
				t.Errorf("FindName() = %v, want %v", got, tt.want)
			}
		})
	}
}
