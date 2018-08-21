// +build wasm

package canvas

// Object represents an object that can be drawn on the canvas.
type Object interface {
	// Stroke draws an outline of the Object.
	Stroke(Point)

	// Fill fills the area of the canvas represented by the Object.
	Fill(Point)

	// Set various styles having to do with drawing lines. For more
	// information, see the CanvasRenderingContext2D documentation.
	SetLineWidth(float64)
	SetLineCap(LineCap)
	SetLineJoin(LineJoin)
	SetMiterLimit(float64)
	SetLineDash([]float64)
}

type pathObj struct {
	c    *Canvas
	path func(*PathBuilder)

	style
}

func (p pathObj) Fill(pt Point) {
	pb := &PathBuilder{p.c}

	pb.begin(pt)
	p.path(pb)

	p.set(p.c)
	p.c.ctx.Call("fill")
}

func (p pathObj) Stroke(pt Point) {
	pb := &PathBuilder{p.c}

	pb.begin(pt)
	p.path(pb)

	p.set(p.c)
	p.c.ctx.Call("stroke")
}

// PathBuilder is a type that is passed to a function used to create a
// custom, path-based object. See the documentation for
// (*Canvas).Path().
type PathBuilder struct {
	c *Canvas
}

func (pb *PathBuilder) begin(p Point) {
	pb.c.ctx.Call("beginPath")
	pb.c.ctx.Call("moveTo", p.X, p.Y)
}

// Adds the rectangle r to the path.
func (pb PathBuilder) Rect(r Rectangle) {
	r = r.Canon()
	pb.c.ctx.Call("rect", r.Min.X, r.Min.Y, r.Dx(), r.Dy())
}

// Moves the current postion to p.
func (pb PathBuilder) MoveTo(p Point) {
	pb.c.ctx.Call("moveTo", p.X, p.Y)
}

// Add a line from the current position to p to the path.
func (pb PathBuilder) Line(p Point) {
	pb.c.ctx.Call("lineTo", p.X, p.Y)
}

// Add a bezier curve to the path.
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

// Add a quadratic curve to the path.
func (pb PathBuilder) Quadratic(cp, end Point) {
	pb.c.ctx.Call("quadraticCurveTo",
		cp.X,
		cp.Y,
		end.X,
		end.Y,
	)
}

// Add an arc to the path.
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

	style
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
