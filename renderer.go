package sk

import (
	"container/heap"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	totalLayers uint
	layers      []RenderLayer
	itemBuffer  []*renderItem
}

func NewRenderer() *Renderer {
	r := &Renderer{
		totalLayers: 0,
		layers:      make([]RenderLayer, 0),
	}
	r.AddLayer()
	return r
}

func (r *Renderer) AddLayer() uint {
	layer := RenderLayer{}

	r.totalLayers++
	r.layers = append(r.layers, layer)
	heap.Init(&layer)

	return r.totalLayers - 1
}

func (r *Renderer) Queue(layer uint, zindex int, image *ebiten.Image, op *ebiten.DrawImageOptions) {
	if layer >= r.totalLayers {
		layer = r.totalLayers - 1
	}

	if len(r.itemBuffer) > 0 {
		r.itemBuffer[0].zindex = zindex
		r.itemBuffer[0].image = image
		r.itemBuffer[0].op = op
		heap.Push(&r.layers[layer], r.itemBuffer[0])
		r.itemBuffer = r.itemBuffer[1:]
		return
	} else {
		heap.Push(&r.layers[layer], &renderItem{zindex, len(r.layers[layer]), image, op})
	}
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	for _, layer := range r.layers {
		for _, item := range layer {
			screen.DrawImage(item.image, item.op)
		}
	}
}

func (r *Renderer) Clear() {
	layers := make([]RenderLayer, 0)
	for _, layer := range r.layers {
		for _, item := range layer {
			r.itemBuffer = append(r.itemBuffer, item)
		}

		l := RenderLayer{}
		heap.Init(&l)
		layers = append(layers, l)
	}

	r.layers = layers
}
