package main

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "Case 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 6,
							},
							Right: &TreeNode{
								Val: 7,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 8,
							Left: &TreeNode{
								Val: 9,
							},
						},
					},
				},
			},
			want: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name: "Case 3",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "Case 4",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
