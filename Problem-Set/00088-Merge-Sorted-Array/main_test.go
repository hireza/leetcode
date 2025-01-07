package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

func stringToIntSlice(s string) []int {
	splitS := strings.Split(s, " ")
	var intValues []int

	for _, str := range splitS {
		num, _ := strconv.Atoi(str)
		intValues = append(intValues, num)
	}

	return intValues
}

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Case 2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "Case 3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := capturer.CaptureOutput(func() {
				merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			})

			out = strings.ReplaceAll(out, "\n", "")
			out = strings.ReplaceAll(out, "[", "")
			out = strings.ReplaceAll(out, "]", "")

			res := stringToIntSlice(out)
			assert.Equal(t, tt.want, res)

		})
	}
}
