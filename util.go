package canvas

import (
	"github.com/gopherjs/gopherjs/js"
)

func getFloats(obj *js.Object) (ret []float64) {
	if (obj == nil) || (obj.Length() == 0) {
		return nil
	}

	ret = make([]float64, 0, obj.Length())
	for i := 0; i < obj.Length(); i++ {
		ret = append(ret, obj.Index(i).Float())
	}

	return
}
