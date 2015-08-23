package canvas

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"image"
	"image/color"
)

func hexColor(c color.Color) string {
	r, g, b, _ := c.RGBA()
	r = r * 0xff / 0xffff
	g = g * 0xff / 0xffff
	b = b * 0xff / 0xffff

	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

type Canvas struct {
	elem *js.Object
	ctx  *js.Object
}

func NewCanvas(elem *js.Object) *Canvas {
	return &Canvas{
		elem: elem,
		ctx:  elem.Call("getContext", "2d"),
	}
}

func (c *Canvas) Text(text string, p image.Point, font *Font) {
	oldFont := c.ctx.Get("font")
	defer func() {
		c.ctx.Set("font", oldFont)
	}()

	c.ctx.Set("font", font.String())
	c.ctx.Call("fillText", text, p.X, p.Y)
}

func (c *Canvas) Rect(r image.Rectangle, col color.Color) {
	oldStyle := c.ctx.Get("fillStyle")
	defer func() {
		c.ctx.Set("fillStyle", oldStyle)
	}()

	c.ctx.Set("fillStyle", hexColor(col))
	c.ctx.Call("fillRect", r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

func (c *Canvas) Width() int {
	return c.ctx.Get("width").Int()
}

func (c *Canvas) Height() int {
	return c.ctx.Get("height").Int()
}

type Font struct {
	Size int
	Name string
}

func (f Font) String() string {
	return fmt.Sprintf("%v %v", f.Size, f.Name)
}
