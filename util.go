// +build wasm

package canvas

import (
	"syscall/js"
)

func getFloats(obj js.Value) (ret []float64) {
	if (obj == js.Undefined()) || (obj == js.Null()) || (obj.Length() == 0) {
		return nil
	}

	ret = make([]float64, 0, obj.Length())
	for i := 0; i < obj.Length(); i++ {
		ret = append(ret, obj.Index(i).Float())
	}

	return
}
