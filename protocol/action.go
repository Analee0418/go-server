// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     lueey.avsc
 */
package protocol

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

	ActionRequest_customer_signin Action = 6

	ActionRequest_customer_auction_bid Action = 7

	ActionRequest_customer_join_queue Action = 8

	ActionRequest_customer_build_signature Action = 9

	ActionRequest_host_connect Action = 10

	ActionRequest_host_set_start_time Action = 11

	ActionRequest_host_switch_state Action = 12

	ActionRequest_host_choice_auction_goods Action = 13

	ActionRequest_host_auction_goods Action = 14

	ActionMessage_session Action = 15

	ActionMessage_room_info Action = 16

	ActionMessage_room_waiting_customers Action = 17

	ActionMessage_room_chat_ends Action = 18

	ActionMessage_customers_info Action = 19

	ActionMessage_customers_auction_info Action = 20

	ActionMessage_cars_model Action = 21

	ActionMessage_contract Action = 22

	ActionMessage_global_info Action = 23

	ActionMessage_auction_info Action = 24

	ActionMessage_forward_to_customer Action = 25

	ActionMessage_forward_to_sales_advisor Action = 26

	ActionMessage_broadcast Action = 27

	ActionTips Action = 28

	ActionError_message Action = 29

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

	case ActionRequest_customer_signin:
		return "request_customer_signin"

	case ActionRequest_customer_auction_bid:
		return "request_customer_auction_bid"

	case ActionRequest_customer_join_queue:
		return "request_customer_join_queue"

	case ActionRequest_customer_build_signature:
		return "request_customer_build_signature"

	case ActionRequest_host_connect:
		return "request_host_connect"

	case ActionRequest_host_set_start_time:
		return "request_host_set_start_time"

	case ActionRequest_host_switch_state:
		return "request_host_switch_state"

	case ActionRequest_host_choice_auction_goods:
		return "request_host_choice_auction_goods"

	case ActionRequest_host_auction_goods:
		return "request_host_auction_goods"

	case ActionMessage_session:
		return "message_session"

	case ActionMessage_room_info:
		return "message_room_info"

	case ActionMessage_room_waiting_customers:
		return "message_room_waiting_customers"

	case ActionMessage_room_chat_ends:
		return "message_room_chat_ends"

	case ActionMessage_customers_info:
		return "message_customers_info"

	case ActionMessage_customers_auction_info:
		return "message_customers_auction_info"

	case ActionMessage_cars_model:
		return "message_cars_model"

	case ActionMessage_contract:
		return "message_contract"

	case ActionMessage_global_info:
		return "message_global_info"

	case ActionMessage_auction_info:
		return "message_auction_info"

	case ActionMessage_forward_to_customer:
		return "message_forward_to_customer"

	case ActionMessage_forward_to_sales_advisor:
		return "message_forward_to_sales_advisor"

	case ActionMessage_broadcast:
		return "message_broadcast"

	case ActionTips:
		return "tips"

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

	case "request_customer_signin":
		return ActionRequest_customer_signin, nil

	case "request_customer_auction_bid":
		return ActionRequest_customer_auction_bid, nil

	case "request_customer_join_queue":
		return ActionRequest_customer_join_queue, nil

	case "request_customer_build_signature":
		return ActionRequest_customer_build_signature, nil

	case "request_host_connect":
		return ActionRequest_host_connect, nil

	case "request_host_set_start_time":
		return ActionRequest_host_set_start_time, nil

	case "request_host_switch_state":
		return ActionRequest_host_switch_state, nil

	case "request_host_choice_auction_goods":
		return ActionRequest_host_choice_auction_goods, nil

	case "request_host_auction_goods":
		return ActionRequest_host_auction_goods, nil

	case "message_session":
		return ActionMessage_session, nil

	case "message_room_info":
		return ActionMessage_room_info, nil

	case "message_room_waiting_customers":
		return ActionMessage_room_waiting_customers, nil

	case "message_room_chat_ends":
		return ActionMessage_room_chat_ends, nil

	case "message_customers_info":
		return ActionMessage_customers_info, nil

	case "message_customers_auction_info":
		return ActionMessage_customers_auction_info, nil

	case "message_cars_model":
		return ActionMessage_cars_model, nil

	case "message_contract":
		return ActionMessage_contract, nil

	case "message_global_info":
		return ActionMessage_global_info, nil

	case "message_auction_info":
		return ActionMessage_auction_info, nil

	case "message_forward_to_customer":
		return ActionMessage_forward_to_customer, nil

	case "message_forward_to_sales_advisor":
		return ActionMessage_forward_to_sales_advisor, nil

	case "message_broadcast":
		return ActionMessage_broadcast, nil

	case "tips":
		return ActionTips, nil

	case "error_message":
		return ActionError_message, nil

	}

	return -1, fmt.Errorf("invalid value for Action: '%s'", raw)
}
