package kdtree

import (
	"fmt"
	"math"
	"strings"
)

func distance[T any](a, b KDPoint[T], dstFn KDistanceCalculator[T]) float64 {
	d := 0.0
	for i := 0; i < a.Dimensions(); i++ {
		d += math.Pow(dstFn(a, b, i), 2)
	}
	return d
}

func (t *KDTree[T]) print() {
	grid := buildTreeGrid(t.Root)
	for _, row := range grid {
		for _, c := range row {
			fmt.Print(c)
		}
		fmt.Println()
		fmt.Println()
	}
}

func buildTreeGrid[T any](root *Node[T]) [][]string {
	if root == nil {
		return [][]string{}
	}

	h := maxDepth(root)
	col := int(math.Pow(2, float64(h+1)) - 1)
	res := make([][]string, h+1)

	for i := 0; i < h+1; i++ {
		row := make([]string, col)
		// init res 2d arr
		for j := 0; j < col; j++ {
			row[j] = ""
		}
		res[i] = row
	}

	maxLen := fillNode(root, 0, col, 0, res)

	for i := 0; i < h+1; i++ {
		for j := 0; j < col; j++ {
			if res[i][j] == "" {
				res[i][j] = strings.Repeat(" ", maxLen)
			}
		}
	}

	return res
}

func fillNode[T any](n *Node[T], l, r, h int, res [][]string) int {
	if n == nil {
		return 1
	}

	maxLen := 0
	var mid int = (l + r) / 2
	if s, ok := n.Point.(fmt.Stringer); ok == true {
		res[h][mid] = s.String()
	} else {
		res[h][mid] = fmt.Sprintf("%v", n.Point)
	}

	if len(res[h][mid]) > maxLen {
		maxLen = len(res[h][mid])
	}

	if n.Left != nil {
		fillNode(n.Left, l, mid, h+1, res)
	}

	if n.Right != nil {
		fillNode(n.Right, mid+1, r, h+1, res)
	}

	return maxLen
}

func closer[T any](target, p1, p2 KDPoint[T], dstFn KDistanceCalculator[T]) KDPoint[T] {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	if distance(target, p1, dstFn) < distance(target, p2, dstFn) {
		return p1
	}
	return p2
}

func maxDepth[T any](n *Node[T]) int {
	if n == nil {
		return 0
	}

	lDepth := maxDepth(n.Left)
	rDepth := maxDepth(n.Right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func traverse[T any](node *Node[T], depth int, fn func(*Node[T], int)) {
	if node == nil {
		return
	}
	fn(node, depth)
	traverse(node.Left, depth+1, fn)
	traverse(node.Right, depth+1, fn)
}
