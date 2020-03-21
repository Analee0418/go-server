// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package avro

import (
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/vm"
)

  
type Action int32

const (

	ActionHeartbeat Action = 0

	ActionRequest_sales_advisor_signin Action = 1

	ActionRequest_sales_advisor_receiving_customers Action = 2

	ActionRequest_sales_advisor_leave_customers Action = 3

	ActionRequest_sales_advisor_build_contract Action = 4

	ActionRequest_sales_advisor_confirm_paid Action = 5

	ActionMessage_room_info Action = 6

	ActionMessage_room_waiting_customers Action = 7

	ActionMessage_customers_info Action = 8

	ActionMessage_customers_auction_info Action = 9

	ActionMessage_cars_model Action = 10

	ActionMessage_contract Action = 11

	ActionError_message Action = 12

)

func (e Action) String() string {
	switch e {

	case ActionHeartbeat:
		return "heartbeat"

	case ActionRequest_sales_advisor_signin:
		return "request_sales_advisor_signin"

	case ActionRequest_sales_advisor_receiving_customers:
		return "request_sales_advisor_receiving_customers"

	case ActionRequest_sales_advisor_leave_customers:
		return "request_sales_advisor_leave_customers"

	case ActionRequest_sales_advisor_build_contract:
		return "request_sales_advisor_build_contract"

	case ActionRequest_sales_advisor_confirm_paid:
		return "request_sales_advisor_confirm_paid"

	case ActionMessage_room_info:
		return "message_room_info"

	case ActionMessage_room_waiting_customers:
		return "message_room_waiting_customers"

	case ActionMessage_customers_info:
		return "message_customers_info"

	case ActionMessage_customers_auction_info:
		return "message_customers_auction_info"

	case ActionMessage_cars_model:
		return "message_cars_model"

	case ActionMessage_contract:
		return "message_contract"

	case ActionError_message:
		return "error_message"

	}
	return "unknown"
}

func writeAction(r Action, w io.Writer) error {
	return vm.WriteInt(int32(r), w)
}

func NewActionValue(raw string) (r Action, err error) {
	switch raw {

	case "heartbeat":
		return ActionHeartbeat, nil

	case "request_sales_advisor_signin":
		return ActionRequest_sales_advisor_signin, nil

	case "request_sales_advisor_receiving_customers":
		return ActionRequest_sales_advisor_receiving_customers, nil

	case "request_sales_advisor_leave_customers":
		return ActionRequest_sales_advisor_leave_customers, nil

	case "request_sales_advisor_build_contract":
		return ActionRequest_sales_advisor_build_contract, nil

	case "request_sales_advisor_confirm_paid":
		return ActionRequest_sales_advisor_confirm_paid, nil

	case "message_room_info":
		return ActionMessage_room_info, nil

	case "message_room_waiting_customers":
		return ActionMessage_room_waiting_customers, nil

	case "message_customers_info":
		return ActionMessage_customers_info, nil

	case "message_customers_auction_info":
		return ActionMessage_customers_auction_info, nil

	case "message_cars_model":
		return ActionMessage_cars_model, nil

	case "message_contract":
		return ActionMessage_contract, nil

	case "error_message":
		return ActionError_message, nil

	}

	return -1, fmt.Errorf("invalid value for Action: '%s'", raw)
}