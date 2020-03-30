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

  
type RequestSalesAdvisorReceivingCustomers struct {

	
	
		Customers_info *MessageCustomersInfo
	

}

func NewRequestSalesAdvisorReceivingCustomers() (*RequestSalesAdvisorReceivingCustomers) {
	return &RequestSalesAdvisorReceivingCustomers{}
}

func DeserializeRequestSalesAdvisorReceivingCustomers(r io.Reader) (*RequestSalesAdvisorReceivingCustomers, error) {
	t := NewRequestSalesAdvisorReceivingCustomers()
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

func DeserializeRequestSalesAdvisorReceivingCustomersFromSchema(r io.Reader, schema string) (*RequestSalesAdvisorReceivingCustomers, error) {
	t := NewRequestSalesAdvisorReceivingCustomers()

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

func writeRequestSalesAdvisorReceivingCustomers(r *RequestSalesAdvisorReceivingCustomers, w io.Writer) error {
	var err error
	
	err = writeMessageCustomersInfo( r.Customers_info, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *RequestSalesAdvisorReceivingCustomers) Serialize(w io.Writer) error {
	return writeRequestSalesAdvisorReceivingCustomers(r, w)
}

func (r *RequestSalesAdvisorReceivingCustomers) Schema() string {
	return "{\"fields\":[{\"name\":\"customers_info\",\"type\":{\"fields\":[{\"name\":\"mobile\",\"type\":[\"null\",\"string\"]},{\"name\":\"idcard\",\"type\":[\"null\",\"string\"]},{\"name\":\"username\",\"type\":[\"null\",\"string\"]}],\"name\":\"MessageCustomersInfo\",\"namespace\":\"proto\",\"type\":\"record\"}}],\"name\":\"proto.RequestSalesAdvisorReceivingCustomers\",\"type\":\"record\"}"
}

func (r *RequestSalesAdvisorReceivingCustomers) SchemaName() string {
	return "proto.RequestSalesAdvisorReceivingCustomers"
}

func (_ *RequestSalesAdvisorReceivingCustomers) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) SetInt(v int32) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) SetLong(v int64) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) SetString(v string) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RequestSalesAdvisorReceivingCustomers) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.Customers_info = NewMessageCustomersInfo()
	
		
		
			return r.Customers_info
		
	
	}
	panic("Unknown field index")
}

func (r *RequestSalesAdvisorReceivingCustomers) SetDefault(i int) {
	switch (i) {
	
        
	
	}
	panic("Unknown field index")
}

func (_ *RequestSalesAdvisorReceivingCustomers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorReceivingCustomers) Finalize() { }
