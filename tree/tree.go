package tree

type TreeNode[T any] struct {
	Value T

	Child []*TreeNode[T]
}

func (t *TreeNode[T]) Add(val T) *TreeNode[T] {
	n := &TreeNode[T]{
		Value: val,
	}
	t.Child = append(t.Child, n)
	return n
}

func (t *TreeNode[T]) Preorder(fn func(val T)) {
	if t == nil {
		return
	}
	fn(t.Value)

	for _, n := range t.Child {
		n.Preorder(fn)
	}
}

func (t *TreeNode[T]) Postorder(fn func(val T)) {
	if t == nil {
		return
	}

	for _, n := range t.Child {
		n.Postorder(fn)
	}

	fn(t.Value)
}

func (t *TreeNode[T]) BFS(fn func(val T)) {
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, t)

	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		fn(front.Value)

		queue = append(queue, front.Child...)

	}
}

// for 문을 사용한 dfs
func (t *TreeNode[T]) DFS(fn func(val T)) {
	stack := []*TreeNode[T]{}
	stack = append(stack, t)

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fn(last.Value)

		stack = append(stack, last.Child...)
	}
}
