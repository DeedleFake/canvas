package canvas

import (
	"image"
)

type Object interface {
	Stroke(image.Point)
	Fill(image.Point)
}

type pathObj struct {
	c    *Canvas
	path func(*PathBuilder)
}

func (p pathObj) Fill(pt image.Point) {
	pb := &PathBuilder{p.c}

	pb.begin(pt)
	p.path(pb)

	p.c.ctx.Call("fill")
}

func (p pathObj) Stroke(pt image.Point) {
	pb := &PathBuilder{p.c}

	pb.begin(pt)
	p.path(pb)

	p.c.ctx.Call("stroke")
}

type PathBuilder struct {
	c *Canvas
}

func (pb *PathBuilder) begin(p image.Point) {
	pb.c.ctx.Call("beginPath")
	pb.c.ctx.Call("moveTo", p.X, p.Y)
}

func (pb PathBuilder) Rect(r image.Rectangle) {
	r = r.Canon()
	pb.c.ctx.Call("rect", r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

func (pb PathBuilder) MoveTo(p image.Point) {
	pb.c.ctx.Call("moveTo", p.X, p.Y)
}

func (pb PathBuilder) Line(p image.Point) {
	pb.c.ctx.Call("lineTo", p.X, p.Y)
}

func (pb PathBuilder) Bezier(cp1, cp2, end image.Point) {
	pb.c.ctx.Call("bezierCurveTo",
		cp1.X,
		cp1.Y,
		cp2.X,
		cp2.Y,
		end.X,
		end.Y,
	)
}

func (pb PathBuilder) Quadratic(cp, end image.Point) {
	pb.c.ctx.Call("quadraticCurveTo",
		cp.X,
		cp.Y,
		end.X,
		end.Y,
	)
}

func (pb PathBuilder) Arc(c image.Point, r float64, sa, ea float64, cc bool) {
	pb.c.ctx.Call("arc",
		c.X,
		c.Y,
		r,
		sa,
		ea,
		cc,
	)
}

type textObj struct {
	c    *Canvas
	text string
	mw   int
}

func (t textObj) Fill(p image.Point) {
	if t.mw < 0 {
		t.c.ctx.Call("fillText", t.text, p.X, p.Y)
		return
	}

	t.c.ctx.Call("fillText", t.text, p.X, p.Y, t.mw)
}

func (t textObj) Stroke(p image.Point) {
	if t.mw < 0 {
		t.c.ctx.Call("strokeText", t.text, p.X, p.Y)
		return
	}

	t.c.ctx.Call("strokeText", t.text, p.X, p.Y, t.mw)
}
