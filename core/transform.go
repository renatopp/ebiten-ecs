package core

import (
	sk "github.com/renatopp/skald"
)

type DTransform struct {
	enabled  bool
	posX     float64
	posY     float64
	scaleX   float64
	scaleY   float64
	rotation float64 // in degrees
	version  uint64
}

func (c *DTransform) IsEnabled() bool {
	return c.enabled
}

func (c *DTransform) SetEnabled(enabled bool) {
	if c.enabled == enabled {
		return
	}

	c.enabled = enabled
	c.version++
}

func (c *DTransform) GetPosition() (x, y float64) {
	return c.posX, c.posY
}

func (c *DTransform) X() float64 {
	return c.posX
}

func (c *DTransform) Y() float64 {
	return c.posY
}

func (c *DTransform) MoveTo(x, y float64) {
	c.posX = x
	c.posY = y
	c.version++
}

func (c *DTransform) MoveBy(dx, dy float64) {
	c.posX += dx
	c.posY += dy
	c.version++
}

func (c *DTransform) GetScale() (x, y float64) {
	return c.scaleX, c.scaleY
}

func (c *DTransform) ScaleTo(x, y float64) {
	c.scaleX = x
	c.scaleY = y
	c.version++
}

func (c *DTransform) ScaleBothTo(s float64) {
	c.scaleX = s
	c.scaleY = s
	c.version++
}

func (c *DTransform) ScaleBy(dx, dy float64) {
	c.scaleX += dx
	c.scaleY += dy
	c.version++
}

func (c *DTransform) GetRotation() float64 {
	return c.rotation
}

func (c *DTransform) RotateTo(deg float64) {
	c.rotation = deg
	c.version++
}

func (c *DTransform) RotateBy(deg float64) {
	c.rotation += deg
	c.version++
}

func (c *DTransform) Reset() {
	c.posX = 0
	c.posY = 0
	c.scaleX = 1
	c.scaleY = 1
	c.rotation = 0
	c.version = 0
}

var Transform = sk.NewComponentWithOptions(sk.ComponentOptions[DTransform]{
	OnCreated: func(c *DTransform) {
		c.enabled = true
		c.Reset()
	},
})
