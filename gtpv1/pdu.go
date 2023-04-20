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
	CreatePDPContextRequest               MessageType = 16
	CreatePDPContextResponse              MessageType = 17
	UpdatePDPContextRequest               MessageType = 18
	UpdatePDPContextResponse              MessageType = 19
	DeletePDPContextRequest               MessageType = 20
	DeletePDPContextResponse              MessageType = 21
	InitiatePDPContextActivationRequest   MessageType = 22
	InitiatePDPContextActivationResponse  MessageType = 23
	ErrorIndication                       MessageType = 26
	PDUNotificationRequest                MessageType = 27
	PDUNotificationResponse               MessageType = 28
	PDUNotificationRejectRequest          MessageType = 30
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
	"Echo Request", "Echo Response", "Version Not Supported", "Node Alive Request", "Node Alive Response", // 5
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

// ExtensionHeaderType is an enum of GTPv1 Extension Header types
type ExtensionHeaderType uint8

const (
	NoMoreHeaders                          ExtensionHeaderType = 0
	MBMSSupportIndication                  ExtensionHeaderType = 1
	MSInfoChangeReportingSupportIndication ExtensionHeaderType = 2
	PDCPPDUNumber                          ExtensionHeaderType = 0xc0
	SuspendRequest                         ExtensionHeaderType = 0xc1
	SuspendRespoonse                       ExtensionHeaderType = 0xc2
)

// ExtensionHeader represents a GTPv1 extension header.  Contents must be in network byte order and
// must not include the next header type.
type ExtensionHeader struct {
	Type     ExtensionHeaderType
	Contents []byte
}

func (h *ExtensionHeader) LengthInDoubleWords() uint8 {
	// all defined Extension Headers are naturally bounded on a double-word, so assuming
	// a valid number of bytes were encoded, integer division should be correct.
	return uint8((len(h.Contents) + 2) / 4)
}

// PDU represents a GTPv1 PDU.  Version field is omitted because it is always '1' and
// PT flag is also omitted, because it is always set to 1b.  Length excludes the
// mandatory header (the first 8 bytes).
type PDU struct {
	Type                  MessageType
	IncludeSequenceNumber bool
	IncludeNPDUNumber     bool
	Length                uint16
	TEID                  uint32
	SequenceNumber        uint16
	NPDUNumber            uint8
	ExtensionHeaders      []*ExtensionHeader
	InformationElements   []*IE
}

// NewPDU constructs a new base GTPv1 PDU.  It uses a builder pattern to
// add non-mandatory elements, including a Sequence Number and Extension headers.
// If you change struct values after construction, Encode() may not operate as expected and may
// even panic, so the struct values should usually be treated as read-only.
// This version of the constructor will panic if the length of the IEs exceeds
// the maximum PDU length.  If you want to be able to catch this condition,
// construct the PDU struct manually.
func NewPDU(pduType MessageType, teid uint32) *PDU {
	return &PDU{
		Type:   pduType,
		Length: 0,
		TEID:   teid,
	}
}

// UseSequenceNumber adds a sequence number to the header and sets the S flag
func (pdu *PDU) UseSequenceNumber(sequenceNumber uint16) *PDU {
	pdu.SequenceNumber = sequenceNumber

	if !pdu.IncludeSequenceNumber {
		pdu.IncludeSequenceNumber = true
		pdu.Length += 2
	}

	return pdu
}

// UseNPDUNumber sets the priority field and the priority presence flag
func (pdu *PDU) UseNPDUNumber(number uint8) *PDU {
	pdu.NPDUNumber = number

	if !pdu.IncludeNPDUNumber {
		pdu.IncludeNPDUNumber = true
		pdu.Length += 1
	}

	return pdu
}

// WithExtensionHeaders sets the PDU Extension headers to the set provided.  There
// is no copy made, so the provided headers should not be modified after they are provided
// here.  panic if the set causes the PDU to exceed its maximum allowable length.
func (pdu *PDU) WithExtensionHeaders(headers []*ExtensionHeader) *PDU {
	if len(pdu.ExtensionHeaders) > 0 {
		for _, previousHeader := range pdu.ExtensionHeaders {
			pdu.Length -= uint16((int(previousHeader.LengthInDoubleWords()) * 4))
		}
	}

	pdu.ExtensionHeaders = headers

	for _, header := range headers {
		if (len(header.Contents)+2)%4 != 0 {
			panic("extension header is not an even multiple of 4 in length")
		}

		if len(header.Contents) != 1 {
			panic(fmt.Sprintf("all defined extension headers have a length of 1, but inserted header has a length of (%d)", len(header.Contents)))
		}

		headerLengthInBytes := uint16(header.LengthInDoubleWords() * 4)
		if 65535-headerLengthInBytes > pdu.Length {
			panic("extension headers exceed allowed length for a GTPv1 PDU")
		}

		pdu.Length += headerLengthInBytes
	}

	return pdu
}

// WithInformationElements add the associated IEs to the PDU.  The IEs are not
// copied, so they should not be modified after adding them.
func (pdu *PDU) WithInformationElements(ies []*IE) *PDU {
	if len(pdu.InformationElements) > 0 {
		for _, ie := range pdu.InformationElements {
			pdu.Length -= ie.encodedLength()
		}
	}

	pdu.InformationElements = ies

	for _, ie := range ies {
		ieLength := ie.encodedLength()

		if pdu.Length > 65535-ieLength {
			panic("information elements are too large for a single PDU")
		}

		pdu.Length += ie.encodedLength()
	}

	return pdu
}

// Encode encodes the GTPv1 PDU as a byte stream in network byte order,
// suitable for trasmission.
func (pdu *PDU) Encode() []byte {
	encoded := make([]byte, pdu.Length+8)

	encoded[0] = 0x20

	encoded[1] = byte(pdu.Type)
	binary.BigEndian.PutUint16(encoded[2:4], pdu.Length)
	binary.BigEndian.PutUint32(encoded[4:8], pdu.TEID)

	indexOfNextByteToWrite := 8

	if pdu.IncludeSequenceNumber {
		encoded[0] |= 0x02
		binary.BigEndian.PutUint16(encoded[8:10], pdu.SequenceNumber)
		indexOfNextByteToWrite = 10
	}

	if pdu.IncludeNPDUNumber {
		encoded[0] |= 0x01
		encoded[indexOfNextByteToWrite] = pdu.NPDUNumber
		indexOfNextByteToWrite++
	}

	if len(pdu.ExtensionHeaders) > 0 {
		encoded[0] |= 0x04

		for _, header := range pdu.ExtensionHeaders {
			encoded[indexOfNextByteToWrite] = byte(header.Type)
			encoded[indexOfNextByteToWrite+1] = header.LengthInDoubleWords()
			i := indexOfNextByteToWrite + 2 + len(header.Contents)
			copy(encoded[indexOfNextByteToWrite+2:i], header.Contents)
			indexOfNextByteToWrite = i
		}

		encoded[indexOfNextByteToWrite] = byte(NoMoreHeaders)
		indexOfNextByteToWrite++
	}

	for _, ie := range pdu.InformationElements {
		encodedIE := ie.Encode()
		i := indexOfNextByteToWrite + len(encodedIE)
		copy(encoded[indexOfNextByteToWrite:i], encodedIE)
		indexOfNextByteToWrite = i
	}

	return encoded
}

// DecodePDU decodes a stream of bytes that contain either exactly one well-formed
// GTPv2 PDU, or two GTPv2 PDUs when the piggyback flag on the first is set to true.
// Returns an error if the stream cannot be decoded into one or two PDUs.
func DecodePDU(stream []byte) (pdu *PDU, err error) {
	return nil, nil
}
