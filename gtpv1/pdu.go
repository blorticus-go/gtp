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

var pduTypeIsDefined = []bool{
	false,
	true, true, true, true, true, // 5
	true, true, false, false, false, // 10
	false, false, false, false, false, // 15
	true, true, true, true, true, // 20
	true, true, true, false, false, // 25
	true, true, true, false, true, // 30
	true, true, true, true, true, // 35
	true, true, false, false, false, // 40
	false, false, false, false, false, // 45
	false, false, true, true, true, // 50
	true, true, true, true, true, // 55
	true, true, true, true, true, // 60
	false, false, false, false, false, // 65
	false, false, false, false, false, // 70
	false, false, false, false, false, // 75
	false, false, false, false, false, // 80
	false, false, false, false, false, // 85
	false, false, false, false, false, // 90
	false, false, false, false, false, // 95
	true, true, true, true, true, // 100
	true, true, true, true, true, // 105
	false, false, false, false, false, // 110
	false, true, true, true, true, // 115
	true, true, true, true, true, // 120
	true, false, false, false, false, // 125
	false, false, true, true, false, // 130
	false, false, false, false, false, // 135
	false, false, false, false, false, // 140
	false, false, false, false, false, // 145
	false, false, false, false, false, // 150
	false, false, false, false, false, // 155
	false, false, false, false, false, // 160
	false, false, false, false, false, // 165
	false, false, false, false, false, // 170
	false, false, false, false, false, // 175
	false, false, false, false, false, // 180
	false, false, false, false, false, // 185
	false, false, false, false, false, // 190
	false, false, false, false, false, // 195
	false, false, false, false, false, // 200
	false, false, false, false, false, // 205
	false, false, false, false, false, // 210
	false, false, false, false, false, // 215
	false, false, false, false, false, // 220
	false, false, false, false, false, // 225
	false, false, false, false, false, // 230
	false, false, false, false, false, // 235
	false, false, false, false, true, // 240
	true, false, false, false, false, // 245
	false, false, false, false, false, // 250
	false, false, false, true, true, // 255
}

var messageNames = []string{
	"Reserved",
	"Echo Request", "Echo Response", "Version Not Supported", "Node Alive Request", "Node Alive Response", // 5
	"Redirection Request", "Redirection Response", "Reserved", "Reserved", "Reserved", // 10
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 15
	"Create PDP Context Request", "Create PDP Context Response", "Update PDP Context Request", "Update PDP Context Response", "Delete PDP Context Request", // 20
	"Delete PDP Context Response", "Initiate PDP Context Activation Request", "Initiate PDP Context Activation Response", "Reserved", "Reserved", // 25
	"Error Indication", "PDU Notification Request", "PDU Notification Response", "Reserved", "PDU Notification Reject Response", // 30
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

var extensionHeaderTypeIsDefined = map[uint8]bool{
	0:    true,
	1:    true,
	2:    true,
	0xc0: true,
	0xc1: true,
	0xc2: true,
}

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
// mandatory header (the first 8 bytes).  If the Type is GPDU, InformationElements must
// be empty.  If Type is not GPDU, the TPDU field must be empty.  If TPDU is populated,
// it must be the tunnelled T-PDU in network byte order.
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
	TPDU                  []byte
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

// NewGPDU returns a GTPv1 G-PDU (type 255) with the T-PDU data, which
// must be in network byte order.  The tpdu is not copied, so if that
// is needed, the caller should copy() first.
func NewGPDU(teid uint32, tpdu []byte) *PDU {
	if len(tpdu) > 65535-8 {
		panic("G-PDU exceeds datagram maximum length")
	}

	return &PDU{
		Type:   GPDU,
		Length: uint16(len(tpdu)),
		TEID:   teid,
		TPDU:   tpdu,
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

// HeaderPadByteCount returns the number of padding bytes required to align the
// header on a 32-bit boundary
func (pdu *PDU) HeaderPadByteCount() uint8 {
	totalHeaderLength := 8
	if pdu.IncludeSequenceNumber {
		totalHeaderLength += 2
	}
	if pdu.IncludeNPDUNumber {
		totalHeaderLength++
	}
	if len(pdu.ExtensionHeaders) > 0 {
		totalHeaderLength++
	}

	if totalHeaderLength&0x03 != 0 {
		return uint8(4 - (totalHeaderLength & 0x03))
	}

	return 0
}

// Encode encodes the GTPv1 PDU as a byte stream in network byte order,
// suitable for trasmission.
func (pdu *PDU) Encode() []byte {
	headerByteCount := pdu.HeaderPadByteCount()
	pdu.Length += uint16(headerByteCount)

	encoded := make([]byte, pdu.Length+8)

	encoded[0] = 0x30

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

	for i := 0; i < int(headerByteCount); i++ {
		encoded[indexOfNextByteToWrite] = 0
		indexOfNextByteToWrite++
	}

	for _, ie := range pdu.InformationElements {
		encodedIE := ie.Encode()
		i := indexOfNextByteToWrite + len(encodedIE)
		copy(encoded[indexOfNextByteToWrite:i], encodedIE)
		indexOfNextByteToWrite = i
	}

	if pdu.TPDU != nil {
		copy(encoded[indexOfNextByteToWrite:], pdu.TPDU)
	}

	return encoded
}

// DecodePDU decodes the complete bytes from a UDP datagram that contains exactly one well-formed
// GTPv1 PDU.  Returns an error if the stream cannot be decoded.  All data from
// stream that are used are copied.
func DecodePDU(datagram []byte) (pdu *PDU, err error) {
	if len(datagram) == 0 {
		return nil, fmt.Errorf("incoming stream is zero length")
	}

	if datagram[0]&0x20 == 0 {
		return nil, fmt.Errorf("incorrect version identifier for GTPv1")
	}

	bytesRemainingToProcess := len(datagram)

	if bytesRemainingToProcess < 8 {
		return nil, fmt.Errorf("few bytes than minimum required for gtpv1 PDU")
	}

	pduType := datagram[1]
	if !pduTypeIsDefined[pduType] {
		return nil, fmt.Errorf("message type (%d) is not defined", pduType)
	}

	pdu = &PDU{
		Type:   MessageType(pduType),
		Length: binary.BigEndian.Uint16(datagram[2:4]),
		TEID:   binary.BigEndian.Uint32(datagram[4:8]),
	}

	bytesRemainingToProcess -= 8

	if pdu.Length != uint16(bytesRemainingToProcess) {
		return nil, fmt.Errorf("length field value (%d) does not match stream length (%d) less the fixed header length (8)", pdu.Length, len(datagram))
	}

	if datagram[0]&0x02 != 0 { // Sequence Number flag
		if bytesRemainingToProcess < 2 {
			return nil, fmt.Errorf("insufficient bytes in datagram to include a sequence number")
		}

		pdu.SequenceNumber = binary.BigEndian.Uint16(datagram[8:10])
		pdu.IncludeSequenceNumber = true

		bytesRemainingToProcess -= 2
	}

	if datagram[0]&0x01 != 0 { // NPDU Number flag
		if bytesRemainingToProcess < 1 {
			return nil, fmt.Errorf("insufficient bytes in datagram to include a N-PDU number")
		}

		pdu.NPDUNumber = datagram[10]
		pdu.IncludeNPDUNumber = true

		bytesRemainingToProcess -= 1
	}

	if datagram[0]&0x04 != 0 { // Extension Header flag
		if bytesRemainingToProcess < 1 {
			return nil, fmt.Errorf("expected extension header but ran out of bytes in datagram")
		}

		extensionHeaders := make([]*ExtensionHeader, 0, 1)

		for i := len(datagram) - bytesRemainingToProcess; datagram[i] != byte(NoMoreHeaders); {
			if !extensionHeaderTypeIsDefined[uint8(datagram[i])] {
				return nil, fmt.Errorf("extension header of type (0x%02x) is not defined", datagram[i])
			}

			if bytesRemainingToProcess < 2 {
				return nil, fmt.Errorf("expected extension header but ran out of bytes in datagram")
			}

			nextHeaderLengthInBytes := uint16(datagram[i+1]) * 4

			if bytesRemainingToProcess < int(nextHeaderLengthInBytes) {
				return nil, fmt.Errorf("expected extension header but ran out of bytes in datagram")
			}

			contents := make([]byte, int(nextHeaderLengthInBytes))
			copy(contents, datagram[i+2:i+(int(nextHeaderLengthInBytes)-2)])

			extensionHeaders = append(extensionHeaders, &ExtensionHeader{
				Type:     ExtensionHeaderType(datagram[i]),
				Contents: contents,
			})

			i += int(nextHeaderLengthInBytes)
			bytesRemainingToProcess -= int(nextHeaderLengthInBytes)
		}

		pdu.ExtensionHeaders = extensionHeaders
	}

	if pdu.Type != GPDU {
		ieSet := make([]*IE, 0, 10)

		for i := len(datagram) - bytesRemainingToProcess; bytesRemainingToProcess > 0; {
			extractedIE, bytesConsumed, err := DecodeIE(datagram[i:])
			if err != nil {
				return nil, err
			}

			ieSet = append(ieSet, extractedIE)

			i += bytesConsumed
			bytesRemainingToProcess -= bytesConsumed
		}

		pdu.InformationElements = ieSet
	} else {
		tpdu := make([]byte, bytesRemainingToProcess)
		copy(tpdu, datagram[len(datagram)-bytesRemainingToProcess:])
		pdu.TPDU = tpdu
	}

	return pdu, nil
}
