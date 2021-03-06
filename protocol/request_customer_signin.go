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
	"github.com/actgardner/gogen-avro/compiler"
)

  
type RequestCustomerSignin struct {

	
	
		Mobile *MobileUnion
	

	
	
		Idcard *IdcardUnion
	

	
	
		Username *UsernameUnion
	

}

func NewRequestCustomerSignin() (*RequestCustomerSignin) {
	return &RequestCustomerSignin{}
}

func DeserializeRequestCustomerSignin(r io.Reader) (*RequestCustomerSignin, error) {
	t := NewRequestCustomerSignin()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func DeserializeRequestCustomerSigninFromSchema(r io.Reader, schema string) (*RequestCustomerSignin, error) {
	t := NewRequestCustomerSignin()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err	
	}
	return t, err
}

func writeRequestCustomerSignin(r *RequestCustomerSignin, w io.Writer) error {
	var err error
	
	err = writeMobileUnion( r.Mobile, w)
	if err != nil {
		return err			
	}
	
	err = writeIdcardUnion( r.Idcard, w)
	if err != nil {
		return err			
	}
	
	err = writeUsernameUnion( r.Username, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *RequestCustomerSignin) Serialize(w io.Writer) error {
	return writeRequestCustomerSignin(r, w)
}

func (r *RequestCustomerSignin) Schema() string {
	return "{\"fields\":[{\"name\":\"mobile\",\"type\":[\"null\",\"string\"]},{\"name\":\"idcard\",\"type\":[\"null\",\"string\"]},{\"name\":\"username\",\"type\":[\"null\",\"string\"]}],\"name\":\"proto.RequestCustomerSignin\",\"type\":\"record\"}"
}

func (r *RequestCustomerSignin) SchemaName() string {
	return "proto.RequestCustomerSignin"
}

func (_ *RequestCustomerSignin) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) SetInt(v int32) { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) SetLong(v int64) { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) SetString(v string) { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RequestCustomerSignin) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.Mobile = NewMobileUnion()
	
		
		
			return r.Mobile
		
	
	case 1:
		
			r.Idcard = NewIdcardUnion()
	
		
		
			return r.Idcard
		
	
	case 2:
		
			r.Username = NewUsernameUnion()
	
		
		
			return r.Username
		
	
	}
	panic("Unknown field index")
}

func (r *RequestCustomerSignin) SetDefault(i int) {
	switch (i) {
	
        
	
        
	
        
	
	}
	panic("Unknown field index")
}

func (_ *RequestCustomerSignin) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *RequestCustomerSignin) Finalize() { }
