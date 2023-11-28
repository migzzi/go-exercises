package kdtree

// Represents a K-Dimensional data point.
type KDPoint[T any] interface {
	GetDimensionValue(n int) T //Returns the value of the nth dimension
	Dimensions() int           //Returns the number of dimensions
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
	panic("not implemented")
}

func (t *KDTree[T]) Insert(p KDPoint[T]) {
	panic("not implemented")
}

func buildTree[T any](points []KDPoint[T], dstFn KDistanceCalculator[T]) *Node[T] {
	panic("not implemented")
}

func (t *KDTree[T]) ForEach(fn func(*Node[T], int)) {
	traverse(t.Root, 0, fn)
}
