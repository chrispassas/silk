package silk

import (
	"net"
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
			Flow{
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
			Flow{
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

//getTestDataFTRWIPV6V2 68 byte records
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
			Flow{
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
			Flow{
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
			Flow{
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
			Flow{
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
