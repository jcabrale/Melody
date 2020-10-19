package events

import (
	"encoding/base64"
	"encoding/json"

	"github.com/google/gopacket/layers"
)

//type LogDataIface interface {
//	String() (string, error)
//}

type Payload struct {
	Content   string `json:"content"`
	Base64    string `json:"base64"`
	Truncated bool   `json:"truncated"`
}

type BaseLogData struct {
	NsTimestamp string              `json:"ns_timestamp"`
	Timestamp   string              `json:"timestamp"`
	Session     string              `json:"session"`
	Type        string              `json:"type"`
	SourceIP    string              `json:"src_ip"`
	DestPort    uint                `json:"dst_port"`
	Tags        []string            `json:"pk_tags"`
	Metadata    map[string]string   `json:"metadata"`
	References  map[string][]string `json:"references"`
	Statements  []string            `json:"statements"`
}

type ICMPv4EventLog struct {
	ICMPv4 ICMPv4LogData `json:"icmpv4"`
	IP     IPLogData     `json:"ip"`
	BaseLogData
}

type TCPEventLog struct {
	TCP TCPLogData `json:"tcp"`
	IP  IPLogData  `json:"ip"`
	BaseLogData
}

type UDPEventLog struct {
	UDP UDPLogData `json:"udp"`
	IP  IPLogData  `json:"ip"`
	BaseLogData
}

type HTTPEventLog struct {
	HTTP HTTPLogData `json:"http"`
	BaseLogData
}

type ICMPv4LogData struct {
	TypeCode layers.ICMPv4TypeCode `json:"type_code"`
	Checksum uint16                `json:"checksum"`
	Id       uint16                `json:"id"`
	Seq      uint16                `json:"seq"`
}

type TCPLogData struct {
	Window     uint16  `json:"window"`
	Seq        uint32  `json:"seq"`
	Ack        uint32  `json:"ack"`
	DataOffset uint8   `json:"data_offset"`
	Flags      string  `json:"flags"`
	Urgent     uint16  `json:"urgent"`
	Payload    Payload `json:"payload"`
}

type UDPLogData struct {
	Payload  Payload `json:"payload"`
	Length   uint16  `json:"length"`
	Checksum uint16  `json:"checksum"`
}

type IPLogData struct {
	Version    uint8             `json:"version"`
	IHL        uint8             `json:"ihl"`
	TOS        uint8             `json:"tos"`
	Length     uint16            `json:"length"`
	Id         uint16            `json:"id"`
	Fragbits   string            `json:"fragbits"`
	FragOffset uint16            `json:"frag_offset"`
	TTL        uint8             `json:"ttl"`
	Protocol   layers.IPProtocol `json:"protocol"`
}

type HTTPLogData struct {
	Verb       string            `json:"verb"`
	Proto      string            `json:"proto"`
	RequestURI string            `json:"URI"`
	RemoteAddr string            `json:"remote_address"`
	SourcePort uint64            `json:"src_port"`
	DestHost   string            `json:"dst_host"`
	Headers    map[string]string `json:"headers"`
	Errors     []string          `json:"errors"`
	Body       Payload           `json:"body"`
	IsTLS      bool              `json:"is_tls"`
}

func (eventLog TCPEventLog) String() (string, error) {
	data, err := json.Marshal(eventLog)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (eventLog HTTPEventLog) String() (string, error) {
	data, err := json.Marshal(eventLog)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (eventLog ICMPv4EventLog) String() (string, error) {
	data, err := json.Marshal(eventLog)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (eventLog UDPEventLog) String() (string, error) {
	data, err := json.Marshal(eventLog)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewPayload(data []byte, maxLength uint64) Payload {
	var pl = Payload{}

	if uint64(len(data)) > maxLength {
		data = data[:maxLength]
		pl.Truncated = true
	}
	pl.Content = string(data)
	pl.Base64 = base64.StdEncoding.EncodeToString(data)
	return pl
}
