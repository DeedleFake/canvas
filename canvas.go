package canvas

import (
	"errors"
	"fmt"
	"github.com/gopherjs/gopherjs/js"
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

func New(elem *js.Object) (*Canvas, error) {
	if js.Global.Get("CanvasRenderingContext2D") == nil {
		return nil, errors.New("Browser doesn't support canvas 2D")
	}

	return &Canvas{
		elem: elem,
		ctx:  elem.Call("getContext", "2d"),
	}, nil
}

func (c *Canvas) Rect(r Rectangle) Object {
	r = r.Canon()

	p := func(pb *PathBuilder) {
		pb.Rect(r)
	}

	return &pathObj{
		c:    c,
		path: p,
	}
}

func (c *Canvas) Path(p func(*PathBuilder)) Object {
	return &pathObj{
		c:    c,
		path: p,
	}
}

func (c *Canvas) Text(text string, mw float64) Object {
	return &textObj{
		c:    c,
		text: text,
		mw:   mw,
	}
}

func (c *Canvas) Clear(r Rectangle) {
	r = r.Canon()
	c.ctx.Call("clearRect", r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

func (c *Canvas) MeasureText(text string) (tm TextMetrics) {
	return TextMetrics{c.ctx.Call("measureText", text)}
}

func (c *Canvas) Width() float64 {
	return c.elem.Get("width").Float()
}

func (c *Canvas) Height() float64 {
	return c.elem.Get("height").Float()
}

type Font struct {
	Size int
	Name string
}

func (f Font) String() string {
	return fmt.Sprintf("%v %v", f.Size, f.Name)
}
