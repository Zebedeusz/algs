package datastructures

import (
	"reflect"
	"testing"
)

func TestStackLinkedList_Push(t *testing.T) {
	type fields struct {
		stackList []int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "adding to empty stack",
			fields: fields{
				stackList: []int{},
			},
			args: args{
				v: 1,
			},
			want: []int{1},
		},
		{
			name: "adding to stack",
			fields: fields{
				stackList: []int{1, 2, 3, 4, 5},
			},
			args: args{
				v: 7,
			},
			want: []int{7, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := CreateLinkedListFromSlice(tt.fields.stackList)
			s := &StackLinkedList{
				first: list.First,
			}
			s.Push(tt.args.v)
			got := LinkedListToSlice(&LinkedList{
				First: s.first,
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StackLinkedList.Push() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestStackLinkedList_Pop(t *testing.T) {
	type fields struct {
		stackList []int
	}
	tests := []struct {
		name               string
		fields             fields
		want               int
		wantRemainingStack []int
	}{
		{
			name: "pop from empty stack",
			fields: fields{
				stackList: []int{},
			},
			want:               -1,
			wantRemainingStack: []int{},
		},
		{
			name: "pop from stack",
			fields: fields{
				stackList: []int{1, 2, 3, 4, 5},
			},
			want:               1,
			wantRemainingStack: []int{2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := CreateLinkedListFromSlice(tt.fields.stackList)
			s := &StackLinkedList{
				first: list.First,
			}
			if got := s.Pop(); got != tt.want {
				t.Errorf("StackLinkedList.Pop() = %v, want %v", got, tt.want)
			}
			remainingStack := LinkedListToSlice(&LinkedList{
				First: s.first,
			})
			if !reflect.DeepEqual(remainingStack, tt.wantRemainingStack) {
				t.Errorf("remainingStack = %v, wantRemainingStack = %v", remainingStack, tt.wantRemainingStack)
			}
		})
	}
}
