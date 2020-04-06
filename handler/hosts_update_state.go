package handler

import (
	"fmt"
	"log"
	"net"
	"runtime/debug"

	"com.lueey.shop/model"
	avro "com.lueey.shop/protocol"
	"com.lueey.shop/utils"
)

type HostsUpdateState struct {
	HandlerSelector
}

func (h *HostsUpdateState) setConn(conn *net.Conn) {
	h.conn = conn
}

func (h *HostsUpdateState) selected(s *model.Session) {
	h.session = s
}

func (h *HostsUpdateState) do(msg avro.Message) {
	log.Printf("##### Will update global state later. #####\n%s", debug.Stack())
	var state int32 = int32(msg.Host_switch_state.RequestHostSwitchState.GlobalState)
	utils.ChangeGlobalState(fmt.Sprintf("Update global state to %v", state))
}
