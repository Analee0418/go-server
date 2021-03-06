package model

import avro "com.lueey.shop/protocol"

func GenerateMessage(action avro.Action) *avro.Message {
	return &avro.Message{
		Action: action,

		SessionId: avro.NewSessionIdUnion(),

		// 销售端

		Sales_advisor_signin: avro.NewSales_advisor_signinUnion(),

		Sales_advisor_receiving_customers: avro.NewSales_advisor_receiving_customersUnion(),

		Sales_advisor_build_contract: avro.NewSales_advisor_build_contractUnion(),

		Sales_advisor_confirm_paid: avro.NewSales_advisor_confirm_paidUnion(),

		// 用户端

		Customer_signin: avro.NewCustomer_signinUnion(),

		Customer_auction_bid: avro.NewCustomer_auction_bidUnion(),

		Request_customer_build_signature: avro.NewRequest_customer_build_signatureUnion(),

		Customer_start_game: avro.NewCustomer_start_gameUnion(),

		Customer_upload_game_score: avro.NewCustomer_upload_game_scoreUnion(),

		Customer_update_state: avro.NewCustomer_update_stateUnion(),

		// 推送消息

		Message_session: avro.NewMessage_sessionUnion(),

		Message_room_info: avro.NewMessage_room_infoUnion(),

		Message_customer_info: avro.NewMessage_customer_infoUnion(),

		Message_room_waiting_customers: avro.NewMessage_room_waiting_customersUnion(),

		Message_auction_info: avro.NewMessage_auction_infoUnion(),

		Message_cars_model: avro.NewMessage_cars_modelUnion(),

		Message_forward_to_customer: avro.NewMessage_forward_to_customerUnion(),

		Message_forward_to_sales_advisor: avro.NewMessage_forward_to_sales_advisorUnion(),

		Message_broadcast: avro.NewMessage_broadcastUnion(),

		Message_global_state: avro.NewMessage_global_stateUnion(),

		Message_game_config: avro.NewMessage_game_configUnion(),

		Message_game_result: avro.NewMessage_game_resultUnion(),

		// 提示消息

		Tips: avro.NewTipsUnion(),

		Error_message: avro.NewError_messageUnion(),
	}
}
