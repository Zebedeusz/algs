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

func TestKthToLastElement(t *testing.T) {
	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "k at the beginning of the list",
			args: args{
				list: []int{1, 2, 3, 4},
				k:    1,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "k at zero index",
			args: args{
				list: []int{1, 2, 3, 4},
				k:    0,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "k at the end of the list",
			args: args{
				list: []int{1, 2, 3, 4},
				k:    4,
			},
			want: []int{4},
		},
		{
			name: "k in the middle of the list",
			args: args{
				list: []int{1, 2, 3, 4},
				k:    2,
			},
			want: []int{2, 3, 4},
		},
		{
			name: "k greater than length of the list",
			args: args{
				list: []int{1, 2, 3, 4},
				k:    5,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ds.CreateLinkedListFromSlice(tt.args.list)
			if got := KthToLastElement(&list, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KthToLastElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMiddleNode(t *testing.T) {
	type args struct {
		list []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "2nd node",
			args: args{
				list: []int{1, 2, 3},
				k:    2,
			},
			want: []int{1, 3},
		},
		{
			name: "before-last node",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6},
				k:    5,
			},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "1st node",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6},
				k:    1,
			},
			want: []int{2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ds.CreateLinkedListFromSlice(tt.args.list)
			node := ds.KthNode(&list, tt.args.k)
			DeleteMiddleNode(node)
			if got := ds.LinkedListToSlice(&list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteMiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	type args struct {
		list []int
		x    int
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
				x:    0,
			},
			want: []int{},
		},
		{
			name: "x in the beginning of the list",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				x:    1,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "x at the end of the list",
			args: args{
				list: []int{4, 5, 1, 2, 3},
				x:    3,
			},
			want: []int{1, 2, 4, 5, 3},
		},
		{
			name: "x in the middle",
			args: args{
				list: []int{6, 7, 5, 4, 3, 2},
				x:    5,
			},
			want: []int{4, 3, 2, 6, 7, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ds.CreateLinkedListFromSlice(tt.args.list)
			Partition(&list, tt.args.x)
			if got := ds.LinkedListToSlice(&list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumReversedLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "both lists empty",
			args: args{
				list1: []int{},
				list2: []int{},
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "list 1 empty",
			args: args{
				list1: []int{},
				list2: []int{1},
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "list 2 empty",
			args: args{
				list1: []int{1},
				list2: []int{},
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "one digit numbers",
			args: args{
				list1: []int{1},
				list2: []int{1},
			},
			want:    []int{2},
			wantErr: false,
		},
		{
			name: "numbers of same length",
			args: args{
				list1: []int{1, 1, 1},
				list2: []int{2, 5, 7},
			},
			want:    []int{3, 6, 8},
			wantErr: false,
		},
		{
			name: "numbers of different length",
			args: args{
				list1: []int{1, 1, 1, 7},
				list2: []int{2, 5, 7},
			},
			want:    []int{3, 6, 8, 7},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := ds.CreateLinkedListFromSlice(tt.args.list1)
			list2 := ds.CreateLinkedListFromSlice(tt.args.list2)
			got, err := SumReversedLists(&list1, &list2)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumReversedLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(ds.LinkedListToSlice(got), tt.want) {
				t.Errorf("SumReversedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumReversedListsV2(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "both lists empty",
			args: args{
				list1: []int{},
				list2: []int{},
			},
			want: []int{},
		},
		{
			name: "list 1 empty",
			args: args{
				list1: []int{},
				list2: []int{1},
			},
			want: []int{1},
		},
		{
			name: "list 2 empty",
			args: args{
				list1: []int{1},
				list2: []int{},
			},
			want: []int{1},
		},
		{
			name: "one digit numbers",
			args: args{
				list1: []int{1},
				list2: []int{1},
			},
			want: []int{2},
		},
		{
			name: "numbers of same length",
			args: args{
				list1: []int{1, 1, 1},
				list2: []int{2, 5, 7},
			},
			want: []int{3, 6, 8},
		},
		{
			name: "numbers of different length",
			args: args{
				list1: []int{1, 1, 1, 7},
				list2: []int{2, 5, 7},
			},
			want: []int{3, 6, 8, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := ds.CreateLinkedListFromSlice(tt.args.list1)
			list2 := ds.CreateLinkedListFromSlice(tt.args.list2)
			got := SumReversedListsV2(&list1, &list2)
			if !reflect.DeepEqual(ds.LinkedListToSlice(got), tt.want) {
				t.Errorf("SumReversedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
