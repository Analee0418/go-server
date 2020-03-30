package model

import avro "com.lueey.shop/protocol"

func GenerateMessage(action avro.Action) *avro.Message {
	return &avro.Message{
		Action: action,

		SessionId: avro.NewSessionIdUnion(),

		// 销售端

		Sales_advisor_signin: avro.NewSales_advisor_signinUnion(),

		Sales_advisor_receiving_customers: avro.NewSales_advisor_receiving_customersUnion(),

		Sales_advisor_leave_customers: avro.NewSales_advisor_leave_customersUnion(),

		Sales_advisor_build_contract: avro.NewSales_advisor_build_contractUnion(),

		Sales_advisor_confirm_paid: avro.NewSales_advisor_confirm_paidUnion(),

		// 用户端

		Customer_signin: avro.NewCustomer_signinUnion(),

		Customer_auction_bid: avro.NewCustomer_auction_bidUnion(),

		Customer_join_queue: avro.NewCustomer_join_queueUnion(),

		// 主持人端

		Host_switch_state: avro.NewHost_switch_stateUnion(),

		// 推送消息

		Message_session: avro.NewMessage_sessionUnion(),

		Message_room_info: avro.NewMessage_room_infoUnion(),

		Message_room_waiting_customers: avro.NewMessage_room_waiting_customersUnion(),

		Message_auction_info: avro.NewMessage_auction_infoUnion(),

		Message_cars_model: avro.NewMessage_cars_modelUnion(),

		// 提示消息

		Tips: avro.NewTipsUnion(),

		Error_message: avro.NewError_messageUnion(),
	}
}
