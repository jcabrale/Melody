package logger

import (
	"os"

	"gitlab.com/Alvoras/pinknoise/internal/events"
)

type Logger struct {
	File *os.File
}

func (l *Logger) WriteTCPIPEvent(rawData events.TCPEventLog) error {
	data, err := rawData.String()
	if err != nil {
		return err
	}
	_, err = l.File.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return nil
}
func (l *Logger) WriteHTTPEvent(rawData events.HTTPEventLog) error {
	data, err := rawData.String()
	if err != nil {
		return err
	}
	_, err = l.File.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) WriteICMPv4Event(rawData events.ICMPv4EventLog) error {
	data, err := rawData.String()
	if err != nil {
		return err
	}
	_, err = l.File.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return nil
}
