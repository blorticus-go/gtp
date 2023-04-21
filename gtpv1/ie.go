package gtpv1

import (
	"encoding/binary"
	"fmt"
)

// IEType represents the various IE types for GTPv1
type IEType uint8

// These represent possible GTPv1 IE types.  In some cases, includes the
// full name and its abbreviation (e.g., for IMSI)
const (
	Cause                                 = 1
	InternationalMobileSubscriberIdentity = 2
	IMSI                                  = 2
	RouteingAreaIdentity                  = 3
	RAI                                   = 3
	TemporaryLogicalLinkIdentity          = 4
	LI                                    = 4
	PacketTMSI                            = 5
	PTMSI                                 = 5
	ReorderingRequired                    = 8
	AuthenticationTriplet                 = 9
	MAPCause                              = 11
	PTMSISignature                        = 12
	MSValidated                           = 13
	Recovery                              = 14
	SelectionMode                         = 15
	TunnelEndpointIdentifierDataI         = 16
	TunnelEndpointIdentifierControl       = 17
	TunnelEndpointIdentifierDataII        = 18
	TeardownInd                           = 19
	NSAPI                                 = 20
	RANAPCause                            = 21
	RABContext                            = 22
	RadioPrioritySMS                      = 23
	RadioPriority                         = 24
	PacketFlowId                          = 25
	ChargingCharacteristics               = 26
	TraceReference                        = 27
	TraceType                             = 28
	MSNotReachableReason                  = 29
	ChargingID                            = 127
	EndUserAddress                        = 128
	MMContext                             = 129
	PDPContext                            = 130
	AccessPointName                       = 131
	ProtocolConfigurationOptions          = 132
	GSNAddress                            = 133
	MSInternationalPSTNISDNNumber         = 134
	MSISDN                                = 134
	QualityofServiceProfile               = 135
	AuthenticationQuintuplet              = 136
	TrafficFlowTemplate                   = 137
	TargetIdentification                  = 138
	UTRANTransparentContainer             = 139
	RABSetupInformation                   = 140
	ExtensionHeaderTypeList               = 141
	TriggerId                             = 142
	OMCIdentity                           = 143
	RANTransparentContainer               = 144
	PDPContextPrioritization              = 145
	AdditionalRABSetupInformation         = 146
	SGSNNumber                            = 147
	CommonFlags                           = 148
	APNRestriction                        = 149
	RadioPriorityLCS                      = 150
	RATType                               = 151
	UserLocationInformation               = 152
	MSTimeZone                            = 153
	IMEI                                  = 154
	SV                                    = 154
	CAMELChargingInformationContainer     = 155
	MBMSUEContext                         = 156
	TemporaryMobileGroupIdentity          = 157
	TMGI                                  = 157
	RIMRoutingAddress                     = 158
	MBMSProtocolConfigurationOptions      = 159
	MBMSServiceArea                       = 160
	SourceRNCPDCPcontextinfo              = 161
	AdditionalTraceInfo                   = 162
	HopCounter                            = 163
	SelectedPLMNID                        = 164
	MBMSSessionIdentifier                 = 165
	MBMS2G3GIndicator                     = 166
	EnhancedNSAPI                         = 167
	MBMSSessionDuration                   = 168
	AdditionalMBMSTraceInfo               = 169
	MBMSSessionRepetitionNumber           = 170
	MBMSTimeToDataTransfer                = 171
	BSSContainer                          = 173
	CellIdentification                    = 174
	PDUNumbers                            = 175
	BSSGPCause                            = 176
	RequiredMBMSbearercapabilities        = 177
	RIMRoutingAddressDiscriminator        = 178
	ListofsetupPFCs                       = 179
	PSHandoverXIDParameters               = 180
	MSInfoChangeReportingAction           = 181
	DirectTunnelFlags                     = 182
	CorrelationID                         = 183
	BearerControlMode                     = 184
	MBMSFlowIdentifier                    = 185
	MBMSIPMulticastDistribution           = 186
	MBMSDistributionAcknowledgement       = 187
	ReliableINTERRATHANDOVERINFO          = 188
	RFSPIndex                             = 189
	FullyQualifiedDomainName              = 190
	FQDN                                  = 190
	EvolvedAllocationRetentionPriorityI   = 191
	EvolvedAllocationRetentionPriorityII  = 192
	ExtendedCommonFlags                   = 193
	UserCSGInformation                    = 194
	UCI                                   = 194
	CSGInformationReportingAction         = 195
	CSGID                                 = 196
	CSGMembershipIndication               = 197
	CMI                                   = 197
	AggregateMaximumBitRate               = 198
	AMBR                                  = 198
	UENetworkCapability                   = 199
	UEAMBR                                = 200
	APNAMBRwithNSAPI                      = 201
	GGSNBackOffTime                       = 202
	SignallingPriorityIndication          = 203
	SignallingPriorityIndicationwithNSAPI = 204
	Higherbitratesthan16Mbpsflag          = 205
	AdditionalMMcontextforSRVCC           = 207
	AdditionalflagsforSRVCC               = 208
	STNSR                                 = 209
	CMSISDN                               = 210
	ExtendedRANAPCause                    = 211
	eNodeBID                              = 212
	SelectionModewithNSAPI                = 213
	ULITimestamp                          = 214
	LocalHomeNetworkID                    = 215
	LHNID                                 = 215
	CNOperatorSelectionEntity             = 216
	ChargingGatewayAddress                = 251
	PrivateExtension                      = 255
)

var ieNames = []string{
	"Reserved",                                                                                                                                                 // 0
	"Cause", "International Mobile Subscriber Identity (IMSI)", "Routeing Area Identity (RAI)", "Temporary Logical Link Identity (LI)", "Packet TMSI (P-TMSI)", // 5
	"Reserved", "Reserved", "Reordering Required", "Authentication Triplet", "Reserved", // 10
	"MAP Cause", "P-TMSI Signature", "MS Validated", "Recovery", "Selection Mode", // 15
	"Tunnel Endpoint Identifier Data I", "Tunnel Endpoint Identifier Control", "Tunnel Endpoint Identifier Data II", "Teardown Ind", "NSAPI", // 20
	"RANAP Cause", "RAB Context", "Radio Priority SMS", "Radio Priority", "Packet Flow Id", // 25
	"Charging Characteristics", "Trace Reference", "Trace Type", "MS Not Reachable Reason", "Reserved", // 30
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 35
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 40
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 45
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 50
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 55
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 60
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 65
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 70
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 75
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 80
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 85
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 90
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 95
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 100
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 105
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 110
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 115
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 120
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 125
	"Reserved", "Charging ID", "End User Address", "MM Context", "PDP Context", // 130
	"Access Point Name", "Protocol Configuration Options", "GSN Address", "MS International PSTN/ISDN Number (MSISDN)", "Quality of Service Profile", // 135
	"Authentication Quintuplet", "Traffic Flow Template", "Target Identification", "UTRAN Transparent Container", "RAB Setup Information", // 140
	"Extension Header Type List", "Trigger Id", "OMC Identity", "RAN Transparent Container", "PDP Context Prioritization", // 145
	"Additional RAB Setup Information", "SGSN Number", "Common Flags", "APN Restriction", "Radio Priority LCS", // 150
	"RAT Type", "User Location Information", "MS Time Zone", "IMEI(SV)", "CAMEL Charging Information Container", // 155
	"MBMS UE Context", "Temporary Mobile Group Identity (TMGI)", "RIM Routing Address", "MBMS Protocol Configuration Options", "MBMS Service Area", // 160
	"Source RNC PDCP context info", "Additional Trace Info", "Hop Counter", "Selected PLMN ID", "MBMS Session Identifier", // 165
	"MBMS 2G/3G Indicator", "Enhanced NSAPI", "MBMS Session Duration", "Additional MBMS Trace Info", "MBMS Session Repetition Number", // 170
	"MBMS Time To Data Transfer", "Reserved", "BSS Container", "Cell Identification", "PDU Numbers", // 175
	"BSSGP Cause", "Required MBMS bearer capabilities", "RIM Routing Address Discriminator", "List of set-up PFCs", "PS Handover XID Parameters", // 180
	"MS Info Change Reporting Action", "Direct Tunnel Flags", "Correlation-ID", "Bearer Control Mode", "MBMS Flow Identifier", // 185
	"MBMS IP Multicast Distribution", "MBMS Distribution Acknowledgement", "Reliable INTER RAT HANDOVER INFO", "RFSP Index", "Fully Qualified Domain Name (FQDN)", // 190
	"Evolved Allocation/Retention Priority I", "Evolved Allocation/Retention Priority II", "Extended Common Flags", "User CSG Information (UCI)", "CSG Information Reporting Action", // 195
	"CSG ID", "CSG Membership Indication (CMI)", "Aggregate Maximum Bit Rate (AMBR)", "UE Network Capability", "UE-AMBR", // 200
	"APN-AMBR with NSAPI", "GGSN Back-Off Time", "Signalling Priority Indication", "Signalling Priority Indication with NSAPI", "Higher bitrates than 16 Mbps flag", // 205
	"Reserved", "Additional MM context for SRVCC", "Additional flags for SRVCC", "STN-SR", "C-MSISDN", // 210
	"Extended RANAP Cause", "eNodeB ID", "Selection Mode with NSAPI", "ULI Timestamp", "Local Home Network ID (LHN-ID) with NSAPI", // 215
	"CN Operator Selection Entity", "Reserved", "Reserved", "Reserved", "Reserved", // 220
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 225
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 230
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 235
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 240
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 245
	"Reserved", "Reserved", "Reserved", "Reserved", "Reserved", // 250
	"Charging Gateway Address", "Reserved", "Reserved", "Reserved", "Private Extension", // 255
}

var ieSizes = map[uint8]uint16{
	1: 1, 2: 8, 3: 6, 4: 4, 5: 4, 8: 1, 9: 28,
	11: 1, 12: 2, 13: 1, 14: 1, 15: 1, 16: 4, 17: 4, 18: 5, 19: 1, 20: 1,
	21: 1, 22: 9, 23: 1, 24: 1, 25: 2, 26: 2, 27: 2, 28: 2, 29: 1,
	127: 4, 145: 0, 148: 1, 149: 1, 150: 1, 151: 1, 153: 1, 154: 8,
	157: 6, 162: 9, 163: 1, 164: 3, 165: 1, 166: 1, 167: 1, 168: 3,
	169: 8, 170: 1, 171: 1, 174: 17, 175: 9, 176: 1, 178: 1, 181: 1,
	183: 1, 184: 1, 187: 1, 188: 1, 189: 2, 191: 1, 192: 2, 194: 8,
	196: 4, 197: 1, 198: 8, 201: 9, 202: 1, 203: 1, 204: 2, 205: 1,
	208: 1, 211: 2, 213: 2, 214: 4, 216: 1,
}

// NameOfIEForType returns a string identifier (from TS 129.060 section 7.7) for
// a GTPv1 IE based on the type integer
func NameOfIEForType(ieType IEType) string {
	return ieNames[int(ieType)]
}

// IE is a GTPv1 Information Element.  Data is the BigEndian data bytes.
type IE struct {
	Type IEType
	Data []byte
}

func (ie *IE) encodedLength() uint16 {
	if ie.Type < 128 {
		if dataLength, fixedSizeIsInTheMap := ieSizes[uint8(ie.Type)]; fixedSizeIsInTheMap {
			return dataLength + 1
		} else {
			panic(fmt.Sprintf("information element of type (%d) has no defined length", ie.Type))
		}
	} else {
		return uint16(len(ie.Data)) + 3
	}
}

// DecodeIE consumes bytes from the start of a stream to produce a GTPv1 IE.
// streamBytesConsumed is the number of bytes from the start of stream that
// were consumed to produce IE.
func DecodeIE(stream []byte) (ie *IE, streamBytesConsumed int, err error) {
	if len(stream) < 4 {
		return nil, 0, fmt.Errorf("insufficient octets in stream for a complete GTPv1 IE")
	}

	ieType := uint8(stream[0])

	ie = &IE{
		Type: IEType(ieType),
	}

	if ieType < 128 {
		if dataLength, thisIsAValidIE := ieSizes[ieType]; !thisIsAValidIE {
			return nil, 0, fmt.Errorf("no defined length for IE of type (0x%02x)", ieType)
		} else {
			if len(stream) < int(dataLength)+1 {
				return nil, 0, fmt.Errorf("for IE of type (0x%02x) insufficient bytes in stream", ieType)
			}

			ie.Data = make([]byte, dataLength)
			copy(ie.Data, stream[1:dataLength+1])

			streamBytesConsumed = int(dataLength) + 1
		}
	} else {
		if len(stream) < 3 {
			return nil, 0, fmt.Errorf("insufficient bytes to retreive TLV length for IE of type (0x%02x)", ieType)
		}

		dataLength := binary.BigEndian.Uint16(stream[1:3])

		if len(stream) < int(dataLength)+3 {
			return nil, 0, fmt.Errorf("for IE of type (0x%02x) insufficient bytes in stream", ieType)
		}

		ie.Data = make([]byte, dataLength)
		copy(ie.Data, stream[3:dataLength+3])

		streamBytesConsumed = int(dataLength) + 3
	}

	return ie, streamBytesConsumed, nil
}

// NewIEWithRawData creates a new GTPv1 IE, providing it with the data as
// a raw byte array.  The data are validated against the length for IEs
// that have a fixed data length.  The data are not copied, so if you require
// that, you must manually copy() the data first.  The data must be in
// network byte order (i.e., big endian order).  This method panics on
// an error.  Use NewIEWithRawDataErrorable() to make the error catchable.
func NewIEWithRawData(ieType IEType, data []byte) *IE {
	ie, err := NewIEWithRawDataErrorable(ieType, data)

	if err != nil {
		panic(err)
	}

	return ie
}

// NewIEWithRawDataErrorable does the same as NewIEWithRawData() but
// returns an error if it occurs, rather than panicing.
func NewIEWithRawDataErrorable(ieType IEType, data []byte) (*IE, error) {
	if len(data) > 65535 {
		return nil, fmt.Errorf("data length %d exceeds maximum for an Information Element", len(data))
	}

	if fixedIELength, ieIsFixedLength := ieSizes[uint8(ieType)]; ieIsFixedLength {
		if len(data) != int(fixedIELength) {
			return nil, fmt.Errorf("an IE of type (0x%02x) has a fixed length of (%d) but recevied (%d) bytes", ieType, fixedIELength, len(data))
		}
	}

	return &IE{
		Type: ieType,
		Data: data,
	}, nil
}

// Encode encodes the Information Element as a series of
// bytes in network byte order.  There is no effort to validate
// that the IE Data field is correct for the type.  This permits
// the creation of structurally correct but semantically incorrect IEs.
// There is also no effort to validate that the data length is correct.
func (ie *IE) Encode() []byte {
	var encodedBytes []byte
	if uint8(ie.Type) < 128 {
		encodedBytes = make([]byte, len(ie.Data)+1)
		encodedBytes[0] = byte(ie.Type)
		copy(encodedBytes[1:], ie.Data)
	} else {
		encodedBytes = make([]byte, len(ie.Data)+3)
		encodedBytes[0] = byte(ie.Type)
		binary.BigEndian.PutUint16(encodedBytes[1:3], uint16(len(ie.Data)))
		copy(encodedBytes[3:], ie.Data)
	}

	return encodedBytes
}
