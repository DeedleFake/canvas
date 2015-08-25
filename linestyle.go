package canvas

import (
	"fmt"
)

type LineStyle struct {
	LineWidth  float64
	LineCap    LineCap
	LineJoin   LineJoin
	MiterLimit float64
	LineDash   []float64
}

func (ls *LineStyle) set(c *Canvas) *LineStyle {
	if ls == nil {
		return nil
	}

	r := &LineStyle{
		LineWidth:  c.ctx.Get("lineWidth").Float(),
		LineCap:    parseLineCap(c.ctx.Get("lineCap").String()),
		LineJoin:   parseLineJoin(c.ctx.Get("lineJoin").String()),
		MiterLimit: c.ctx.Get("miterLimit").Float(),
		LineDash:   getFloats(c.ctx.Get("lineDash")),
	}

	c.ctx.Set("lineWidth", ls.LineWidth)
	c.ctx.Set("lineCap", ls.LineCap.String())
	c.ctx.Set("lineJoin", ls.LineJoin.String())
	c.ctx.Set("miterLimit", ls.MiterLimit)
	c.ctx.Set("lineDash", ls.LineDash)

	return r
}

type LineCap int

const (
	ButtCap LineCap = iota
	RoundCap
	SquareCap
)

func parseLineCap(lc string) LineCap {
	switch lc {
	case "butt":
		return ButtCap
	case "round":
		return RoundCap
	case "square":
		return SquareCap
	}

	panic("Unknown lineCap: " + lc)
}

func (lc LineCap) String() string {
	switch lc {
	case ButtCap:
		return "butt"
	case RoundCap:
		return "round"
	case SquareCap:
		return "square"
	}

	panic(fmt.Errorf("Unknown LineCap: %v", int(lc)))
}

type LineJoin int

const (
	BevelJoin LineJoin = iota
	RoundJoin
	MiterJoin
)

func parseLineJoin(lj string) LineJoin {
	switch lj {
	case "bevel":
		return BevelJoin
	case "round":
		return RoundJoin
	case "miter":
		return MiterJoin
	}

	panic("Unknown lineJoin: " + lj)
}

func (lj LineJoin) String() string {
	switch lj {
	case BevelJoin:
		return "bevel"
	case RoundJoin:
		return "round"
	case MiterJoin:
		return "miter"
	}

	panic(fmt.Errorf("Unknown LineJoin: %v", int(lj)))
}
