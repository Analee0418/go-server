package model

import avro "com.lueey.shop/protocol"

func GenerateMessage(action avro.Action) *avro.Message {
	return &avro.Message{
		Action: action,

		Sales_advisor_signin: avro.NewUnionNullRequestSalesAdvisorSignin(),

		Sales_advisor_receiving_customers: avro.NewUnionNullRequestSalesAdvisorReceivingCustomers(),

		Sales_advisor_leave_customers: avro.NewUnionNullRequestSalesAdvisorLeaveCustomers(),

		Sales_advisor_build_contract: avro.NewUnionNullRequestSalesAdvisorBuildContract(),

		Sales_advisor_confirm_paid: avro.NewUnionNullRequestSalesAdvisorConfirmPaid(),

		Message_room_info: avro.NewUnionNullMessageRoomInfo(),

		Message_room_waiting_customers: avro.NewUnionNullMessageRoomWaitingCustomers(),

		Error_message: avro.NewUnionNullString(),
	}
}
