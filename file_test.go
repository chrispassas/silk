package silk

import (
	"bytes"
	"encoding/binary"
	"net"
	"os"
	"testing"
)

//getTestDataFTRWIPV6ROUTINGV6V2 88 byte records
func getTestDataFTRWIPV6ROUTINGV6V2() testDetails {
	return testDetails{
		files: []string{
			"testdata/FT_RWIPV6ROUTING-v2-c0-L.dat",
			"testdata/FT_RWIPV6ROUTING-v2-c0-B.dat",
			"testdata/FT_RWIPV6ROUTING-v2-c1-L.dat",
			"testdata/FT_RWIPV6ROUTING-v2-c1-B.dat",
			"testdata/FT_RWIPV6ROUTING-v2-c2-L.dat",
			"testdata/FT_RWIPV6ROUTING-v2-c2-B.dat",
			"testdata/FT_RWIPV6ROUTING-v2-c3-L.dat",
			"testdata/FT_RWIPV6ROUTING-v2-c3-B.dat",
		},
		flows: []Flow{
			{
				SrcIP:        net.ParseIP("192.168.40.20"),
				DstIP:        net.ParseIP("10.0.40.54"),
				SrcPort:      88,
				DstPort:      60339,
				Proto:        6,
				Packets:      4,
				Bytes:        373,
				Flags:        30,
				StartTimeMS:  1434553200013,
				Duration:     6,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    net.ParseIP("0.0.0.0"),
				ClassType:    1,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
			{
				SrcIP:        net.ParseIP("192.168.20.58"),
				DstIP:        net.ParseIP("128.63.2.53"),
				SrcPort:      29070,
				DstPort:      53,
				Proto:        17,
				Packets:      1,
				Bytes:        74,
				Flags:        0,
				StartTimeMS:  1434553200025,
				Duration:     0,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    net.ParseIP("0.0.0.0"),
				ClassType:    1,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
		},
	}
}

//getTestDataFTRWIPV6V2 56 byte records
func getTestDataFTRWIPV6V2() testDetails {
	return testDetails{
		files: []string{
			"testdata/FT_RWIPV6-v2-c0-L.dat",
			"testdata/FT_RWIPV6-v2-c0-B.dat",
			"testdata/FT_RWIPV6-v2-c1-L.dat",
			"testdata/FT_RWIPV6-v2-c1-B.dat",
			"testdata/FT_RWIPV6-v2-c2-L.dat",
			"testdata/FT_RWIPV6-v2-c2-B.dat",
			"testdata/FT_RWIPV6-v2-c3-L.dat",
			"testdata/FT_RWIPV6-v2-c3-B.dat",
		},
		flows: []Flow{
			{
				SrcIP:        net.ParseIP("192.168.40.20"),
				DstIP:        net.ParseIP("10.0.40.54"),
				SrcPort:      88,
				DstPort:      60339,
				Proto:        6,
				Packets:      4,
				Bytes:        373,
				Flags:        30,
				StartTimeMS:  1434553200013,
				Duration:     6,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    nil,
				ClassType:    0,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
			{
				SrcIP:        net.ParseIP("192.168.20.58"),
				DstIP:        net.ParseIP("128.63.2.53"),
				SrcPort:      29070,
				DstPort:      53,
				Proto:        17,
				Packets:      1,
				Bytes:        74,
				Flags:        0,
				StartTimeMS:  1434553200025,
				Duration:     0,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    nil,
				ClassType:    0,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
		},
	}
}

//getTestDataFTRWIPV6V2 68 byte records
func getTestDataFTRWIPV6V1() testDetails {
	return testDetails{
		files: []string{
			"testdata/FT_RWIPV6-v1-c0-L.dat",
			"testdata/FT_RWIPV6-v1-c0-B.dat",
			"testdata/FT_RWIPV6-v1-c1-L.dat",
			"testdata/FT_RWIPV6-v1-c1-B.dat",
			"testdata/FT_RWIPV6-v1-c2-L.dat",
			"testdata/FT_RWIPV6-v1-c2-B.dat",
			"testdata/FT_RWIPV6-v1-c3-L.dat",
			"testdata/FT_RWIPV6-v1-c3-B.dat",
		},
		flows: []Flow{
			{
				SrcIP:        net.ParseIP("192.168.40.20"),
				DstIP:        net.ParseIP("10.0.40.54"),
				SrcPort:      88,
				DstPort:      60339,
				Proto:        6,
				Packets:      4,
				Bytes:        373,
				Flags:        30,
				StartTimeMS:  1434553200013,
				Duration:     6,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    nil,
				ClassType:    1,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
			{
				SrcIP:        net.ParseIP("192.168.20.58"),
				DstIP:        net.ParseIP("128.63.2.53"),
				SrcPort:      29070,
				DstPort:      53,
				Proto:        17,
				Packets:      1,
				Bytes:        74,
				Flags:        0,
				StartTimeMS:  1434553200025,
				Duration:     0,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    nil,
				ClassType:    1,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
		},
	}
}

//getTestDataFTRWGENERICV5 52 byte records
func getTestDataFTRWGENERICV5() testDetails {
	return testDetails{
		files: []string{
			"testdata/FT_RWGENERIC-v5-c0-B.dat",
			"testdata/FT_RWGENERIC-v5-c0-L.dat",
			"testdata/FT_RWGENERIC-v5-c1-B.dat",
			"testdata/FT_RWGENERIC-v5-c1-L.dat",
		},

		flows: []Flow{
			{
				SrcIP:        net.ParseIP("192.168.40.20"),
				DstIP:        net.ParseIP("10.0.40.54"),
				SrcPort:      88,
				DstPort:      60339,
				Proto:        6,
				Packets:      4,
				Bytes:        373,
				Flags:        30,
				StartTimeMS:  1434553200013,
				Duration:     6,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    net.ParseIP("0.0.0.0"),
				ClassType:    1,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
			{
				SrcIP:        net.ParseIP("192.168.20.58"),
				DstIP:        net.ParseIP("128.63.2.53"),
				SrcPort:      29070,
				DstPort:      53,
				Proto:        17,
				Packets:      1,
				Bytes:        74,
				Flags:        0,
				StartTimeMS:  1434553200025,
				Duration:     0,
				Sensor:       3,
				SNMPIn:       0,
				SNMPOut:      0,
				NextHopIP:    net.ParseIP("0.0.0.0"),
				ClassType:    1,
				InitalFlags:  0,
				SessionFlags: 0,
				Attributes:   0,
				Application:  0,
			},
		},
	}
}

type testDetails struct {
	files []string
	flows []Flow
}

//TestFiles Test all supported file formats. Reads each file and
//verify first 2 rows all fields match expected test values.
func TestFiles(t *testing.T) {
	t.Logf("TestFiles()")
	var testData = []testDetails{
		getTestDataFTRWIPV6V2(),          //56 byte
		getTestDataFTRWIPV6V1(),          //68 byte
		getTestDataFTRWIPV6ROUTINGV6V2(), //88 byte
		getTestDataFTRWGENERICV5(),       //52 byte
	}
	for _, testFlowData := range testData {
		for _, filePath := range testFlowData.files {
			var sf File
			var err error
			if sf, err = OpenFile(filePath); err != nil {
				t.Errorf("OpenFile file:%s error:%s", filePath, err)
			}

			if len(sf.Flows) != 245340 {
				t.Errorf("File:%s Rows found:%d, rows expected:%d", filePath, len(sf.Flows), 245340)
			}
			if len(sf.Flows) < len(testFlowData.flows) {
				t.Errorf("Test file:%s has fewer rows:%d then test:%d", filePath, len(sf.Flows), len(testFlowData.flows))
			}

			for x := 0; x < len(testFlowData.flows); x++ {

				if len(sf.Flows) <= x {
					break
				}

				if sf.Flows[x].SrcPort != testFlowData.flows[x].SrcPort {
					t.Errorf("Test SrcPort:%d not equal to data row:%d file:%s", testFlowData.flows[x].SrcPort, sf.Flows[x].SrcPort, filePath)
				}
				if sf.Flows[x].DstPort != testFlowData.flows[x].DstPort {
					t.Errorf("Test DstPort:%d not equal to data row:%d file:%s", testFlowData.flows[x].DstPort, sf.Flows[x].DstPort, filePath)
				}
				if sf.Flows[x].SrcIP.String() != testFlowData.flows[x].SrcIP.String() {
					t.Errorf("Test SrcIP:%s not equal to data row:%s file:%s", testFlowData.flows[x].SrcIP.String(), sf.Flows[x].SrcIP.String(), filePath)
				}
				if sf.Flows[x].DstIP.String() != testFlowData.flows[x].DstIP.String() {
					t.Errorf("Test DstIP:%s not equal to data row:%s file:%s", testFlowData.flows[x].DstIP.String(), sf.Flows[x].DstIP.String(), filePath)
				}
				if sf.Flows[x].Proto != testFlowData.flows[x].Proto {
					t.Errorf("Test Proto:%d not equal to data row:%d file:%s", testFlowData.flows[x].Proto, sf.Flows[x].Proto, filePath)
				}
				if sf.Flows[x].Packets != testFlowData.flows[x].Packets {
					t.Errorf("Test Packets:%d not equal to data row:%d file:%s", testFlowData.flows[x].Packets, sf.Flows[x].Packets, filePath)
				}
				if sf.Flows[x].Bytes != testFlowData.flows[x].Bytes {
					t.Errorf("Test Bytes:%d not equal to data row:%d file:%s", testFlowData.flows[x].Bytes, sf.Flows[x].Bytes, filePath)
				}
				if sf.Flows[x].Flags != testFlowData.flows[x].Flags {
					t.Errorf("Test Flags:%d not equal to data row:%d file:%s", testFlowData.flows[x].Flags, sf.Flows[x].Flags, filePath)
				}
				if sf.Flows[x].StartTimeMS != testFlowData.flows[x].StartTimeMS {
					t.Errorf("Test StartTimeMS:%d not equal to data row:%d file:%s", testFlowData.flows[x].StartTimeMS, sf.Flows[x].StartTimeMS, filePath)
				}
				if sf.Flows[x].Sensor != testFlowData.flows[x].Sensor {
					t.Errorf("Test Sensor:%d not equal to data row:%d file:%s", testFlowData.flows[x].Sensor, sf.Flows[x].Sensor, filePath)
				}
				if sf.Flows[x].Duration != testFlowData.flows[x].Duration {
					t.Errorf("Test Duration:%d not equal to data row:%d file:%s", testFlowData.flows[x].Duration, sf.Flows[x].Duration, filePath)
				}
				if sf.Flows[x].SNMPIn != testFlowData.flows[x].SNMPIn {
					t.Errorf("Test SNMPIn:%d not equal to data row:%d file:%s", testFlowData.flows[x].SNMPIn, sf.Flows[x].SNMPIn, filePath)
				}
				if sf.Flows[x].SNMPOut != testFlowData.flows[x].SNMPOut {
					t.Errorf("Test SNMPOut:%d not equal to data row:%d file:%s", testFlowData.flows[x].SNMPOut, sf.Flows[x].SNMPOut, filePath)
				}
				if sf.Flows[x].NextHopIP.Equal(testFlowData.flows[x].NextHopIP) == false {
					t.Errorf("Test NextHopIP:%s not equal to data row:%s file:%s", testFlowData.flows[x].NextHopIP.String(), sf.Flows[x].NextHopIP.String(), filePath)
				}
				if sf.Flows[x].ClassType != testFlowData.flows[x].ClassType {
					t.Errorf("Test ClassType:%d not equal to data row:%d file:%s", testFlowData.flows[x].ClassType, sf.Flows[x].ClassType, filePath)
				}
				if sf.Flows[x].InitalFlags != testFlowData.flows[x].InitalFlags {
					t.Errorf("Test InitalFlags:%d not equal to data row:%d file:%s", testFlowData.flows[x].InitalFlags, sf.Flows[x].InitalFlags, filePath)
				}
				if sf.Flows[x].SessionFlags != testFlowData.flows[x].SessionFlags {
					t.Errorf("Test SessionFlags:%d not equal to data row:%d file:%s", testFlowData.flows[x].SessionFlags, sf.Flows[x].SessionFlags, filePath)
				}
				if sf.Flows[x].Attributes != testFlowData.flows[x].Attributes {
					t.Errorf("Test Attributes:%d not equal to data row:%d file:%s", testFlowData.flows[x].Attributes, sf.Flows[x].Attributes, filePath)
				}
				if sf.Flows[x].Application != testFlowData.flows[x].Application {
					t.Errorf("Test Application:%d not equal to data row:%d file:%s", testFlowData.flows[x].Application, sf.Flows[x].Application, filePath)
				}
			}
		}
	}

}

//TestParseRecordOffsetsAreRecordRelative is a regression test for a bug where the
//ClassType, InitalFlags, SessionFlags and Attributes fields were read using offsets
//relative to the start of the read buffer instead of offsets relative to the current
//record's start. Every record after the first in a block silently inherited record 0's
//values for those fields. This builds a minimal, synthetic (uncompressed, 52-byte
//record) silk file with two records whose values differ, so the bug can't hide behind
//coincidentally-equal fixture data the way it did in TestFiles/TestParse.
func TestParseRecordOffsetsAreRecordRelative(t *testing.T) {
	const recordSize = 52

	buildHeader := func() []byte {
		h := make([]byte, 0, recordSize)
		h = append(h, 0xde, 0xad, 0xbe, 0xef) // MagicNumber (unchecked by parseHeader)
		h = append(h, 0)                      // FileFlags: 0 = little endian
		h = append(h, 0)                      // RecordFormat
		h = append(h, 5)                      // FileVersion
		h = append(h, 0)                      // Compression: 0 = none
		h = append(h, 0, 0, 0, 1)             // SilkVersion
		var recSize [2]byte
		binary.BigEndian.PutUint16(recSize[:], recordSize)
		h = append(h, recSize[:]...) // RecordSize
		h = append(h, 0, 5)          // RecordVersion

		// Minimal terminating variable-length header: id=0, length=8, no content.
		h = append(h, 0, 0, 0, 0) // ID
		h = append(h, 0, 0, 0, 8) // Length

		// parseHeader pads the header out to a RecordSize boundary.
		for len(h) < recordSize {
			h = append(h, 0)
		}
		return h
	}

	// Field byte offsets below match getOffsets()'s recordSize==52 case in file.go.
	buildRecord := func(srcPort uint16, classType, initalFlags, sessionFlags, attributes uint8) []byte {
		r := make([]byte, recordSize)
		binary.LittleEndian.PutUint16(r[12:14], srcPort) // startSrcPort:endSrcPort
		r[17] = classType                                // startClassType
		r[21] = initalFlags                              // startInitalFlags
		r[22] = sessionFlags                             // startSessionFlags
		r[23] = attributes                               // startAttributes
		return r
	}

	var data []byte
	data = append(data, buildHeader()...)
	data = append(data, buildRecord(100, 1, 2, 3, 4)...)
	data = append(data, buildRecord(200, 11, 12, 13, 14)...)

	receiver := NewSliceFlowReceiver(2)
	if err := Parse(bytes.NewReader(data), receiver); err != nil {
		t.Fatalf("Parse() error:%s", err)
	}

	if len(receiver.Flows) != 2 {
		t.Fatalf("expected 2 flows, got %d", len(receiver.Flows))
	}

	expected := []Flow{
		{SrcPort: 100, ClassType: 1, InitalFlags: 2, SessionFlags: 3, Attributes: 4},
		{SrcPort: 200, ClassType: 11, InitalFlags: 12, SessionFlags: 13, Attributes: 14},
	}

	for x, exp := range expected {
		got := receiver.Flows[x]
		if got.SrcPort != exp.SrcPort {
			t.Errorf("record %d: SrcPort = %d, want %d", x, got.SrcPort, exp.SrcPort)
		}
		if got.ClassType != exp.ClassType {
			t.Errorf("record %d: ClassType = %d, want %d", x, got.ClassType, exp.ClassType)
		}
		if got.InitalFlags != exp.InitalFlags {
			t.Errorf("record %d: InitalFlags = %d, want %d", x, got.InitalFlags, exp.InitalFlags)
		}
		if got.SessionFlags != exp.SessionFlags {
			t.Errorf("record %d: SessionFlags = %d, want %d", x, got.SessionFlags, exp.SessionFlags)
		}
		if got.Attributes != exp.Attributes {
			t.Errorf("record %d: Attributes = %d, want %d", x, got.Attributes, exp.Attributes)
		}
	}
}

//TestParse Test all supported file formats using an alternative flow receiver
func TestParse(t *testing.T) {
	t.Logf("TestFiles()")
	var testData = []testDetails{
		getTestDataFTRWIPV6V2(),          //56 byte
		getTestDataFTRWIPV6V1(),          //68 byte
		getTestDataFTRWIPV6ROUTINGV6V2(), //88 byte
		getTestDataFTRWGENERICV5(),       //52 byte
	}
	for _, testFlowData := range testData {
		for _, filePath := range testFlowData.files {
			var err error
			receiver := NewChannelFlowReceiver(0)
			flows := make([]Flow, 0, 245340)
			reader, err := os.Open(filePath)
			if err != nil {
				t.Error(err)
				continue
			}

			go func() {
				if err = Parse(reader, receiver); err != nil {
					t.Errorf("OpenFile file:%s error:%s", filePath, err)
				}
			}()

			for flow := range receiver.Read() {
				flows = append(flows, flow)
			}

			if err = reader.Close(); err != nil {
				t.Error("unable to close reader")
			}

			if len(flows) != 245340 {
				t.Errorf("File:%s Rows found:%d, rows expected:%d", filePath, len(flows), 245340)
			}
			if len(flows) < len(testFlowData.flows) {
				t.Errorf("Test file:%s has fewer rows:%d then test:%d", filePath, len(flows), len(testFlowData.flows))
			}

			for x := 0; x < len(testFlowData.flows); x++ {

				if len(flows) <= x {
					break
				}

				if flows[x].SrcPort != testFlowData.flows[x].SrcPort {
					t.Errorf("Test SrcPort:%d not equal to data row:%d file:%s", testFlowData.flows[x].SrcPort, flows[x].SrcPort, filePath)
				}
				if flows[x].DstPort != testFlowData.flows[x].DstPort {
					t.Errorf("Test DstPort:%d not equal to data row:%d file:%s", testFlowData.flows[x].DstPort, flows[x].DstPort, filePath)
				}
				if flows[x].SrcIP.String() != testFlowData.flows[x].SrcIP.String() {
					t.Errorf("Test SrcIP:%s not equal to data row:%s file:%s", testFlowData.flows[x].SrcIP.String(), flows[x].SrcIP.String(), filePath)
				}
				if flows[x].DstIP.String() != testFlowData.flows[x].DstIP.String() {
					t.Errorf("Test DstIP:%s not equal to data row:%s file:%s", testFlowData.flows[x].DstIP.String(), flows[x].DstIP.String(), filePath)
				}
				if flows[x].Proto != testFlowData.flows[x].Proto {
					t.Errorf("Test Proto:%d not equal to data row:%d file:%s", testFlowData.flows[x].Proto, flows[x].Proto, filePath)
				}
				if flows[x].Packets != testFlowData.flows[x].Packets {
					t.Errorf("Test Packets:%d not equal to data row:%d file:%s", testFlowData.flows[x].Packets, flows[x].Packets, filePath)
				}
				if flows[x].Bytes != testFlowData.flows[x].Bytes {
					t.Errorf("Test Bytes:%d not equal to data row:%d file:%s", testFlowData.flows[x].Bytes, flows[x].Bytes, filePath)
				}
				if flows[x].Flags != testFlowData.flows[x].Flags {
					t.Errorf("Test Flags:%d not equal to data row:%d file:%s", testFlowData.flows[x].Flags, flows[x].Flags, filePath)
				}
				if flows[x].StartTimeMS != testFlowData.flows[x].StartTimeMS {
					t.Errorf("Test StartTimeMS:%d not equal to data row:%d file:%s", testFlowData.flows[x].StartTimeMS, flows[x].StartTimeMS, filePath)
				}
				if flows[x].Sensor != testFlowData.flows[x].Sensor {
					t.Errorf("Test Sensor:%d not equal to data row:%d file:%s", testFlowData.flows[x].Sensor, flows[x].Sensor, filePath)
				}
				if flows[x].Duration != testFlowData.flows[x].Duration {
					t.Errorf("Test Duration:%d not equal to data row:%d file:%s", testFlowData.flows[x].Duration, flows[x].Duration, filePath)
				}
				if flows[x].SNMPIn != testFlowData.flows[x].SNMPIn {
					t.Errorf("Test SNMPIn:%d not equal to data row:%d file:%s", testFlowData.flows[x].SNMPIn, flows[x].SNMPIn, filePath)
				}
				if flows[x].SNMPOut != testFlowData.flows[x].SNMPOut {
					t.Errorf("Test SNMPOut:%d not equal to data row:%d file:%s", testFlowData.flows[x].SNMPOut, flows[x].SNMPOut, filePath)
				}
				if flows[x].NextHopIP.Equal(testFlowData.flows[x].NextHopIP) == false {
					t.Errorf("Test NextHopIP:%s not equal to data row:%s file:%s", testFlowData.flows[x].NextHopIP.String(), flows[x].NextHopIP.String(), filePath)
				}
				if flows[x].ClassType != testFlowData.flows[x].ClassType {
					t.Errorf("Test ClassType:%d not equal to data row:%d file:%s", testFlowData.flows[x].ClassType, flows[x].ClassType, filePath)
				}
				if flows[x].InitalFlags != testFlowData.flows[x].InitalFlags {
					t.Errorf("Test InitalFlags:%d not equal to data row:%d file:%s", testFlowData.flows[x].InitalFlags, flows[x].InitalFlags, filePath)
				}
				if flows[x].SessionFlags != testFlowData.flows[x].SessionFlags {
					t.Errorf("Test SessionFlags:%d not equal to data row:%d file:%s", testFlowData.flows[x].SessionFlags, flows[x].SessionFlags, filePath)
				}
				if flows[x].Attributes != testFlowData.flows[x].Attributes {
					t.Errorf("Test Attributes:%d not equal to data row:%d file:%s", testFlowData.flows[x].Attributes, flows[x].Attributes, filePath)
				}
				if flows[x].Application != testFlowData.flows[x].Application {
					t.Errorf("Test Application:%d not equal to data row:%d file:%s", testFlowData.flows[x].Application, flows[x].Application, filePath)
				}
			}
		}
	}

}

//getBenchFileList returns list of all supported test files
func getBenchFileList() []string {
	return []string{
		"testdata/FT_RWIPV6-v1-c0-L.dat",
		"testdata/FT_RWIPV6-v1-c0-B.dat",
		"testdata/FT_RWIPV6-v1-c1-L.dat",
		"testdata/FT_RWIPV6-v1-c1-B.dat",
		"testdata/FT_RWIPV6-v1-c2-L.dat",
		"testdata/FT_RWIPV6-v1-c2-B.dat",
		"testdata/FT_RWIPV6-v1-c3-L.dat",
		"testdata/FT_RWIPV6-v1-c3-B.dat",
		"testdata/FT_RWIPV6-v2-c0-L.dat",
		"testdata/FT_RWIPV6-v2-c0-B.dat",
		"testdata/FT_RWIPV6-v2-c1-L.dat",
		"testdata/FT_RWIPV6-v2-c1-B.dat",
		"testdata/FT_RWIPV6-v2-c2-L.dat",
		"testdata/FT_RWIPV6-v2-c2-B.dat",
		"testdata/FT_RWIPV6-v2-c3-L.dat",
		"testdata/FT_RWIPV6-v2-c3-B.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c0-L.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c0-B.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c1-L.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c1-B.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c2-L.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c2-B.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c3-L.dat",
		"testdata/FT_RWIPV6ROUTING-v2-c3-B.dat",
		"testdata/FT_RWGENERIC-v5-c0-B.dat",
		"testdata/FT_RWGENERIC-v5-c0-L.dat",
		"testdata/FT_RWGENERIC-v5-c1-B.dat",
		"testdata/FT_RWGENERIC-v5-c1-L.dat",
	}
}

//TestOpenNonExistingFile test opening non-existing file
func TestOpenNonExistingFile(t *testing.T) {
	var fakeFile = "fake_file_path.dat"
	var err error
	if _, err = OpenFile(fakeFile); err == nil {
		t.Errorf("OpenFile file:%s does not exist and shoulf fail", fakeFile)
	}
}

//BenchmarkReadFile read all test files to allow benchmarking how fast files can be read.
func BenchmarkReadFile(b *testing.B) {
	var err error
	var sf File
	var filePath string
	for n := 0; n < b.N; n++ {
		for _, filePath = range getBenchFileList() {
			if sf, err = OpenFile(filePath); err != nil {
				b.Errorf("OpenFile error:%s", err)
			}
			if len(sf.Flows) != 245340 {
				b.Errorf("File:%s Rows found:%d, rows expected:%d", filePath, len(sf.Flows), 245340)
			}
		}
	}
}
