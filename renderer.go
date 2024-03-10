package sk

import (
	"container/heap"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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

func (r *Renderer) Draw(screen *ebiten.Image, ops *ebiten.DrawImageOptions) {
	for _, layer := range r.layers {
		for _, item := range layer {
			screen.DrawImage(item.image, item.op)
		}
	}

	im := ebiten.NewImage(1000, 1000)
	vector.DrawFilledCircle(im, 500, 500, 100, color.White, false)
	vector.DrawFilledCircle(im, 500, 500, 1, color.RGBA{255, 0, 0, 255}, false)
	vector.DrawFilledCircle(im, 700, 500, 50, color.White, false)
	vector.DrawFilledCircle(im, 300, 500, 50, color.White, false)
	vector.DrawFilledCircle(im, 500, 700, 50, color.White, false)
	vector.DrawFilledCircle(im, 500, 300, 50, color.White, false)

	ops2 := &ebiten.DrawImageOptions{}
	ops2.GeoM.Translate(PIXELS_PER_UNIT*SCREEN_WIDTH/2, PIXELS_PER_UNIT*SCREEN_HEIGHT/2)
	ops2.GeoM.Translate(-500, -500)

	// ops2.GeoM.Concat(ops.GeoM)

	screen.DrawImage(im, ops2)
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
