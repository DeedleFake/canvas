// +build wasm

package canvas

import (
	"errors"
	"fmt"
	"image/color"
	"syscall/js"
)

func hexColor(c color.Color) string {
	r, g, b, _ := c.RGBA()
	r = r * 0xff / 0xffff
	g = g * 0xff / 0xffff
	b = b * 0xff / 0xffff

	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// Canvas represents an HTML <canvas> element and its associated 2D
// rendering context.
type Canvas struct {
	elem js.Value
	ctx  js.Value
}

// New returns a new Canvas that wraps the <canvas> elem. This will
// most likely be retrieved using
//
//     js.Global.Get("document").Call("getElementById", id)
//
// If the browser doesn't support CanvasRenderingContext2D, an error
// is returned.
func New(elem js.Value) (*Canvas, error) {
	if js.Global().Get("CanvasRenderingContext2D") == js.Undefined() {
		// TODO: Is this necessary? I would think that anything that
		// supports WebAssembly should probably also have support for
		// canvases. Maybe Node.js?
		return nil, errors.New("Browser doesn't support canvas 2D")
	}

	return &Canvas{
		elem: elem,
		ctx:  elem.Call("getContext", "2d"),
	}, nil
}

// Rect returns an Object that draws a rectangle with the given
// bounds.
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

// Path returns an Object that draws a path build by the function p.
// For more information, see the documentation for PathBuilder.
func (c *Canvas) Path(p func(*PathBuilder)) Object {
	return &pathObj{
		c:    c,
		path: p,
	}
}

// Text returns an Object that draws the string text. If mw is a
// positive number, than it specifies the maximumWidth parameter of
// the canvas context's fillText() and drawText() methods.
func (c *Canvas) Text(text string, mw float64) Object {
	return &textObj{
		c:    c,
		text: text,
		mw:   mw,
	}
}

// Clear clears the pixels in the rectangle specified by r.
func (c *Canvas) Clear(r Rectangle) {
	r = r.Canon()
	c.ctx.Call("clearRect", r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

// MeasureText returns a TextMetrics containing information about the
// string text.
func (c *Canvas) MeasureText(text string) (tm TextMetrics) {
	return TextMetrics{c.ctx.Call("measureText", text)}
}

// Width returns the width of the canvas.
func (c *Canvas) Width() float64 {
	return c.elem.Get("width").Float()
}

// Height returns the height of the canvas.
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
