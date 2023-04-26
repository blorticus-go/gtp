package gtpv1_test

import (
	"net"
	"testing"

	"github.com/blorticus-go/gtp/gtpv1"
	"github.com/blorticus-go/protodef"
	"github.com/go-test/deep"
)

type v1PDUComparable struct {
	testName     string
	encodedBytes []byte
	matchingPdu  *gtpv1.PDU
}

type v1PDUNamesComparable struct {
	expectedName string
	pduType      gtpv1.MessageType
}

type gtpuPDUComparable struct {
	testName        string
	tunnelledPacket protodef.IPProtocolPDU
	encodedBytes    []byte
	tunnelID        uint32
	sequenceNumber  *uint16
}

func TestPDUNames(t *testing.T) {
	// This test set is mostly to make sure the list doesn't accidentally
	// get shifted if values are changed
	testCases := []v1PDUNamesComparable{
		{"Reserved", 0},
		{"Echo Request", 1},
		{"Send Routeing Information for GPRS", 33},
		{"MS Info Change Notification Response", 129},
		{"Reserved", 172},
		{"G-PDU", 255},
	}

	for _, testCase := range testCases {
		if retrievedName := gtpv1.NameOfMessageForType(testCase.pduType); retrievedName != testCase.expectedName {
			t.Errorf("For PDU Message Type (%d), expected name = (%s), got = (%s)", testCase.pduType, testCase.expectedName, retrievedName)
		}
	}
}

func TestPDUEncodeValid(t *testing.T) {
	testCases := []v1PDUComparable{
		{
			testName: "Properly formatted Create PDP Context Request Encode()",
			matchingPdu: gtpv1.NewPDU(gtpv1.CreatePDPContextRequest, 0xaabbccdd).WithInformationElements([]*gtpv1.IE{
				gtpv1.NewIEWithRawData(gtpv1.IMSI, []byte{0, 1, 2, 3, 4, 5, 6, 7}),
				gtpv1.NewIEWithRawData(gtpv1.TunnelEndpointIdentifierDataI, []byte{0x0a, 0x0b, 0x0c, 0x0d}),
				gtpv1.NewIEWithRawData(gtpv1.NSAPI, []byte{0x00}),
				gtpv1.NewIEWithRawData(gtpv1.GSNAddress, []byte{192, 168, 1, 1}),
				gtpv1.NewIEWithRawData(gtpv1.GSNAddress, []byte{10, 10, 10, 10}),
				gtpv1.NewIEWithRawData(gtpv1.QualityofServiceProfile, []byte{0x01, 0x02, 0x03}),
			}),
			encodedBytes: []byte{
				0x30, 0x10, 0x00, 0x24, 0xaa, 0xbb, 0xcc, 0xdd,
				0x02, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
				0x10, 0x0a, 0x0b, 0x0c, 0x0d,
				0x14, 0x00,
				133, 0, 4, 192, 168, 1, 1,
				133, 0, 4, 10, 10, 10, 10,
				135, 0, 3, 0x01, 0x02, 0x03,
			},
		},
	}

	for testCaseIndex, testCase := range testCases {
		encoded := testCase.matchingPdu.Encode()

		if err := compareByteArrays(testCase.encodedBytes, encoded); err != nil {
			t.Errorf("on test number (%d): %s", testCaseIndex+1, err.Error())
		}
	}
}

func TestPDUDecodeValid(t *testing.T) {
	testCases := []v1PDUComparable{
		{
			testName: "Properly formatted Create PDP Context Request Encode()",
			matchingPdu: gtpv1.NewPDU(gtpv1.CreatePDPContextRequest, 0xaabbccdd).WithInformationElements([]*gtpv1.IE{
				gtpv1.NewIEWithRawData(gtpv1.IMSI, []byte{0, 1, 2, 3, 4, 5, 6, 7}),
				gtpv1.NewIEWithRawData(gtpv1.TunnelEndpointIdentifierDataI, []byte{0x0a, 0x0b, 0x0c, 0x0d}),
				gtpv1.NewIEWithRawData(gtpv1.NSAPI, []byte{0x00}),
				gtpv1.NewIEWithRawData(gtpv1.GSNAddress, []byte{192, 168, 1, 1}),
				gtpv1.NewIEWithRawData(gtpv1.GSNAddress, []byte{10, 10, 10, 10}),
				gtpv1.NewIEWithRawData(gtpv1.QualityofServiceProfile, []byte{0x01, 0x02, 0x03}),
			}),
			encodedBytes: []byte{
				0x30, 0x10, 0x00, 0x24, 0xaa, 0xbb, 0xcc, 0xdd,
				0x02, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
				0x10, 0x0a, 0x0b, 0x0c, 0x0d,
				0x14, 0x00,
				133, 0, 4, 192, 168, 1, 1,
				133, 0, 4, 10, 10, 10, 10,
				135, 0, 3, 0x01, 0x02, 0x03,
			},
		},
	}

	for _, testCase := range testCases {
		decoded, err := gtpv1.DecodePDU(testCase.encodedBytes)
		if err != nil {
			t.Errorf("failed to decode PDU stream: %s", err.Error())
		}

		if diff := deep.Equal(testCase.matchingPdu, decoded); diff != nil {
			t.Errorf("[%s]: %s", testCase.testName, diff)
		}
	}
}

func TestGPDUEncode(t *testing.T) {
	sequenceNumbers := []uint16{0x28db}

	testCases := []gtpuPDUComparable{
		{
			testName:       "Properly formatted GPDU 01 Encoded()",
			tunnelID:       1,
			sequenceNumber: &sequenceNumbers[0],
			tunnelledPacket: protodef.NewIPv4Packet(net.IPv4(202, 11, 40, 158), net.IPv4(192, 168, 40, 178)).WithPayloadFromPdu(
				&protodef.ICMPv4PDU{
					Header: protodef.ICMPv4Header{
						Type:               protodef.ICMPv4EchoRequest,
						Code:               0,
						Checksum:           0xbee7,
						TypeSpecificHeader: []byte{0x00, 0x00, 0x28, 0x7b, 0x04, 0x11, 0x20, 0x4b, 0xf4, 0x3d, 0x0d, 0x00},
					},
					Data: []byte{0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
						0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
						0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37},
				},
			),
			encodedBytes: []byte{0x32, 0xff, 0x00, 0x58, 0x00, 0x00, 0x00, 0x01, 0x28, 0xdb, 0x00, 0x00,
				0x45, 0x00, 0x00, 0x54, 0x00, 0x00, 0x00, 0x00, 0x40, 0x01, 0x9e, 0xa5, 0xca, 0x0b, 0x28, 0x9e,
				0xc0, 0xa8, 0x28, 0xb2, 0x08, 0x00, 0xbe, 0xe7, 0x00, 0x00, 0x28, 0x7b, 0x04, 0x11, 0x20, 0x4b,
				0xf4, 0x3d, 0x0d, 0x00, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20, 0x21, 0x22, 0x23,
				0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30, 0x31, 0x32, 0x33,
				0x34, 0x35, 0x36, 0x37,
			},
		},
	}

	for _, testCase := range testCases {
		tpdu, err := testCase.tunnelledPacket.MarshallToNetworkByteOrder()
		if err != nil {
			t.Fatalf("[%s]: failed to marshall packet: %s", testCase.testName, err.Error())
		}

		pdu := gtpv1.NewGPDU(testCase.tunnelID, tpdu)
		if testCase.sequenceNumber != nil {
			pdu.UseSequenceNumber(*testCase.sequenceNumber)
		}
		encoded := pdu.Encode()

		if diff := deep.Equal(testCase.encodedBytes, encoded); diff != nil {
			t.Errorf("[%s]: %s", testCase.testName, diff)
		}
	}

}
