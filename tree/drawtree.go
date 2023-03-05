package tree

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type drawTreeNode[T any] struct {
	Value T
	X     int
	Y     int

	Child []*drawTreeNode[T]
}

func makeDrawTree[T any](node *TreeNode[T], level int, order *int) *drawTreeNode[T] {
	// order *int 값은 변경된 값이 유지되어야 하기 때문애 pointer 형으로
	if node == nil {
		return nil
	}

	// x, y 좌표가 셋팅된 tree node를 그리기 위해 생성
	drawNode := &drawTreeNode[T]{
		Value: node.Value,
		Y:     level,
		X:     *order,
	}

	//in-order
	// 왼쪽먼저 오른쪽
	// 여기서는 자식들을 반으로 나누어서 왼쪽자식들 본 후 오른쪽 자식들
	half := len(node.Child) / 2

	// pointer를 값 type으로 바꾸어서 1증가
	(*order)++

	// 왼쪽
	for i := 0; i < half; i++ {
		child := node.Child[i]
		drawNode.Child = append(drawNode.Child, makeDrawTree(child, level-1, order))
	}

	// setting x 좌표
	drawNode.X = *order
	(*order)++

	// 오른쪽
	for i := half; i < len(node.Child); i++ {
		child := node.Child[i]
		drawNode.Child = append(drawNode.Child, makeDrawTree(child, level-1, order))
	}

	return drawNode
}

func SaveTreeGraph[T any](t *TreeNode[T], filepath string) error {
	var order int
	drawNode := makeDrawTree(t, 0, &order)
	if drawNode == nil {
		return fmt.Errorf("empty Tree")
	}

	p := plot.New()

	// drawing

	var xys plotter.XYs
	// 위치 정보 모아주는 function
	drawNode.getLocations(&xys)
	points, err := plotter.NewScatter(xys)
	if err != nil {
		return err
	}

	points.Shape = draw.CircleGlyph{}
	points.Color = color.Gray16{}
	points.Radius = vg.Points(20)

	// 라인
	err = drawLines(drawNode, p)
	if err != nil {
		return err
	}

	p.Add(points)

	// add labels
	err = addLabel(drawNode, p)
	if err != nil {
		return err
	}

	return p.Save(1000, 600, filepath)
}

func (d *drawTreeNode[T]) getLocations(xys *plotter.XYs) {
	// node의 x,y 좌표 값을 append
	*xys = append(*xys, plotter.XY{
		X: float64(d.X),
		Y: float64(d.Y),
	})

	for _, c := range d.Child {
		c.getLocations(xys)
	}
}

func drawLines[T any](node *drawTreeNode[T], p *plot.Plot) error {
	for _, c := range node.Child {
		pts := plotter.XYs{
			{X: float64(node.X), Y: float64(node.Y)},
			{X: float64(c.X), Y: float64(c.Y)},
		}

		line, err := plotter.NewLine(pts)
		if err != nil {
			return err
		}

		scatter, err := plotter.NewScatter(pts)
		if err != nil {
			return err
		}

		p.Add(line, scatter)

		err = drawLines(c, p)
		if err != nil {
			return nil
		}

	}
	return nil
}

func addLabel[T any](node *drawTreeNode[T], p *plot.Plot) error {
	label, err := plotter.NewLabels(plotter.XYLabels{
		XYs: []plotter.XY{
			{X: float64(node.X), Y: float64(node.Y)},
		},
		Labels: []string{fmt.Sprint(node.Value)},
	})
	if err != nil {
		return err
	}

	p.Add(label)
	for _, c := range node.Child {
		err = addLabel(c, p)
		if err != nil {
			return err
		}
	}

	return nil
}
