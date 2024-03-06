package render

import "github.com/hajimehoshi/ebiten/v2"

type renderItem struct {
	zindex int
	index  int
	image  *ebiten.Image
	op     *ebiten.DrawImageOptions
}

type RenderLayer []*renderItem

func (q RenderLayer) Len() int {
	return len(q)
}

func (q RenderLayer) Less(i, j int) bool {
	return q[i].zindex < q[j].zindex
}

func (q RenderLayer) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *RenderLayer) Push(x any) {
	item := x.(*renderItem)
	item.index = len(*q)
	*q = append(*q, item)
}

func (q *RenderLayer) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1
	*q = old[0 : n-1]
	return item
}
