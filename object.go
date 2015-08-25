package canvas

type Object interface {
	Stroke(Point)
	Fill(Point)

	SetLineStyle(*LineStyle)
}

type pathObj struct {
	c    *Canvas
	path func(*PathBuilder)
	ls   *LineStyle
}

func (p pathObj) Fill(pt Point) {
	pb := &PathBuilder{p.c}

	pb.begin(pt)
	p.path(pb)

	defer p.ls.set(p.c).set(p.c)
	p.c.ctx.Call("fill")
}

func (p pathObj) Stroke(pt Point) {
	pb := &PathBuilder{p.c}

	pb.begin(pt)
	p.path(pb)

	defer p.ls.set(p.c).set(p.c)
	p.c.ctx.Call("stroke")
}

func (p pathObj) SetLineStyle(ls *LineStyle) {
	p.ls = ls
}

type PathBuilder struct {
	c *Canvas
}

func (pb *PathBuilder) begin(p Point) {
	pb.c.ctx.Call("beginPath")
	pb.c.ctx.Call("moveTo", p.X, p.Y)
}

func (pb PathBuilder) Rect(r Rectangle) {
	r = r.Canon()
	pb.c.ctx.Call("rect", r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

func (pb PathBuilder) MoveTo(p Point) {
	pb.c.ctx.Call("moveTo", p.X, p.Y)
}

func (pb PathBuilder) Line(p Point) {
	pb.c.ctx.Call("lineTo", p.X, p.Y)
}

func (pb PathBuilder) Bezier(cp1, cp2, end Point) {
	pb.c.ctx.Call("bezierCurveTo",
		cp1.X,
		cp1.Y,
		cp2.X,
		cp2.Y,
		end.X,
		end.Y,
	)
}

func (pb PathBuilder) Quadratic(cp, end Point) {
	pb.c.ctx.Call("quadraticCurveTo",
		cp.X,
		cp.Y,
		end.X,
		end.Y,
	)
}

func (pb PathBuilder) Arc(c Point, r float64, sa, ea float64, cc bool) {
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
	mw   float64
}

func (t textObj) Fill(p Point) {
	if t.mw < 0 {
		t.c.ctx.Call("fillText", t.text, p.X, p.Y)
		return
	}

	t.c.ctx.Call("fillText", t.text, p.X, p.Y, t.mw)
}

func (t textObj) Stroke(p Point) {
	if t.mw < 0 {
		t.c.ctx.Call("strokeText", t.text, p.X, p.Y)
		return
	}

	t.c.ctx.Call("strokeText", t.text, p.X, p.Y, t.mw)
}

func (t textObj) SetLineStyle(*LineStyle) {
}
