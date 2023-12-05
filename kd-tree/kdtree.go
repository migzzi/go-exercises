package kdtree

import (
	"math"
	"sort"
)

// Represents a K-Dimensional data point.
type KDPoint[T any] interface {
	GetDimensionValue(n int) T
	Dimensions() int
}

type KDistanceCalculator[T any] func(a, b KDPoint[T], dim int) float64

// a node in our tree
type Node[T any] struct {
	Point KDPoint[T] //Holds the data point.

	//Pointers to child nodes.
	Left  *Node[T]
	Right *Node[T]
}

type KDTree[T any] struct {
	Root  *Node[T]
	Size  int
	dstFn KDistanceCalculator[T] //Function to calculate the distance between two points at a given dimension.
}

func NewKDTree[T any](points []KDPoint[T], dstFn KDistanceCalculator[T]) *KDTree[T] {
	if dstFn == nil {
		panic("dstFn cannot be nil")
	}

	return &KDTree[T]{dstFn: dstFn, Root: buildTree(points, dstFn), Size: len(points)}
}

func (t *KDTree[T]) SearchNearest(target KDPoint[T]) KDPoint[T] {
	if t.Root == nil {
		return nil
	}
	return t.searchNearest(t.Root, target, 0)
}

func (t *KDTree[T]) searchNearest(root *Node[T], target KDPoint[T], depth int) KDPoint[T] {
	if root == nil {
		return nil
	}

	dim := depth % target.Dimensions()
	var nextBranch, oppositeBranch *Node[T]

	if t.dstFn(target, root.Point, dim) < 0 {
		nextBranch = root.Left
		oppositeBranch = root.Right
	} else {
		nextBranch = root.Right
		oppositeBranch = root.Left
	}

	tempBest := t.searchNearest(nextBranch, target, depth+1)
	best := closer(target, tempBest, root.Point, t.dstFn)

	distToBest := distance(target, best, t.dstFn)
	distToRoot := math.Pow(t.dstFn(target, root.Point, dim), 2)

	if distToRoot < distToBest {
		tempBest = t.searchNearest(oppositeBranch, target, depth+1)
		best = closer(target, tempBest, best, t.dstFn)
	}

	return best
}

func (t *KDTree[T]) Insert(p KDPoint[T]) {
	if t.Root == nil {
		t.Root = &Node[T]{Point: p}
		return
	}
	t.insert(t.Root, p, 0)
}

func (t *KDTree[T]) insert(n *Node[T], p KDPoint[T], depth int) {
	dim := depth % p.Dimensions()
	if t.dstFn(p, n.Point, dim) < 0 {
		if n.Left == nil {
			n.Left = &Node[T]{Point: p}
			return
		}
		t.insert(n.Left, p, depth+1)
	} else {
		if n.Right == nil {
			n.Right = &Node[T]{Point: p}
			return
		}
		t.insert(n.Right, p, depth+1)
	}
}

func buildTree[T any](points []KDPoint[T], dstFn KDistanceCalculator[T]) *Node[T] {
	return buildTreeWithSortedPoints(points, 0, dstFn)
}

func buildTreeWithSortedPoints[T any](
	points []KDPoint[T],
	depth int,
	dstFn KDistanceCalculator[T]) *Node[T] {
	if len(points) == 0 {
		return nil
	}
	pointsCpy := make([]KDPoint[T], len(points))
	copy(pointsCpy, points)
	dim := depth % points[0].Dimensions()
	median := len(pointsCpy) / 2

	sort.Slice(pointsCpy, func(a, b int) bool {
		return dstFn(pointsCpy[a], pointsCpy[b], dim) < 0
	})

	return &Node[T]{
		Point: pointsCpy[median],
		Left:  buildTreeWithSortedPoints(pointsCpy[:median], depth+1, dstFn),
		Right: buildTreeWithSortedPoints(pointsCpy[median+1:], depth+1, dstFn),
	}
}

func (t *KDTree[T]) ForEach(fn func(*Node[T], int)) {
	traverse(t.Root, 0, fn)
}
