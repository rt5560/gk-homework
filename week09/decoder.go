package main

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

const (
	PacketLen       = 4
	HeaderLen       = 2
	ProtocolVersion = 2
	Operation       = 4
	SequenceId      = 4
)

func decoder(data []byte) (map[string]interface{}, error) {
	AllLen := PacketLen + HeaderLen + ProtocolVersion + Operation + SequenceId
	if len(data) <= AllLen {
		return nil, errors.New("error")
	}
	decoderData := map[string]interface{}{}
	// PacketLen
	packetLen := binary.BigEndian.Uint32(data[:PacketLen])
	decoderData["packetLen"] = packetLen
	// HeaderLen
	HeaderLenStart := PacketLen
	HeaderLenEnd := PacketLen + HeaderLen
	headerLen := binary.BigEndian.Uint16(data[HeaderLenStart:HeaderLenEnd])
	decoderData["headerLen"] = headerLen
	// ProtocolVersion
	versionEnd := HeaderLenEnd + ProtocolVersion
	version := binary.BigEndian.Uint16(data[HeaderLenEnd:versionEnd])
	decoderData["version"] = version
	// Operation
	operationEnd := versionEnd + Operation
	operation := binary.BigEndian.Uint32(data[versionEnd:operationEnd])
	decoderData["operation"] = operation
	// SequenceId
	sequenceEnd := operationEnd + SequenceId
	sequence := binary.BigEndian.Uint32(data[versionEnd:sequenceEnd])
	decoderData["sequence"] = sequence

	body := string(data[AllLen:])
	decoderData["body"] = body
	return decoderData, nil
}
