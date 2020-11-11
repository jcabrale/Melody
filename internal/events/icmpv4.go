package events

import (
	"time"

	"github.com/bonjourmalware/melody/internal/events/helpers"
	"github.com/bonjourmalware/melody/internal/events/logdata"

	"github.com/bonjourmalware/melody/internal/config"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type ICMPv4Event struct {
	//ICMPv4Header *layers.ICMPv4
	LogData logdata.ICMPv4EventLog
	BaseEvent
	helpers.IPv4Layer
	helpers.ICMPv4Layer
}

func NewICMPv4Event(packet gopacket.Packet) (*ICMPv4Event, error) {
	var ev = &ICMPv4Event{}
	ev.Kind = config.ICMPv4Kind

	ev.Session = "n/a"
	ev.Timestamp = packet.Metadata().Timestamp

	ICMPv4Header, _ := packet.Layer(layers.LayerTypeICMPv4).(*layers.ICMPv4)
	ev.ICMPv4Layer = helpers.ICMPv4Layer{Header: ICMPv4Header}

	IPHeader, _ := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
	ev.IPv4Layer = helpers.IPv4Layer{Header: IPHeader}
	ev.SourceIP = ev.IPv4Layer.Header.SrcIP.String()
	ev.Additional = make(map[string]string)
	ev.Tags = make(Tags)

	return ev, nil
}

func (ev ICMPv4Event) ToLog() EventLog {
	ev.LogData = logdata.ICMPv4EventLog{}
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

	ev.LogData.ICMPv4 = logdata.ICMPv4LogData{
		TypeCode:     ev.ICMPv4Layer.Header.TypeCode,
		Type:         ev.ICMPv4Layer.Header.TypeCode.Type(),
		Code:         ev.ICMPv4Layer.Header.TypeCode.Code(),
		TypeCodeName: ev.ICMPv4Layer.Header.TypeCode.String(),
		Checksum:     ev.ICMPv4Layer.Header.Checksum,
		Id:           ev.ICMPv4Layer.Header.Id,
		Seq:          ev.ICMPv4Layer.Header.Seq,
	}

	ev.LogData.IP = logdata.NewIPv4LogData(ev.IPv4Layer)
	ev.LogData.Additional = ev.Additional

	return ev.LogData
}