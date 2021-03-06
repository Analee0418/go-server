// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package protocol

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
)

func writeMapDouble(r *MapDouble, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteDouble(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapDouble struct {
	keys []string
	values []float64
	M map[string]float64
}

func NewMapDouble() *MapDouble{
	return &MapDouble {
		keys: make([]string, 0),
		values: make([]float64, 0),
		M: make(map[string]float64),
	}
}

func (_ *MapDouble) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MapDouble) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MapDouble) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MapDouble) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MapDouble) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MapDouble) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MapDouble) SetString(v string) { panic("Unsupported operation") }
func (_ *MapDouble) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *MapDouble) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapDouble) SetDefault(i int) { panic("Unsupported operation") }
func (r *MapDouble) Finalize() { 
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapDouble) AppendMap(key string) types.Field { 
	r.keys = append(r.keys, key)
	var v float64
	
	r.values = append(r.values, v)
	
	return (*types.Double)(&r.values[len(r.values)-1])
	
}

func (_ *MapDouble) AppendArray() types.Field { panic("Unsupported operation") }

