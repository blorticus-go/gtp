package gtpv1_test

import (
	"testing"

	"github.com/blorticus-go/gtp/gtpv1"
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
