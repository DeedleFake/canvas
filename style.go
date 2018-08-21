// +build wasm

package canvas

import (
	"fmt"
)

type style struct {
	lineWidth  float64
	lineCap    LineCap
	lineJoin   LineJoin
	miterLimit float64
	lineDash   []float64
}

func (s *style) set(c *Canvas) {
	c.ctx.Set("lineWidth", s.lineWidth)
	c.ctx.Set("lineCap", s.lineCap.String())
	c.ctx.Set("lineJoin", s.lineJoin.String())
	c.ctx.Set("miterLimit", s.miterLimit)
	c.ctx.Set("lineDash", s.lineDash)
}

func (s *style) SetLineWidth(width float64) {
	s.lineWidth = width
}

func (s *style) SetLineCap(lc LineCap) {
	s.lineCap = lc
}

func (s *style) SetLineJoin(join LineJoin) {
	s.lineJoin = join
}

func (s *style) SetMiterLimit(limit float64) {
	s.miterLimit = limit
}

func (s *style) SetLineDash(dash []float64) {
	s.lineDash = dash
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
