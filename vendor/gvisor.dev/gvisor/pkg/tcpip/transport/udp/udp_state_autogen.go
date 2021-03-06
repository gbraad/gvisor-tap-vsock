// automatically generated by stateify.

package udp

import (
	"gvisor.dev/gvisor/pkg/state"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
)

func (u *udpPacket) StateTypeName() string {
	return "pkg/tcpip/transport/udp.udpPacket"
}

func (u *udpPacket) StateFields() []string {
	return []string{
		"udpPacketEntry",
		"senderAddress",
		"destinationAddress",
		"packetInfo",
		"data",
		"timestamp",
		"tos",
	}
}

func (u *udpPacket) beforeSave() {}

func (u *udpPacket) StateSave(stateSinkObject state.Sink) {
	u.beforeSave()
	var dataValue buffer.VectorisedView = u.saveData()
	stateSinkObject.SaveValue(4, dataValue)
	stateSinkObject.Save(0, &u.udpPacketEntry)
	stateSinkObject.Save(1, &u.senderAddress)
	stateSinkObject.Save(2, &u.destinationAddress)
	stateSinkObject.Save(3, &u.packetInfo)
	stateSinkObject.Save(5, &u.timestamp)
	stateSinkObject.Save(6, &u.tos)
}

func (u *udpPacket) afterLoad() {}

func (u *udpPacket) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &u.udpPacketEntry)
	stateSourceObject.Load(1, &u.senderAddress)
	stateSourceObject.Load(2, &u.destinationAddress)
	stateSourceObject.Load(3, &u.packetInfo)
	stateSourceObject.Load(5, &u.timestamp)
	stateSourceObject.Load(6, &u.tos)
	stateSourceObject.LoadValue(4, new(buffer.VectorisedView), func(y interface{}) { u.loadData(y.(buffer.VectorisedView)) })
}

func (e *endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/udp.endpoint"
}

func (e *endpoint) StateFields() []string {
	return []string{
		"TransportEndpointInfo",
		"DefaultSocketOptionsHandler",
		"waiterQueue",
		"uniqueID",
		"rcvReady",
		"rcvList",
		"rcvBufSizeMax",
		"rcvBufSize",
		"rcvClosed",
		"sndBufSize",
		"sndBufSizeMax",
		"state",
		"dstPort",
		"ttl",
		"multicastTTL",
		"multicastAddr",
		"multicastNICID",
		"portFlags",
		"lastError",
		"boundBindToDevice",
		"boundPortFlags",
		"sendTOS",
		"shutdownFlags",
		"multicastMemberships",
		"effectiveNetProtos",
		"owner",
		"ops",
	}
}

func (e *endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	var rcvBufSizeMaxValue int = e.saveRcvBufSizeMax()
	stateSinkObject.SaveValue(6, rcvBufSizeMaxValue)
	var lastErrorValue string = e.saveLastError()
	stateSinkObject.SaveValue(18, lastErrorValue)
	stateSinkObject.Save(0, &e.TransportEndpointInfo)
	stateSinkObject.Save(1, &e.DefaultSocketOptionsHandler)
	stateSinkObject.Save(2, &e.waiterQueue)
	stateSinkObject.Save(3, &e.uniqueID)
	stateSinkObject.Save(4, &e.rcvReady)
	stateSinkObject.Save(5, &e.rcvList)
	stateSinkObject.Save(7, &e.rcvBufSize)
	stateSinkObject.Save(8, &e.rcvClosed)
	stateSinkObject.Save(9, &e.sndBufSize)
	stateSinkObject.Save(10, &e.sndBufSizeMax)
	stateSinkObject.Save(11, &e.state)
	stateSinkObject.Save(12, &e.dstPort)
	stateSinkObject.Save(13, &e.ttl)
	stateSinkObject.Save(14, &e.multicastTTL)
	stateSinkObject.Save(15, &e.multicastAddr)
	stateSinkObject.Save(16, &e.multicastNICID)
	stateSinkObject.Save(17, &e.portFlags)
	stateSinkObject.Save(19, &e.boundBindToDevice)
	stateSinkObject.Save(20, &e.boundPortFlags)
	stateSinkObject.Save(21, &e.sendTOS)
	stateSinkObject.Save(22, &e.shutdownFlags)
	stateSinkObject.Save(23, &e.multicastMemberships)
	stateSinkObject.Save(24, &e.effectiveNetProtos)
	stateSinkObject.Save(25, &e.owner)
	stateSinkObject.Save(26, &e.ops)
}

func (e *endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.TransportEndpointInfo)
	stateSourceObject.Load(1, &e.DefaultSocketOptionsHandler)
	stateSourceObject.Load(2, &e.waiterQueue)
	stateSourceObject.Load(3, &e.uniqueID)
	stateSourceObject.Load(4, &e.rcvReady)
	stateSourceObject.Load(5, &e.rcvList)
	stateSourceObject.Load(7, &e.rcvBufSize)
	stateSourceObject.Load(8, &e.rcvClosed)
	stateSourceObject.Load(9, &e.sndBufSize)
	stateSourceObject.Load(10, &e.sndBufSizeMax)
	stateSourceObject.Load(11, &e.state)
	stateSourceObject.Load(12, &e.dstPort)
	stateSourceObject.Load(13, &e.ttl)
	stateSourceObject.Load(14, &e.multicastTTL)
	stateSourceObject.Load(15, &e.multicastAddr)
	stateSourceObject.Load(16, &e.multicastNICID)
	stateSourceObject.Load(17, &e.portFlags)
	stateSourceObject.Load(19, &e.boundBindToDevice)
	stateSourceObject.Load(20, &e.boundPortFlags)
	stateSourceObject.Load(21, &e.sendTOS)
	stateSourceObject.Load(22, &e.shutdownFlags)
	stateSourceObject.Load(23, &e.multicastMemberships)
	stateSourceObject.Load(24, &e.effectiveNetProtos)
	stateSourceObject.Load(25, &e.owner)
	stateSourceObject.Load(26, &e.ops)
	stateSourceObject.LoadValue(6, new(int), func(y interface{}) { e.loadRcvBufSizeMax(y.(int)) })
	stateSourceObject.LoadValue(18, new(string), func(y interface{}) { e.loadLastError(y.(string)) })
	stateSourceObject.AfterLoad(e.afterLoad)
}

func (m *multicastMembership) StateTypeName() string {
	return "pkg/tcpip/transport/udp.multicastMembership"
}

func (m *multicastMembership) StateFields() []string {
	return []string{
		"nicID",
		"multicastAddr",
	}
}

func (m *multicastMembership) beforeSave() {}

func (m *multicastMembership) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.nicID)
	stateSinkObject.Save(1, &m.multicastAddr)
}

func (m *multicastMembership) afterLoad() {}

func (m *multicastMembership) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.nicID)
	stateSourceObject.Load(1, &m.multicastAddr)
}

func (l *udpPacketList) StateTypeName() string {
	return "pkg/tcpip/transport/udp.udpPacketList"
}

func (l *udpPacketList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *udpPacketList) beforeSave() {}

func (l *udpPacketList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *udpPacketList) afterLoad() {}

func (l *udpPacketList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *udpPacketEntry) StateTypeName() string {
	return "pkg/tcpip/transport/udp.udpPacketEntry"
}

func (e *udpPacketEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *udpPacketEntry) beforeSave() {}

func (e *udpPacketEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *udpPacketEntry) afterLoad() {}

func (e *udpPacketEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*udpPacket)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*multicastMembership)(nil))
	state.Register((*udpPacketList)(nil))
	state.Register((*udpPacketEntry)(nil))
}
