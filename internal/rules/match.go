package rules

import (
	"github.com/bonjourmalware/pinknoise/internal/events"
)

func (rule *Rule) Match(ev events.Event) bool {
	var condOK bool
	//ipHeader := ev.GetIPHeader()

	// The rule fails if the source IP is blacklisted
	if len(rule.IPs.BlacklistedIPs) > 0 {
		for _, iprange := range rule.IPs.BlacklistedIPs {
			if iprange.ContainsIPString(ev.GetSourceIP()) {
				return false
			}
		}
	}

	// The rule fails if the source IP is not in the whitelisted addresses
	if len(rule.IPs.WhitelistedIPs) > 0 {
		condOK = false

		for _, iprange := range rule.IPs.WhitelistedIPs {
			if iprange.ContainsIPString(ev.GetSourceIP()) {
				condOK = true
				break
			}
		}

		if !condOK {
			return false
		}
	}

	//if ipHeader != nil {
	//	//if rule.ID != nil {
	//	//	if ev.IPHeader.Id != *rule.ID {
	//	//		return false
	//	//	}
	//	//}
	//
	//	if rule.Options.MatchAll {
	//		if rule.TTL != nil {
	//			if ipHeader.TTL != *rule.TTL {
	//				return false
	//			}
	//		}
	//
	//		if rule.TOS != nil {
	//			if ipHeader.TOS != *rule.TOS {
	//				return false
	//			}
	//		}
	//
	//		if len(rule.Fragbits) > 0 {
	//			condOK = false
	//			for _, fragbits := range rule.Fragbits {
	//				// If at least one of the flag string match exactly, then continue
	//				if uint8(ipHeader.Flags)^(*fragbits) == 0 {
	//					condOK = true
	//					break
	//				}
	//			}
	//			if !condOK {
	//				return false
	//			}
	//		}
	//
	//		return true
	//	}
	//
	//	if rule.TTL != nil {
	//		if ipHeader.TTL == *rule.TTL {
	//			return true
	//		}
	//	}
	//
	//	if rule.TOS != nil {
	//		if ipHeader.TOS == *rule.TOS {
	//			return true
	//		}
	//	}
	//
	//	if len(rule.Fragbits) > 0 {
	//		for _, fragbits := range rule.Fragbits {
	//			// If at least one of the flag string match exactly, then continue
	//			if uint8(ipHeader.Flags)^(*fragbits) == 0 {
	//				return true
	//			}
	//		}
	//	}
	//}

	switch ev.GetKind() {
	case events.UDPKind:
		return rule.MatchUDPEvent(ev)
	case events.TCPKind:
		return rule.MatchTCPEvent(ev)
	case events.ICMPv4Kind:
		return rule.MatchICMPv4Event(ev)
	case events.HTTPKind:
		return rule.MatchHTTPEvent(ev)
	}

	return false
}

func (rule *Rule) MatchICMPv4Event(ev events.Event) bool {
	//var condOK bool

	//if rule.ID != nil {
	//	if ev.IPHeader.Id != *rule.ID {
	//		return false
	//	}
	//}
	// The rule fails if the source IP is blacklisted
	//if len(rule.IPs.BlacklistedIPs) > 0 {
	//	for _, iprange := range rule.IPs.BlacklistedIPs {
	//		if iprange.ContainsIPString(ev.SourceIP) {
	//			return false
	//		}
	//	}
	//}
	//
	//// The rule fails if the source IP is not in the whitelisted addresses
	//if len(rule.IPs.WhitelistedIPs) > 0 {
	//	condOK = false
	//
	//	for _, iprange := range rule.IPs.WhitelistedIPs {
	//		if iprange.ContainsIPString(ev.SourceIP) {
	//			condOK = true
	//			break
	//		}
	//	}
	//
	//	if !condOK {
	//		return false
	//	}
	//}
	//
	//if rule.Options.MatchAll {
	//	if rule.TTL != nil {
	//		if ev.IPHeader.TTL != *rule.TTL {
	//			return false
	//		}
	//	}
	//
	//	if rule.TOS != nil {
	//		if ev.IPHeader.TOS != *rule.TOS {
	//			return false
	//		}
	//	}
	//
	//	return true
	//}
	//
	//if rule.TTL != nil {
	//	if ev.IPHeader.TTL == *rule.TTL {
	//		return true
	//	}
	//}
	//
	//if rule.TOS != nil {
	//	if ev.IPHeader.TOS == *rule.TOS {
	//		return true
	//	}
	//}

	return false
}

func (rule *Rule) MatchUDPEvent(ev events.Event) bool {
	//var condOK bool

	//
	//// The rule fails if the source IP is blacklisted
	//if len(rule.IPs.BlacklistedIPs) > 0 {
	//	for _, iprange := range rule.IPs.BlacklistedIPs {
	//		if iprange.ContainsIPString(ev.SourceIP) {
	//			return false
	//		}
	//	}
	//}
	//
	//// The rule fails if the source IP is not in the whitelisted addresses
	//if len(rule.IPs.WhitelistedIPs) > 0 {
	//	condOK = false
	//
	//	for _, iprange := range rule.IPs.WhitelistedIPs {
	//		if iprange.ContainsIPString(ev.SourceIP) {
	//			condOK = true
	//			break
	//		}
	//	}
	//
	//	if !condOK {
	//		return false
	//	}
	//}
	//
	//if rule.Options.MatchAll {
	//	if rule.TTL != nil {
	//		if ev.IPHeader.TTL != *rule.TTL {
	//			return false
	//		}
	//	}
	//
	//	if rule.TOS != nil {
	//		if ev.IPHeader.TOS != *rule.TOS {
	//			return false
	//		}
	//	}
	//
	//	//TODO : Add <, > and <> operators
	//	if rule.UDPLength != nil {
	//		if ev.UDPHeader.Length != *rule.UDPLength {
	//			return false
	//		}
	//	}
	//
	//	if rule.Checksum != nil {
	//		if ev.UDPHeader.Checksum != *rule.Checksum {
	//			return false
	//		}
	//	}
	//
	//	if rule.Payload != nil {
	//		if !rule.Payload.Match(ev.UDPHeader.Payload, rule.Options) {
	//			return false
	//		}
	//	}
	//
	//	return true
	//}

	//if rule.TTL != nil {
	//	if ev.IPHeader.TTL == *rule.TTL {
	//		return true
	//	}
	//}
	//
	//if rule.TOS != nil {
	//	if ev.IPHeader.TOS == *rule.TOS {
	//		return true
	//	}
	//}

	udpHeader := ev.GetUDPHeader()

	if len(rule.Ports) > 0 {
		for _, port := range rule.Ports {
			// If at least one port is valid
			if port == uint(udpHeader.DstPort) {
				break
			}
		}
	}

	if rule.Options.MatchAll {
		//TODO : Add <, > and <> operators
		if rule.UDPLength != nil {
			if udpHeader.Length != *rule.UDPLength {
				return false
			}
		}

		if rule.Checksum != nil {
			if udpHeader.Checksum != *rule.Checksum {
				return false
			}
		}

		if rule.Payload != nil {
			if !rule.Payload.Match(udpHeader.Payload, rule.Options) {
				return false
			}
		}

		return true
	}

	//TODO : Add <, > and <> operators
	if rule.UDPLength != nil {
		if udpHeader.Length == *rule.UDPLength {
			return true
		}
	}

	if rule.Checksum != nil {
		if udpHeader.Checksum == *rule.Checksum {
			return true
		}
	}

	if rule.Payload != nil {
		if rule.Payload.Match(udpHeader.Payload, rule.Options) {
			return true
		}
	}

	return false
}

func (rule *Rule) MatchTCPEvent(ev events.Event) bool {
	tcpHeader := ev.GetTCPHeader()

	var condOK bool

	if len(rule.Ports) > 0 {
		for _, port := range rule.Ports {
			// If at least one port is valid
			if port == uint(tcpHeader.DstPort) {
				break
			}
		}
	}

	//// The rule fails if the source IP is blacklisted
	//if len(rule.IPs.BlacklistedIPs) > 0 {
	//	for _, iprange := range rule.IPs.BlacklistedIPs {
	//		if iprange.ContainsIPString(ev.SourceIP) {
	//			return false
	//		}
	//	}
	//}
	//
	//// The rule fails if the source IP is not in the whitelisted addresses
	//if len(rule.IPs.WhitelistedIPs) > 0 {
	//	condOK = false
	//
	//	for _, iprange := range rule.IPs.WhitelistedIPs {
	//		if iprange.ContainsIPString(ev.SourceIP) {
	//			condOK = true
	//			break
	//		}
	//	}
	//
	//	if !condOK {
	//		return false
	//	}
	//}

	if rule.Options.MatchAll {
		if len(rule.Flags) > 0 {
			condOK = false
			for _, flags := range rule.Flags {
				// If at least one of the flag string match exactly, then continue
				if tcpHeader.BaseLayer.Contents[13]^(*flags) == 0 {
					condOK = true
					break
				}
			}
			if !condOK {
				return false
			}
		}

		if rule.Seq != nil {
			if tcpHeader.Seq != *rule.Seq {
				return false
			}
		}

		if rule.Ack != nil {
			if tcpHeader.Ack != *rule.Ack {
				return false
			}
		}

		if rule.Window != nil {
			if tcpHeader.Window != *rule.Window {
				return false
			}
		}

		if rule.Payload != nil {
			if !rule.Payload.Match(tcpHeader.Payload, rule.Options) {
				return false
			}
		}

		//TODO : Add <, > and <> operators
		if rule.Dsize != nil {
			if len(tcpHeader.Payload) != *rule.Dsize {
				return false
			}
		}

		return true
	}

	// else
	if len(rule.Flags) > 0 {
		for _, flags := range rule.Flags {
			// If at least one of the flag string match exactly, then continue
			if tcpHeader.BaseLayer.Contents[13]^(*flags) == 0 {
				return true
			}
		}
	}

	if rule.Seq != nil {
		if tcpHeader.Seq == *rule.Seq {
			return true
		}
	}

	if rule.Ack != nil {
		if tcpHeader.Ack == *rule.Ack {
			return true
		}
	}

	if rule.Window != nil {
		if tcpHeader.Window == *rule.Window {
			return true
		}
	}

	if rule.Payload != nil {
		if rule.Payload.Match(tcpHeader.Payload, rule.Options) {
			return true
		}
	}

	//TODO : Add <, > and <> operators
	if rule.Dsize != nil {
		if len(tcpHeader.Payload) == *rule.Dsize {
			return true
		}
	}

	return false
}

func (rule *Rule) MatchHTTPEvent(ev events.Event) bool {
	httpData := ev.GetHTTPData()

	var condOK bool

	if len(rule.Ports) > 0 {
		var portMatch bool
		for _, port := range rule.Ports {
			// If at least one port is valid
			if port == httpData.DestPort {
				portMatch = true
				break
			}
		}

		if portMatch == false {
			return false
		}
	}

	if rule.Options.MatchAll {
		if rule.URI != nil {
			if rule.URI.Match([]byte(httpData.RequestURI), rule.Options) == false {
				return false
			}
		}

		if rule.Body != nil {
			if rule.Body.Match([]byte(httpData.Body.Content), rule.Options) == false {
				return false
			}
		}

		if rule.Headers != nil {
			condOK = false

			for _, inlineHeader := range httpData.InlineHeaders {
				if rule.Headers.Match([]byte(inlineHeader), rule.Options) == true {
					condOK = true
					break
				}
			}

			if !condOK {
				return false
			}
		}

		if rule.Verb != nil {
			if rule.Verb.Match([]byte(httpData.Verb), rule.Options) == false {
				return false
			}
		}

		if rule.Proto != nil {
			if rule.Proto.Match([]byte(httpData.Proto), rule.Options) == false {
				return false
			}
		}
	}

	// else

	if rule.URI != nil {
		if rule.URI.Match([]byte(httpData.RequestURI), rule.Options) == true {
			return true
		}
	}

	if rule.Body != nil {
		if rule.Body.Match([]byte(httpData.Body.Content), rule.Options) == true {
			return true
		}
	}

	if rule.Headers != nil {
		condOK = false

		for _, inlineHeader := range httpData.InlineHeaders {
			if rule.Headers.Match([]byte(inlineHeader), rule.Options) == true {
				condOK = true
				break
			}
		}

		if !condOK {
			return false
		}
	}

	if rule.Verb != nil {
		if rule.Verb.Match([]byte(httpData.Verb), rule.Options) == true {
			return true
		}
	}

	if rule.Proto != nil {
		if rule.Proto.Match([]byte(httpData.Proto), rule.Options) == true {
			return true
		}
	}

	return true
}
