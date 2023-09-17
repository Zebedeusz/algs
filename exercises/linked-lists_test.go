package exercises

import (
	ds "algs/data_structures"
	"reflect"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty list",
			args: args{
				list: []int{},
			},
			want: []int{},
		},
		{
			name: "no duplicates - nothing removed",
			args: args{
				list: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "duplicate from the end removed",
			args: args{
				list: []int{1, 2, 3, 4, 5, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "duplicate from the middle removed",
			args: args{
				list: []int{1, 2, 3, 3, 4, 5, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "duplicate from the beginning removed",
			args: args{
				list: []int{1, 1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "multiple duplicates removed",
			args: args{
				list: []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 5},
			},
			want: []int{1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ds.CreateLinkedListFromSlice(tt.args.list)
			RemoveDups(&list)
			if got := ds.LinkedListToSlice(&list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDups() = %v, want %v", got, tt.want)
			}
		})
	}
}
