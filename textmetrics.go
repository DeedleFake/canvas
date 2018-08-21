// +build wasm

package canvas

import (
	"syscall/js"
)

type TextMetrics struct {
	tm js.Value
}

func (tm TextMetrics) Width() float64 {
	return tm.tm.Get("width").Float()
}

func (tm TextMetrics) ActualBoundingBoxLeft() float64 {
	return tm.tm.Get("actualBoundingBoxLeft").Float()
}

func (tm TextMetrics) ActualBoundingBoxRight() float64 {
	return tm.tm.Get("actualBoundingBoxRight").Float()
}

func (tm TextMetrics) ActualBoundingBoxAscent() float64 {
	return tm.tm.Get("actualBoundingBoxAscent").Float()
}

func (tm TextMetrics) ActualBoundingBoxDescent() float64 {
	return tm.tm.Get("actualBoundingBoxDescent").Float()
}

func (tm TextMetrics) FontBoundingBoxAscent() float64 {
	return tm.tm.Get("fontBoundingBoxAscent").Float()
}

func (tm TextMetrics) FontBoundingBoxDescent() float64 {
	return tm.tm.Get("fontBoundingBoxDescent").Float()
}

func (tm TextMetrics) EmHeightAscent() float64 {
	return tm.tm.Get("emHeightAscent").Float()
}

func (tm TextMetrics) EmHeightDescent() float64 {
	return tm.tm.Get("emHeightDescent").Float()
}

func (tm TextMetrics) HangingBaseline() float64 {
	return tm.tm.Get("hangingBaseline").Float()
}

func (tm TextMetrics) AlphabeticBaseline() float64 {
	return tm.tm.Get("alphabeticBaseline").Float()
}

func (tm TextMetrics) IdeographicBaseline() float64 {
	return tm.tm.Get("ideographicBaseline").Float()
}
