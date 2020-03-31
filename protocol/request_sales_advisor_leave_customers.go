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

  
type RequestSalesAdvisorLeaveCustomers struct {

	
	
		Customers_info *MessageCustomersInfo
	

}

func NewRequestSalesAdvisorLeaveCustomers() (*RequestSalesAdvisorLeaveCustomers) {
	return &RequestSalesAdvisorLeaveCustomers{}
}

func DeserializeRequestSalesAdvisorLeaveCustomers(r io.Reader) (*RequestSalesAdvisorLeaveCustomers, error) {
	t := NewRequestSalesAdvisorLeaveCustomers()
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

func DeserializeRequestSalesAdvisorLeaveCustomersFromSchema(r io.Reader, schema string) (*RequestSalesAdvisorLeaveCustomers, error) {
	t := NewRequestSalesAdvisorLeaveCustomers()

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

func writeRequestSalesAdvisorLeaveCustomers(r *RequestSalesAdvisorLeaveCustomers, w io.Writer) error {
	var err error
	
	err = writeMessageCustomersInfo( r.Customers_info, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r *RequestSalesAdvisorLeaveCustomers) Serialize(w io.Writer) error {
	return writeRequestSalesAdvisorLeaveCustomers(r, w)
}

func (r *RequestSalesAdvisorLeaveCustomers) Schema() string {
	return "{\"fields\":[{\"name\":\"customers_info\",\"type\":{\"fields\":[{\"name\":\"mobile\",\"type\":[\"null\",\"string\"]},{\"name\":\"mobileRegion\",\"type\":[\"null\",\"string\"]},{\"name\":\"idcard\",\"type\":[\"null\",\"string\"]},{\"name\":\"username\",\"type\":[\"null\",\"string\"]}],\"name\":\"MessageCustomersInfo\",\"namespace\":\"proto\",\"type\":\"record\"}}],\"name\":\"proto.RequestSalesAdvisorLeaveCustomers\",\"type\":\"record\"}"
}

func (r *RequestSalesAdvisorLeaveCustomers) SchemaName() string {
	return "proto.RequestSalesAdvisorLeaveCustomers"
}

func (_ *RequestSalesAdvisorLeaveCustomers) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) SetInt(v int32) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) SetLong(v int64) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) SetString(v string) { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RequestSalesAdvisorLeaveCustomers) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.Customers_info = NewMessageCustomersInfo()
	
		
		
			return r.Customers_info
		
	
	}
	panic("Unknown field index")
}

func (r *RequestSalesAdvisorLeaveCustomers) SetDefault(i int) {
	switch (i) {
	
        
	
	}
	panic("Unknown field index")
}

func (_ *RequestSalesAdvisorLeaveCustomers) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *RequestSalesAdvisorLeaveCustomers) Finalize() { }
