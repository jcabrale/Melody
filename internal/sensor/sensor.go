package sensor

import (
	"fmt"
	"time"

	"gitlab.com/Alvoras/pinknoise/internal/engine"
	"gitlab.com/Alvoras/pinknoise/internal/events"

	"gitlab.com/Alvoras/pinknoise/internal/sessions"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/tcpassembly"
	"gitlab.com/Alvoras/pinknoise/internal/config"
	"gitlab.com/Alvoras/pinknoise/internal/http_assembler"
)

func Start(quitErrChan chan error, shutdownChan chan bool, sensorStoppedChan chan bool) {
	go ReceivePackets(quitErrChan, shutdownChan, sensorStoppedChan)
}

func ReceivePackets(quitErrChan chan error, shutdownChan chan bool, sensorStoppedChan chan bool) {
	// Set up HTTP assembly
	streamFactory := &http_assembler.HttpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)
	var handle *pcap.Handle
	var err error

	if config.Cfg.PcapFile != nil {
		handle, err = pcap.OpenOfflineFile(config.Cfg.PcapFile)
		if err != nil {
			quitErrChan <- err
			close(sensorStoppedChan)
		}
	} else {
		// Open up a pcap handle for packet reads/writes.
		handle, err = pcap.OpenLive(config.Cfg.Interface, 65536, true, pcap.BlockForever)
		if err != nil {
			quitErrChan <- err
			close(sensorStoppedChan)
		}
	}

	defer handle.Close()
	if config.Cfg.BPFFilter != "" {
		if err := handle.SetBPFFilter(config.Cfg.BPFFilter); err != nil {
			quitErrChan <- err
			return
		}
	}

	assemblerFlushTicker := time.Tick(time.Minute)
	sessionsFlushTicker := time.Tick(time.Second * 30)
	src := gopacket.NewPacketSource(handle, handle.LinkType())
	in := src.Packets()

	defer func() {
		assembler.FlushAll()
		sessions.Map.FlushAll()
		close(sensorStoppedChan)
	}()

loop:
	for {
		var packet gopacket.Packet
		select {
		case packet = <-in:
			if packet == nil {
				break loop
			}

			if packet.NetworkLayer() != nil {
				if _, ok := packet.NetworkLayer().(*layers.IPv4); ok {
					// Ignore outgoing packets
					for _, ip := range config.Cfg.HomeNet {
						if packet.NetworkLayer().(*layers.IPv4).SrcIP.String() == ip {
							continue loop
						}
					}
					switch packet.NetworkLayer().(*layers.IPv4).Protocol {
					case layers.IPProtocolICMPv4:
						event, err := events.NewICMPv4Event(packet)
						if err != nil {
							//TODO: write to error log
							fmt.Println("ERROR", err)
							continue
						}

						engine.ICMPv4EventChan <- event

					case layers.IPProtocolTCP:
						event, err := events.NewTCPEvent(packet)
						if err != nil {
							//TODO: write to error log
							fmt.Println("ERROR", err)
							continue
						}

						engine.TCPEventChan <- event

						tcpPacket := packet.TransportLayer().(*layers.TCP)
						assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcpPacket, packet.Metadata().Timestamp)

					default:
						continue loop
					}
				}
			}

		case <-assemblerFlushTicker:
			// Every minute, flush connections that haven't seen activity in the past 2 minutes
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		case <-sessionsFlushTicker:
			// Every 30 seconds, flush inactive flows
			sessions.Map.FlushOlderThan(time.Now().Add(time.Second * -30))
		case <-shutdownChan:
			return
		}
	}

	close(shutdownChan)
}