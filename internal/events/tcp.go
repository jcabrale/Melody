package events

import (
	"strings"
	"time"

	"github.com/bonjourmalware/melody/internal/events/helpers"

	"github.com/bonjourmalware/melody/internal/events/logdata"

	"github.com/bonjourmalware/melody/internal/sessions"

	"github.com/bonjourmalware/melody/internal/config"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type TCPEvent struct {
	LogData logdata.TCPEventLog
	BaseEvent
	helpers.TCPLayer
	helpers.IPv4Layer
	helpers.IPv6Layer
}

func NewTCPEvent(packet gopacket.Packet, IPVersion uint) (*TCPEvent, error) {
	var ev = &TCPEvent{}
	ev.Kind = config.TCPKind
	ev.IPVersion = IPVersion

	ev.Session = sessions.Map.GetUID(packet.TransportLayer().TransportFlow().String())

	switch IPVersion {
	case 4:
		IPHeader, _ := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
		ev.IPv4Layer = helpers.IPv4Layer{Header: IPHeader}
		ev.SourceIP = IPHeader.SrcIP.String()
	case 6:
		IPHeader, _ := packet.Layer(layers.LayerTypeIPv6).(*layers.IPv6)
		ev.IPv6Layer = helpers.IPv6Layer{Header: IPHeader}
		ev.SourceIP = IPHeader.SrcIP.String()

	}

	TCPHeader, _ := packet.Layer(layers.LayerTypeTCP).(*layers.TCP)

	ev.Timestamp = packet.Metadata().Timestamp
	ev.TCPLayer = helpers.TCPLayer{Header: TCPHeader}
	ev.DestPort = uint16(TCPHeader.DstPort)

	ev.Additional = make(map[string]string)
	ev.Tags = make(Tags)

	return ev, nil
}

func (ev TCPEvent) ToLog() EventLog {
	var tcpFlagsStr []string
	//var ipFlagsStr []string

	ev.LogData = logdata.TCPEventLog{}
	ev.LogData.Timestamp = ev.Timestamp.Format(time.RFC3339Nano)
	ev.LogData.Type = ev.Kind
	ev.LogData.SourceIP = ev.SourceIP
	ev.LogData.DestPort = ev.DestPort
	ev.LogData.Session = ev.Session

	// Deduplicate tags
	if len(ev.Tags) == 0 {
		ev.LogData.Tags = []string{}
	} else {
		ev.LogData.Tags = ev.Tags.ToArray()
	}

	switch ev.IPVersion {
	case 4:
		ev.LogData.IP = logdata.NewIPv4LogData(ev.IPv4Layer)
	case 6:
		ev.LogData.IP = logdata.NewIPv6LogData(ev.IPv6Layer)
	}

	ev.LogData.TCP = logdata.TCPLogData{
		Window:     ev.TCPLayer.Header.Window,
		Seq:        ev.TCPLayer.Header.Seq,
		Ack:        ev.TCPLayer.Header.Ack,
		DataOffset: ev.TCPLayer.Header.DataOffset,
		Urgent:     ev.TCPLayer.Header.Urgent,
		Payload:    logdata.NewPayloadLogData(ev.TCPLayer.Header.Payload, config.Cfg.MaxTCPDataSize),
	}

	if ev.TCPLayer.Header.FIN {
		tcpFlagsStr = append(tcpFlagsStr, "F")
	}
	if ev.TCPLayer.Header.SYN {
		tcpFlagsStr = append(tcpFlagsStr, "S")
	}
	if ev.TCPLayer.Header.RST {
		tcpFlagsStr = append(tcpFlagsStr, "R")
	}
	if ev.TCPLayer.Header.PSH {
		tcpFlagsStr = append(tcpFlagsStr, "P")
	}
	if ev.TCPLayer.Header.ACK {
		tcpFlagsStr = append(tcpFlagsStr, "A")
	}
	if ev.TCPLayer.Header.URG {
		tcpFlagsStr = append(tcpFlagsStr, "U")
	}
	if ev.TCPLayer.Header.ECE {
		tcpFlagsStr = append(tcpFlagsStr, "E")
	}
	if ev.TCPLayer.Header.CWR {
		tcpFlagsStr = append(tcpFlagsStr, "C")
	}
	if ev.TCPLayer.Header.NS {
		tcpFlagsStr = append(tcpFlagsStr, "N")
	}

	ev.LogData.TCP.Flags = strings.Join(tcpFlagsStr, "")
	ev.LogData.Additional = ev.Additional

	return ev.LogData
}