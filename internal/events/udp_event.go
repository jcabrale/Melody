package events

import (
	"strconv"
	"strings"
	"time"

	"github.com/bonjourmalware/pinknoise/internal/config"

	"github.com/bonjourmalware/pinknoise/internal/sessions"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type UDPEvent struct {
	//IPHeader  *layers.IPv4
	//UDPHeader *layers.UDP
	LogData   UDPEventLog
	BaseEvent
	UDPLayer
	IPLayer
}

func NewUDPEvent(packet gopacket.Packet) (*UDPEvent, error) {
	var ev = &UDPEvent{}
	ev.Kind = UDPKind

	ev.Session = sessions.Map.GetUID(packet.TransportLayer().TransportFlow().String())

	//IPHeader, _ := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
	//ev.IPHeader = IPHeader
	IPHeader, _ := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
	ev.IPLayer = IPLayer{Header: IPHeader}
	ev.SourceIP = IPHeader.SrcIP.String()

	UDPHeader, _ := packet.Layer(layers.LayerTypeUDP).(*layers.UDP)
	ev.UDPLayer = UDPLayer{Header: UDPHeader}
	ev.DestPort = uint(UDPHeader.DstPort)

	ev.Metadata = make(map[string]string)
	ev.References = make(map[string][]string)
	ev.Statements = []string{}

	return ev, nil
}

func (ev UDPEvent) ToLog() EventLog {
	var ipFlagsStr []string

	ev.LogData = UDPEventLog{}
	ev.LogData.Timestamp = time.Now().Format(time.RFC3339)
	ev.LogData.NsTimestamp = strconv.FormatInt(time.Now().UnixNano(), 10)
	ev.LogData.Type = ev.Kind
	ev.LogData.SourceIP = ev.SourceIP
	ev.LogData.DestPort = ev.DestPort
	ev.LogData.Session = ev.Session

	// Deduplicate tags
	if len(ev.Tags) == 0 {
		ev.LogData.Tags = []string{}
	} else {
		var set = make(map[string]struct{})
		for _, tag := range ev.Tags {
			if _, ok := set[tag]; !ok {
				set[tag] = struct{}{}
			}
		}

		for tag, _ := range set {
			ev.LogData.Tags = append(ev.LogData.Tags, tag)
		}
	}

	ev.LogData.IP = IPLogData{
		Version:    ev.IPLayer.Header.Version,
		IHL:        ev.IPLayer.Header.IHL,
		TOS:        ev.IPLayer.Header.TOS,
		Length:     ev.IPLayer.Header.Length,
		Id:         ev.IPLayer.Header.Id,
		FragOffset: ev.IPLayer.Header.FragOffset,
		TTL:        ev.IPLayer.Header.TTL,
		Protocol:   ev.IPLayer.Header.Protocol,
	}

	if ev.IPLayer.Header.Flags&layers.IPv4EvilBit != 0 {
		ipFlagsStr = append(ipFlagsStr, "EV")
	}
	if ev.IPLayer.Header.Flags&layers.IPv4DontFragment != 0 {
		ipFlagsStr = append(ipFlagsStr, "DF")
	}
	if ev.IPLayer.Header.Flags&layers.IPv4MoreFragments != 0 {
		ipFlagsStr = append(ipFlagsStr, "MF")
	}

	ev.LogData.IP.Fragbits = strings.Join(ipFlagsStr, "")

	ev.LogData.UDP = UDPLogData{
		Payload:  NewPayload(ev.UDPLayer.Header.Payload, config.Cfg.MaxUDPDataSize),
		Length:   ev.UDPLayer.Header.Length,
		Checksum: ev.UDPLayer.Header.Checksum,
	}

	ev.LogData.Metadata = ev.Metadata
	ev.LogData.References = ev.References
	ev.LogData.Statements = ev.Statements

	return ev.LogData
}
