// generated by stringer -type=connectionState; DO NOT EDIT

package tchannel

import "fmt"

const _connectionState_name = "connectionWaitingToRecvInitReqconnectionWaitingToSendInitReqconnectionWaitingToRecvInitResconnectionActiveconnectionStartCloseconnectionInboundClosedconnectionClosed"

var _connectionState_index = [...]uint8{0, 30, 60, 90, 106, 126, 149, 165}

func (i connectionState) String() string {
	i -= 1
	if i < 0 || i+1 >= connectionState(len(_connectionState_index)) {
		return fmt.Sprintf("connectionState(%d)", i+1)
	}
	return _connectionState_name[_connectionState_index[i]:_connectionState_index[i+1]]
}
