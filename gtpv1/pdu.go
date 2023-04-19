package gtpv1

import (
	"encoding/binary"
	"fmt"
)

// MessageType represents possible GTPv1 message type values
type MessageType uint8

// GTPv1 MessageTypes
const (
	EchoRequest                           MessageType = 1
	EchoResponse                          MessageType = 2
	VersionNotSupported                   MessageType = 3
	NodeAliveRequest                      MessageType = 4
	NodeAliveResponse                     MessageType = 5
	RedirectionRequest                    MessageType = 6
	RedirectionResponse                   MessageType = 7
	PDUNotificationRejectResponse         MessageType = 30
	SupportedExtensionHeadersNotification MessageType = 31
	SendRouteingInformationforGPRSRequest MessageType = 32
	SendRouteingInformationforGPRS        MessageType = 33
	FailureReportRequest                  MessageType = 34
	FailureReportResponse                 MessageType = 35
	NoteMSGPRSPresentRequest              MessageType = 36
	NoteMSGPRSPresentResponse             MessageType = 37
	IdentificationRequest                 MessageType = 48
	IdentificationResponse                MessageType = 49
	SGSNContextRequest                    MessageType = 50
	SGSNContextResponse                   MessageType = 51
	SGSNContextAcknowledge                MessageType = 52
	ForwardRelocationRequest              MessageType = 53
	ForwardRelocationResponse             MessageType = 54
	ForwardRelocationComplete             MessageType = 55
	RelocationCancelRequest               MessageType = 56
	RelocationCancelResponse              MessageType = 57
	ForwardSRNSContext                    MessageType = 58
	ForwardRelocationCompleteAcknowledge  MessageType = 59
	ForwardSRNSContextAcknowledge         MessageType = 60
	MBMSNotificationRequest               MessageType = 96
	MBMSNotificationResponse              MessageType = 97
	MBMSNotificationRejectRequest         MessageType = 98
	MBMSNotificationRejectResponse        MessageType = 99
	CreateMBMSContextRequest              MessageType = 100
	CreateMBMSContextResponse             MessageType = 101
	UpdateMBMSContextRequest              MessageType = 102
	UpdateMBMSContextResponse             MessageType = 103
	DeleteMBMSContextRequest              MessageType = 104
	DeleteMBMSContextResponse             MessageType = 105
	MBMSRegistrationRequest               MessageType = 112
	MBMSRegistrationResponse              MessageType = 113
	MBMSDeRegistrationRequest             MessageType = 114
	MBMSDeRegistrationResponse            MessageType = 115
	MBMSSessionStartRequest               MessageType = 116
	MBMSSessionStartResponse              MessageType = 117
	MBMSSessionStopRequest                MessageType = 118
	MBMSSessionStopResponse               MessageType = 119
	MBMSSessionUpdateRequest              MessageType = 120
	MBMSSessionUpdateResponse             MessageType = 121
	MSInfoChangeNotificationRequest       MessageType = 128
	MSInfoChangeNotificationResponse      MessageType = 129
	DataRecordTransferRequest             MessageType = 240
	DataRecordTransferResponse            MessageType = 241
	EndMarker                             MessageType = 254
	GPDU                                  MessageType = 255
)

var messageNames = []string{
	"Reserved",
	"Reserved", "Echo Request", "Echo Response", "Version Not Supported", "Node Alive Request", "Node Alive Response", // 5
	"Redirection Request", "Redirection Response", "Reserved", "Reserved", "Reserved", // 10
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 15
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 20
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 25
	"Reserved", "Reserved", "Reserved", "Reserved", "PDU Notification Reject Response", // 30
	"Supported Extension Headers Notification", "Send Routeing Information for GPRS Request", "Send Routeing Information for GPRS", "Failure Report Request", "Failure Report Response", // 35
	"Note MS GPRS Present Request", "Note MS GPRS Present Response", "Reserved", "Reserved", "Reserved", // 40
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 45
	"Reserved", "Reserved", "Identification Request", "Identification Response", "SGSN Context Request", // 50
	"SGSN Context Response", "SGSN Context Acknowledge", "Forward Relocation Request", "Forward Relocation Response", "Forward Relocation Complete", // 55
	"Relocation Cancel Request", "Relocation Cancel Response", "Forward SRNS Context", "Forward Relocation Complete Acknowledge", "Forward SRNS Context Acknowledge", // 60
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 65
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 70
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 75
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 80
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 85
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 90
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 95
	"MBMS Notification Request", "MBMS Notification Response", "MBMS Notification Reject Request", "MBMS Notification Reject Response", "Create MBMS Context Request", // 100
	"Create MBMS Context Response", "Update MBMS Context Request", "Update MBMS Context Response", "Delete MBMS Context Request", "Delete MBMS Context Response", // 105
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 110
	"Reserved", "MBMS Registration Request", "MBMS Registration Response", "MBMS De-Registration Request", "MBMS De-Registration Response", // 115
	"MBMS Session Start Request", "MBMS Session Start Response", "MBMS Session Stop Request", "MBMS Session Stop Response", "MBMS Session Update Request", // 120
	"MBMS Session Update Response", "Reserved", "Reserved", "Reserved", "Reserved", // 125
	"Reserved", "Reserved", "MS Info Change Notification Request", "MS Info Change Notification Response", "Reserved", // 130
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 135
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 140
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 145
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 150
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 155
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 160
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 165
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 170
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 175
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 180
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 185
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 190
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 195
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 200
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 205
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 210
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 215
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 220
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 225
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 230
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 235
	"Reserved", "Reserved", "Reserved", "Reserved", "Data Record Transfer Request", // 240
	"Data Record Transfer Response", "Reserved", "Reserved", "Reserved", "Reserved", // 245
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 250
	"Reserved", "Reserved", "Reserved", "End Marker", "G-PDU", // 255
}

// NameOfMessageForType returns a string identifier (from TS 129.060 section 7) for
// a GTPv1 IE based on the type integer value
func NameOfMessageForType(msgType MessageType) string {
	return messageNames[int(msgType)]
}

// ExtensionHeader represents a GTPv1 extension header.
type ExtensionHeader struct {
	ExtensionLength uint8
	Contents        []byte
}

// PDU represents a GTPv1 PDU.  Version field is omitted because it is always '1' and
// PT flag is also omitted, because it is always set to 1b.  TotalLength includes
// complete header length, and body length.
type PDU struct {
	Type                MessageType
	TotalLength         uint16
	TEID                uint32
	SequenceNumber      uint16
	ExtensionHeaders    []*ExtensionHeader
	InformationElements []*IE
}

// NewPDU constructs a new base GTPv2 PDU.  It uses a builder pattern to
// add non-mandatory elements, including a TEID and a priority.  A piggybacked
// PDU is added at the time of encoding and revealed on decoding.  If you change
// struct values after construction, Encode() may not operate as expected and may
// even panic, so the struct values should usually be treated as read-only.
// This version of the constructor will panic if the length of the IEs exceeds
// the maximum PDU length.  If you want to be able to catch this condition,
// construct the PDU struct manually.
func NewPDU(pduType MessageType, sequenceNumber uint16, ies []*IE) *PDU {
	pduLength := uint32(8)

	for _, ie := range ies {
		// compute of IE length is data length + 4 bytes for IE header
		pduLength += uint32(len(ie.Data) + 4)
	}

	if pduLength > 0xffff {
		panic("Combined IE lengths exceed maximum PDU length")
	}

	return &PDU{
		Type:                pduType,
		TEID:                0,
		SequenceNumber:      sequenceNumber,
		InformationElements: ies,
		TotalLength:         uint16(pduLength),
	}
}

// AddTEID sets the TEID field and the teid presence flag
func (pdu *PDU) AddTEID(teid uint32) *PDU {
	pdu.TEID = teid
	pdu.TotalLength += 4

	return pdu
}

// AddPriority sets the priority field and the priority presence flag
func (pdu *PDU) AddPriority(priority uint8) *PDU {
	return pdu
}

// Encode encodes the GTPv2 PDU as a byte stream in network byte order,
// suitable for trasmission.
func (pdu *PDU) Encode() []byte {
	// encoded := make([]byte, pdu.TotalLength)

	// encoded[0] = 0x40
	// encoded[1] = uint8(pdu.Type)
	// binary.BigEndian.PutUint16(encoded[2:4], pdu.TotalLength-4)

	// ieOffsetByteIndex := 0

	// if pdu.TEIDFieldIsPresent {
	// 	encoded[0] |= 0x08
	// 	binary.BigEndian.PutUint32(encoded[4:8], pdu.TEID)
	// 	binary.BigEndian.PutUint32(encoded[8:12], pdu.SequenceNumber<<8)

	// 	if pdu.PriorityFieldIsPresent {
	// 		encoded[0] |= 0x04
	// 		encoded[11] = pdu.Priority << 4
	// 	}
	// 	ieOffsetByteIndex = 12
	// } else {
	// 	binary.BigEndian.PutUint32(encoded[4:8], pdu.SequenceNumber<<8)
	// 	ieOffsetByteIndex = 8
	// }

	// for _, ie := range pdu.InformationElements {
	// 	encodedIE := ie.Encode()
	// 	offsetForEndOfIE := ieOffsetByteIndex + len(encodedIE)

	// 	copy(encoded[ieOffsetByteIndex:offsetForEndOfIE], encodedIE)

	// 	ieOffsetByteIndex = offsetForEndOfIE
	// }

	// return encoded
	return nil
}

// DecodePDU decodes a stream of bytes that contain either exactly one well-formed
// GTPv2 PDU, or two GTPv2 PDUs when the piggyback flag on the first is set to true.
// Returns an error if the stream cannot be decoded into one or two PDUs.
func DecodePDU(stream []byte) (pdu *PDU, piggybackedPdu *PDU, err error) {
	piggybackedPdu = nil

	if len(stream) < 8 {
		return nil, nil, fmt.Errorf("stream length (%d) too short for a GTPv2 PDU", len(stream))
	}

	if (stream[0] >> 5) != 2 {
		return nil, nil, fmt.Errorf("GTPv2 PDU version should be 2, but in stream, it is (%d)", (stream[0] >> 5))
	}

	hasPiggybackedPdu := (stream[0] & 0x10) == 0x10

	msgLengthFieldValue := binary.BigEndian.Uint16(stream[2:4])
	totalPduLength := msgLengthFieldValue + 4

	if len(stream) < int(totalPduLength) {
		return nil, nil, fmt.Errorf("GTPv2 PDU length field is (%d), so total length should be (%d), but stream length is (%d)", msgLengthFieldValue, totalPduLength, len(stream))
	}

	if !hasPiggybackedPdu {
		if len(stream) != int(totalPduLength) {
			return nil, nil, fmt.Errorf("GTPv2 PDU length field is (%d), so total length should be (%d), but stream length is (%d)", msgLengthFieldValue, totalPduLength, len(stream))
		}
	} else {
		piggybackedPduStream := stream[totalPduLength:]

		if (piggybackedPduStream[0] & 0x10) != 0 {
			return nil, nil, fmt.Errorf("GTPv2 PDU has piggybacked PDU but the piggyback flag for that piggybacked PDU is not 0")
		}

		piggybackedPdu, _, err = DecodePDU(piggybackedPduStream)

		if err != nil {
			return nil, nil, fmt.Errorf("on piggybacked PDU: %s", err)
		}

		if len(stream) != int(totalPduLength)+int(piggybackedPdu.TotalLength) {
			return nil, nil, fmt.Errorf("stream contains more than single PDU and piggybacked PDU")
		}
	}

	teid := uint32(0)
	sequenceNumber := uint32(0)
	var headerLength int

	if (stream[0] & 0x08) == 0x08 {
		teid = binary.BigEndian.Uint32(stream[4:8])
		sequenceNumber = binary.BigEndian.Uint32(stream[8:12]) >> 8
		headerLength = 12
	} else {
		sequenceNumber = binary.BigEndian.Uint32(stream[4:8]) >> 8
		headerLength = 8
	}

	pdu = &PDU{
		TEID:           teid,
		SequenceNumber: uint16(sequenceNumber),
		TotalLength:    totalPduLength,
		Type:           MessageType(stream[1]),
	}

	ieSet := make([]*IE, 0, 10)

	for i := headerLength; i < int(totalPduLength); {
		nextIEInStream, err := DecodeIE(stream[i:])

		if err != nil {
			return nil, nil, err
		}

		ieSet = append(ieSet, nextIEInStream)

	}

	pdu.InformationElements = ieSet

	return pdu, piggybackedPdu, nil
}
