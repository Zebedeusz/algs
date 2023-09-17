package datastructures

import (
	"reflect"
	"testing"
)

var testList LinkedList = LinkedList{
	First: &LinkedListNode{
		Next: &LinkedListNode{
			Next: &LinkedListNode{
				Next: &LinkedListNode{
					Next: &LinkedListNode{
						Value: 5,
					},
					Value: 4,
				},
				Value: 3,
			},
			Value: 2,
		},
		Value: 1,
	},
}

func TestCreateLinkedListFromSlice(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want LinkedList
	}{
		{
			name: "empty slice",
			args: args{
				slice: []int{},
			},
			want: LinkedList{},
		},
		{
			name: "one element",
			args: args{
				slice: []int{1},
			},
			want: LinkedList{
				First: &LinkedListNode{
					Value: 1,
				},
			},
		},
		{
			name: "test list",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
			},
			want: testList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLinkedListFromSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLinkedListFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		list *LinkedList
		elem int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "add to empty list",
			args: args{
				list: &LinkedList{},
				elem: 1,
			},
			want: []int{1},
		},
		{
			name: "add to test list",
			args: args{
				list: &testList,
				elem: 1,
			},
			want: []int{1, 2, 3, 4, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.list, tt.args.elem)
			if got := LinkedListToSlice(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedListToSlice(t *testing.T) {
	type args struct {
		list *LinkedList
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty list",
			args: args{
				list: &LinkedList{},
			},
			want: []int{},
		},
		{
			name: "test list",
			args: args{
				list: &testList,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkedListToSlice(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedListToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	type args struct {
		list *LinkedList
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty list",
			args: args{
				list: &LinkedList{},
			},
			want: 0,
		},
		{
			name: "test list",
			args: args{
				list: &testList,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Len(tt.args.list); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		list *LinkedList
		elem int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "remove from an empty list",
			args: args{
				list: &LinkedList{},
				elem: 1,
			},
			want: []int{},
		},
		{
			name: "remove first element from test list",
			args: args{
				list: &testList,
				elem: 1,
			},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "remove middle element from test list",
			args: args{
				list: &testList,
				elem: 3,
			},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "remove last element from test list",
			args: args{
				list: &testList,
				elem: 5,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listCopy := CreateLinkedListFromSlice(LinkedListToSlice(tt.args.list))
			Remove(&listCopy, tt.args.elem)
			if got := LinkedListToSlice(&listCopy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
