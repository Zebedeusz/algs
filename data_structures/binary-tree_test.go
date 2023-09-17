package datastructures

import (
	"reflect"
	"testing"
)

//	 1
//	 /\
//	2   3
// /\  /\
// 4 5 6 7
// /
// 8

var multipleNodesTree = &Node{
	Left: &Node{
		Left: &Node{
			Left: &Node{
				value: 8,
			},
			value: 4,
		},
		Right: &Node{
			value: 5,
		},
		value: 2,
	},
	Right: &Node{
		Left: &Node{
			value: 6,
		},
		Right: &Node{
			value: 7,
		},
		value: 3,
	},
	value: 1,
}

var tree = &Tree{Root: multipleNodesTree}

func TestPrintTree(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "a tree with one node",
			args: args{
				root: &Node{
					value: 1,
				},
			},
		},
		{
			name: "a tree with multiple nodes",
			args: args{
				root: multipleNodesTree,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintTree(tt.args.root)
		})
	}
}

func TestGetTreeValues(t *testing.T) {
	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "a tree with one node",
			args: args{
				node: &Node{
					value: 1,
				},
			},
			want: []int{1},
		},
		{
			name: "a tree with multiple nodes",
			args: args{
				node: multipleNodesTree,
			},
			want: []int{8, 4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTreeValues(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTreeValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		tree  *Tree
		value int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "value in the root of the tree",
			args: args{
				tree:  tree,
				value: 1,
			},
			want: multipleNodesTree,
		},
		{
			name: "value at the bottom left of the tree",
			args: args{
				tree:  tree,
				value: 8,
			},
			want: &Node{
				value: 8,
			},
		},
		{
			name: "value at the bottom right of the tree",
			args: args{
				tree:  tree,
				value: 7,
			},
			want: &Node{
				value: 7,
			},
		},
		{
			name: "value in the middle of the tree",
			args: args{
				tree:  tree,
				value: 3,
			},
			want: multipleNodesTree.Right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.tree, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args struct {
		tree  *Tree
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "inserts into empty tree",
			args: args{
				tree:  &Tree{},
				value: 1,
			},
			want: []int{1},
		},
		{
			name: "inserts into level 1",
			args: args{
				tree: &Tree{
					Root: &Node{
						value: 1,
					},
				},
				value: 2,
			},
			want: []int{2, 1},
		},
		{
			name: "inserts into level 1, to the right",
			args: args{
				tree: &Tree{
					Root: &Node{
						Left: &Node{
							value: 2,
						},
						value: 1,
					},
				},
				value: 3,
			},
			want: []int{2, 1, 3},
		},
		{
			name: "inserts into bigger tree",
			args: args{
				tree: &Tree{
					Root: multipleNodesTree,
				},
				value: 12,
			},
			want: []int{8, 4, 12, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.tree, tt.args.value)
			values := GetTreeValues(tt.args.tree.Root)
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("TreeValues() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		tree  *Tree
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "delete an element that does not exist in the tree",
			args: args{
				tree:  tree,
				value: 800,
			},
			want: []int{8, 4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "delete bottom-left element",
			args: args{
				tree:  tree,
				value: 8,
			},
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name: "delete bottom-right element",
			args: args{
				tree:  tree,
				value: 7,
			},
			want: []int{8, 4, 2, 5, 1, 6, 3},
		},
		{
			name: "delete element in the middle",
			args: args{
				tree:  tree,
				value: 3,
			},
			want: []int{8, 4, 2, 5, 1, 7, 6},
		},
		{
			name: "delete element with children nodes that have children nodes",
			args: args{
				tree:  tree,
				value: 2,
			},
			want: []int{8, 4, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treeCopy := &Tree{}
			DeepCopy(tree, treeCopy)

			Delete(treeCopy, tt.args.value)
			values := GetTreeValues(treeCopy.Root)
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("TreeValues() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestCopyValues(t *testing.T) {
	type args struct {
		src *Tree
		dst *Tree
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tree",
			args: args{
				src: tree,
				dst: &Tree{},
			},
			want: []int{8, 4, 2, 5, 1, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyValues(tt.args.src, tt.args.dst)
			values := GetTreeValues(tt.args.dst.Root)
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("TreeValues() = %v, want %v", values, tt.want)
			}
		})
	}
}

func TestDeepCopy(t *testing.T) {
	type args struct {
		src *Tree
		dst *Tree
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tree",
			args: args{
				src: tree,
				dst: &Tree{},
			},
			want: []int{8, 4, 2, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeepCopy(tt.args.src, tt.args.dst)
			values := GetTreeValues(tt.args.dst.Root)
			if !reflect.DeepEqual(values, tt.want) {
				t.Errorf("TreeValues() = %v, want %v", values, tt.want)
			}
		})
	}
}
