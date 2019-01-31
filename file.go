package silk

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net"
	"os"

	"github.com/golang/snappy"
	lzo "github.com/rasky/go-lzo"
)

// SilkTCPStateExpanded constant value defined in silk code
const SilkTCPStateExpanded uint8 = 0x01

const defaultReadSize = 4096

var calMsec = uint32(math.Pow(2, 22) - 1)

var isTCPAnd = uint32(math.Pow(2, 23))

//Flow represents a silk flow row of data
//Depending on type of silk record not all fields are used
//More details on the Flow stuct fields can be found here:
//	https://tools.netsa.cert.org/silk/faq.html#file-formats
type Flow struct {
	startTimeMS56 uint32
	StartTimeMS   uint64
	Duration      uint32
	SrcIP         net.IP
	DstIP         net.IP
	SrcPort       uint16
	DstPort       uint16
	Proto         uint8
	Flags         uint8
	Packets       uint32
	Bytes         uint32
	ClassType     uint8
	Sensor        uint16
	InitalFlags   uint8
	SessionFlags  uint8
	Attributes    uint8
	Application   uint16
	SNMPIn        uint16
	SNMPOut       uint16
	NextHopIP     net.IP
}

// ErrUnsupportedCompression unknown compression type. Currently supported
// 	0 = no compression
// 	1 = zlib
// 	2 = lzo
// 	3 = snappy
var ErrUnsupportedCompression = fmt.Errorf("Unsupported compression")

type offsets struct {
	startStartTime    int
	endStartTime      int
	startDuration     int
	endDuration       int
	startSrcIP        int
	endSrcIP          int
	startDstIP        int
	endDstIP          int
	startSrcPort      int
	endSrcPort        int
	startDstPort      int
	endDstPort        int
	startProto        int
	startTCPFlags     int
	startPackets      int
	endPackets        int
	startBytes        int
	endBytes          int
	startClassType    int
	startSensor       int
	endSensor         int
	startInitalFlags  int
	startSessionFlags int
	startAttributes   int
	startApplication  int
	endApplication    int
	startSNMPIn       int
	endSNMPIn         int
	startSNMPOut      int
	endSNMPOut        int
	startNextHopIP    int
	endNextHopIP      int
}

func getOffsets(recordSize uint16) (o offsets, err error) {

	if recordSize == 88 {
		o.startStartTime = 0
		o.endStartTime = 8
		o.startSrcIP = 40
		o.endSrcIP = 56
		o.startDstIP = 56
		o.endDstIP = 72
		o.startSrcPort = 12
		o.endSrcPort = 14
		o.startDstPort = 14
		o.endDstPort = 16
		o.startProto = 16
		o.startPackets = 32
		o.endPackets = 36
		o.startBytes = 36
		o.endBytes = 40
		o.startDuration = 8
		o.endDuration = 12
		o.startTCPFlags = 20
		o.startClassType = 17
		o.startSensor = 18
		o.endSensor = 20
		o.startInitalFlags = 21
		o.startSessionFlags = 22
		o.startAttributes = 23
		o.startApplication = 24
		o.endApplication = 26
		o.startSNMPIn = 28
		o.endSNMPIn = 30
		o.startSNMPOut = 30
		o.endSNMPOut = 32
		o.startNextHopIP = 72
		o.endNextHopIP = 88
	} else if recordSize == 56 {
		o.startStartTime = 0
		o.endStartTime = 4
		o.startSrcIP = 24
		o.endSrcIP = 40
		o.startDstIP = 40
		o.endDstIP = 56
		o.startSrcPort = 8
		o.endSrcPort = 10
		o.startDstPort = 10
		o.endDstPort = 12
		o.startProto = 4
		o.startPackets = 16
		o.endPackets = 20
		o.startBytes = 20
		o.endBytes = 24
		o.startDuration = 12
		o.endDuration = 16
		o.startTCPFlags = 5
		o.startClassType = 0
		o.startSensor = 0
		o.endSensor = 0
		o.startInitalFlags = 0
		o.startSessionFlags = 0
		o.startAttributes = 0
		o.startApplication = 6
		o.endApplication = 8
		o.startSNMPIn = 0
		o.endSNMPIn = 0
		o.startSNMPOut = 0
		o.endSNMPOut = 0
		o.startNextHopIP = 0
		o.endNextHopIP = 0
	} else if recordSize == 68 {
		o.startStartTime = 0
		o.endStartTime = 8
		o.startSrcIP = 36
		o.endSrcIP = 52
		o.startDstIP = 52
		o.endDstIP = 68
		o.startSrcPort = 12
		o.endSrcPort = 14
		o.startDstPort = 14
		o.endDstPort = 16
		o.startProto = 16
		o.startPackets = 28
		o.endPackets = 32
		o.startBytes = 32
		o.endBytes = 36
		o.startDuration = 8
		o.endDuration = 12
		o.startTCPFlags = 20
		o.startClassType = 17
		o.startSensor = 18
		o.endSensor = 20
		o.startInitalFlags = 21
		o.startSessionFlags = 22
		o.startAttributes = 23
		o.startApplication = 24
		o.endApplication = 26
		o.startSNMPIn = 0
		o.endSNMPIn = 0
		o.startSNMPOut = 0
		o.endSNMPOut = 0
		o.startNextHopIP = 0
		o.endNextHopIP = 0
	} else {
		err = fmt.Errorf("Unsupported record size:%d", recordSize)
	}
	return
}

//File contains header and silk slice of flow records
type File struct {
	Header Header
	Flows  []Flow
}

//OpenFile opens and parses silk file returning silk File struct and Error
func OpenFile(filePath string) (sf File, err error) {
	var f *os.File
	var n int
	var x int
	var start int
	var end int
	var readMod float64
	var shortReadCount = 0
	var ret int64
	var maxShortReadCount = 5
	var recordsCount int
	var silkFlow Flow
	var decompressedBuffer []byte
	var compressedBlockHeader = make([]byte, 8)
	var compressedBlockSize uint32
	var decompressedBlockSize uint32
	var compressedBuffer []byte
	var o offsets
	var ro io.ReadCloser
	var isTCP uint32

	if f, err = os.Open(filePath); err != nil {
		return
	}

	if sf.Header, err = parseHeader(f); err != nil {
		return
	}

	if o, err = getOffsets(sf.Header.RecordSize); err != nil {
		return
	}

	if sf.Header.Compression == 0 {
		mod := math.Floor(float64(defaultReadSize) / float64(sf.Header.RecordSize))
		readSize := (int(mod) * int(sf.Header.RecordSize))
		decompressedBuffer = make([]byte, readSize)
	}

	//TODO ADD flock syscall
	// syscall.Flock(f.Fd(), 1)
	// defer flock unlock

	var blockCount int
	for {
		blockCount++
		// log.Printf("blockCount:%d", blockCount)
		switch sf.Header.Compression {
		case 0:
			if n, err = f.Read(decompressedBuffer); n == 0 && err == io.EOF {
				err = nil
				return
			} else if err != nil {
				err = fmt.Errorf("Read error:%s", err.Error())
				return
			} else if n < int(sf.Header.RecordSize) {
				err = fmt.Errorf("Read:%d smaller then record size:%d", n, sf.Header.RecordSize)
				return
			} else if n < len(decompressedBuffer) {
				shortReadCount++
				if shortReadCount > maxShortReadCount {
					err = fmt.Errorf("Read failed to read requested size:%d %d times in a row", len(decompressedBuffer), maxShortReadCount)
					return
				}
			} else {
				shortReadCount = 0
			}

			recordsCount = n / int(sf.Header.RecordSize)

			if readMod = math.Mod(float64(n), float64(sf.Header.RecordSize)); readMod != 0 {
				if ret, err = f.Seek(-int64(readMod), 1); err != nil {
					err = fmt.Errorf("Seek ret:%d error:%s", ret, err)
					return
				}

				recordsCount = n / int(sf.Header.RecordSize)
			}
		case 3, 2, 1:
			if _, err = f.Read(compressedBlockHeader); err == io.EOF {
				err = nil
				return
			} else if err != nil {
				return
			}
			compressedBlockSize = binary.BigEndian.Uint32(compressedBlockHeader[0:4])
			decompressedBlockSize = binary.BigEndian.Uint32(compressedBlockHeader[4:8])
			if int(compressedBlockSize) > cap(compressedBuffer) {
				compressedBuffer = make([]byte, compressedBlockSize)
			}

			if int(decompressedBlockSize) > cap(decompressedBuffer) {
				decompressedBuffer = make([]byte, decompressedBlockSize)
			}

			if n, err = f.Read(compressedBuffer[:compressedBlockSize]); n == 0 && err == io.EOF {
				err = nil
				return
			} else if err != nil {
				return
			}
			if sf.Header.Compression == 3 {
				if _, err = snappy.Decode(decompressedBuffer, compressedBuffer[:compressedBlockSize]); err != nil {
					return
				}
			} else if sf.Header.Compression == 2 {
				if decompressedBuffer, err = lzo.Decompress1X(bytes.NewReader(compressedBuffer[:compressedBlockSize]), 0, 0); err != nil {
					return
				}
			} else if sf.Header.Compression == 1 {
				if ro, err = zlib.NewReader(bytes.NewReader(compressedBuffer[:compressedBlockSize])); err != nil {
					return
				}
				if _, err = ro.Read(decompressedBuffer); err != nil {
					ro.Close()
					return
				}
				ro.Close()
			}

			recordsCount = int(decompressedBlockSize) / int(sf.Header.RecordSize)
		default:
			err = ErrUnsupportedCompression
			return
		}

		start = 0
		end = int(sf.Header.RecordSize)
		for x = 0; x < recordsCount; x++ {
			//Clear out struct values
			silkFlow.startTimeMS56 = 0
			silkFlow.StartTimeMS = 0
			silkFlow.Duration = 0
			silkFlow.SrcIP = nil
			silkFlow.DstIP = nil
			silkFlow.SrcPort = 0
			silkFlow.DstPort = 0
			silkFlow.Proto = 0
			silkFlow.Flags = 0
			silkFlow.Packets = 0
			silkFlow.Bytes = 0
			silkFlow.ClassType = 0
			silkFlow.Sensor = 0
			silkFlow.InitalFlags = 0
			silkFlow.SessionFlags = 0
			silkFlow.Attributes = 0
			silkFlow.Application = 0
			silkFlow.SNMPIn = 0
			silkFlow.SNMPOut = 0
			silkFlow.NextHopIP = nil

			if sf.Header.FileFlags == 0 {
				//little endian
				if sf.Header.RecordSize == 56 {
					silkFlow.startTimeMS56 = binary.LittleEndian.Uint32(decompressedBuffer[start:end][o.startStartTime:o.endStartTime])
					silkFlow.StartTimeMS = uint64((calMsec & silkFlow.startTimeMS56)) + sf.Header.fileDateMS
					isTCP = silkFlow.startTimeMS56 & isTCPAnd
					if isTCP != 0 {
						silkFlow.Proto = 6
						if (decompressedBuffer[start:end][5] & SilkTCPStateExpanded) != 0 {
							silkFlow.Flags = decompressedBuffer[start:end][3] | decompressedBuffer[start:end][4]
						} else {
							silkFlow.Flags = decompressedBuffer[start:end][4]
						}
					} else {
						silkFlow.Flags = 0
						silkFlow.Proto = decompressedBuffer[start:end][4]
					}
				} else {
					silkFlow.Flags = uint8(decompressedBuffer[start:end][o.startTCPFlags])
					silkFlow.Proto = uint8(decompressedBuffer[start:end][o.startProto])
					silkFlow.StartTimeMS = binary.LittleEndian.Uint64(decompressedBuffer[start:end][o.startStartTime:o.endStartTime])
				}

				silkFlow.SrcIP = net.ParseIP(net.IP(decompressedBuffer[start:end][o.startSrcIP:o.endSrcIP]).String())
				silkFlow.DstIP = net.ParseIP(net.IP(decompressedBuffer[start:end][o.startDstIP:o.endDstIP]).String())
				silkFlow.SrcPort = binary.LittleEndian.Uint16(decompressedBuffer[start:end][o.startSrcPort:o.endSrcPort])
				silkFlow.DstPort = binary.LittleEndian.Uint16(decompressedBuffer[start:end][o.startDstPort:o.endDstPort])
				silkFlow.Packets = binary.LittleEndian.Uint32(decompressedBuffer[start:end][o.startPackets:o.endPackets])
				silkFlow.Bytes = binary.LittleEndian.Uint32(decompressedBuffer[start:end][o.startBytes:o.endBytes])
				silkFlow.Duration = binary.LittleEndian.Uint32(decompressedBuffer[start:end][o.startDuration:o.endDuration])

				if sf.Header.RecordSize == 88 {
					silkFlow.SNMPIn = binary.LittleEndian.Uint16(decompressedBuffer[start:end][o.startSNMPIn:o.endSNMPIn])
					silkFlow.SNMPOut = binary.LittleEndian.Uint16(decompressedBuffer[start:end][o.startSNMPOut:o.endSNMPOut])
					silkFlow.NextHopIP = net.ParseIP(net.IP(decompressedBuffer[start:end][o.startNextHopIP:o.endNextHopIP]).String())
				} else {
					silkFlow.Application = binary.LittleEndian.Uint16(decompressedBuffer[start:end][o.startApplication:o.endApplication])
				}

				if sf.Header.RecordSize == 88 || sf.Header.RecordSize == 68 {
					silkFlow.ClassType = decompressedBuffer[o.startClassType]
					silkFlow.Sensor = binary.LittleEndian.Uint16(decompressedBuffer[start:end][o.startSensor:o.endSensor])
					silkFlow.InitalFlags = decompressedBuffer[o.startInitalFlags]
					silkFlow.SessionFlags = decompressedBuffer[o.startSessionFlags]
					silkFlow.Attributes = decompressedBuffer[o.startAttributes]
				} else if sf.Header.RecordSize == 56 {
					silkFlow.Sensor = uint16(sf.Header.fileSensor)
				}

			} else {
				//big endian)
				if sf.Header.RecordSize == 56 {
					silkFlow.startTimeMS56 = binary.BigEndian.Uint32(decompressedBuffer[start:end][o.startStartTime:o.endStartTime])
					silkFlow.StartTimeMS = uint64((calMsec & silkFlow.startTimeMS56)) + sf.Header.fileDateMS
					isTCP = silkFlow.startTimeMS56 & isTCPAnd
					if isTCP != 0 {
						silkFlow.Proto = 6
						if (decompressedBuffer[start:end][5] & SilkTCPStateExpanded) != 0 {
							silkFlow.Flags = decompressedBuffer[start:end][3] | decompressedBuffer[start:end][4]
						} else {
							silkFlow.Flags = decompressedBuffer[start:end][4]
						}
					} else {
						silkFlow.Flags = 0
						silkFlow.Proto = decompressedBuffer[start:end][4]
					}
				} else {
					silkFlow.Flags = uint8(decompressedBuffer[start:end][o.startTCPFlags])
					silkFlow.Proto = uint8(decompressedBuffer[start:end][o.startProto])
					silkFlow.StartTimeMS = binary.BigEndian.Uint64(decompressedBuffer[start:end][o.startStartTime:o.endStartTime])
				}
				silkFlow.SrcIP = net.ParseIP(net.IP(decompressedBuffer[start:end][o.startSrcIP:o.endSrcIP]).String())
				silkFlow.DstIP = net.ParseIP(net.IP(decompressedBuffer[start:end][o.startDstIP:o.endDstIP]).String())
				silkFlow.SrcPort = binary.BigEndian.Uint16(decompressedBuffer[start:end][o.startSrcPort:o.endSrcPort])
				silkFlow.DstPort = binary.BigEndian.Uint16(decompressedBuffer[start:end][o.startDstPort:o.endDstPort])
				silkFlow.Packets = binary.BigEndian.Uint32(decompressedBuffer[start:end][o.startPackets:o.endPackets])
				silkFlow.Bytes = binary.BigEndian.Uint32(decompressedBuffer[start:end][o.startBytes:o.endBytes])
				silkFlow.Duration = binary.BigEndian.Uint32(decompressedBuffer[start:end][o.startDuration:o.endDuration])

				if sf.Header.RecordSize == 88 {
					silkFlow.SNMPIn = binary.BigEndian.Uint16(decompressedBuffer[start:end][o.startSNMPIn:o.endSNMPIn])
					silkFlow.SNMPOut = binary.BigEndian.Uint16(decompressedBuffer[start:end][o.startSNMPOut:o.endSNMPOut])
					copy(silkFlow.NextHopIP, decompressedBuffer[start:end][o.startNextHopIP:o.endNextHopIP])
					silkFlow.NextHopIP = net.ParseIP(net.IP(decompressedBuffer[start:end][o.startNextHopIP:o.endNextHopIP]).String())
				} else {
					silkFlow.Application = binary.BigEndian.Uint16(decompressedBuffer[start:end][o.startApplication:o.endApplication])
				}

				if sf.Header.RecordSize == 88 || sf.Header.RecordSize == 68 {
					silkFlow.Sensor = binary.BigEndian.Uint16(decompressedBuffer[start:end][o.startSensor:o.endSensor])
					silkFlow.InitalFlags = decompressedBuffer[o.startInitalFlags]
					silkFlow.SessionFlags = decompressedBuffer[o.startSessionFlags]
					silkFlow.Attributes = decompressedBuffer[o.startAttributes]
					silkFlow.ClassType = decompressedBuffer[o.startClassType]
				} else if sf.Header.RecordSize == 56 {
					silkFlow.Sensor = uint16(sf.Header.fileSensor)
				}
			}

			sf.Flows = append(sf.Flows, silkFlow)
			start += int(sf.Header.RecordSize)
			end += int(sf.Header.RecordSize)
		}
	}

}
