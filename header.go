package silk

import (
	"encoding/binary"
	"fmt"
	"io"
)

//Header is documented here:
//	https://tools.netsa.cert.org/silk/faq.html#file-header
type Header struct {
	MagicNumber   []byte
	FileFlags     uint8
	RecordFormat  uint8
	FileVersion   uint8
	Compression   uint8
	SilkVersion   uint32
	RecordSize    uint16
	RecordVersion uint16
	VarLenHeaders []VarLenHeader
	HeaderLength  int

	fileDateMS uint64
	fileSensor uint32
}

//VarLenHeader is part of the silk header. They contain different things
//like the cli command used to create the file. For some file types the
// variable length header also contains the year/month/day/hour of the file.
type VarLenHeader struct {
	ID      uint32
	Length  uint32
	Content []byte
}

func parseHeader(f io.Reader) (h Header, err error) {
	var n int
	var headerBytes = make([]byte, 16)
	var id, varLengthHeaderLength uint32
	var counter int
	var varLenHeader VarLenHeader
	// var buf = make([]byte, 2^10)

	if n, err = f.Read(headerBytes); err != nil {
		return
	} else if n != 16 {
		err = fmt.Errorf("Failed to read first 16 bytes of header, only read:%d", n)
		return
	}
	counter += 16
	h.MagicNumber = headerBytes[0:4]
	h.FileFlags = headerBytes[4]
	h.RecordFormat = headerBytes[5]
	h.FileVersion = headerBytes[6]
	h.Compression = headerBytes[7]
	h.SilkVersion = binary.BigEndian.Uint32(headerBytes[8:12])
	h.RecordSize = binary.BigEndian.Uint16(headerBytes[12:14])
	h.RecordVersion = binary.BigEndian.Uint16(headerBytes[14:16])

	for {
		var b = make([]byte, 8)
		if n, err = f.Read(b); err != nil {
			return
		} else if n != 8 {
			err = fmt.Errorf("Failed to read 8 bytes of variable length header, only read:%d", n)
			return
		}
		id = binary.BigEndian.Uint32(b[0:4])
		varLengthHeaderLength = binary.BigEndian.Uint32(b[4:8])

		switch id {
		case 0:
			//use value of varLengthHeaderLength above
		case 1:
			varLengthHeaderLength = 24
		case 2:
			//use value of varLengthHeaderLength above
		case 3:
			//use value of varLengthHeaderLength above
		case 4:
			//use value of varLengthHeaderLength above
		case 5:
			//use value of varLengthHeaderLength above
		case 6:
			varLengthHeaderLength = 16
		case 7:
			varLengthHeaderLength = 32
		default:
			err = fmt.Errorf("Unsupported variable length header id:%d", id)
			return
		}

		var varHeaderContent []byte

		if varLengthHeaderLength > 0 {
			varHeaderContent = make([]byte, varLengthHeaderLength-8)
			if _, err = f.Read(varHeaderContent); err != nil {
				return
			}
		}
		counter += int(varLengthHeaderLength)

		varLenHeader.ID = id
		varLenHeader.Length = varLengthHeaderLength
		varLenHeader.Content = varHeaderContent

		if id == 1 {
			h.fileDateMS = binary.BigEndian.Uint64(varHeaderContent[0:8])
			// h.fileSensor = binary.BigEndian.Uint32(varHeaderContent[8:12]) //Correct value but unknown purpose
			h.fileSensor = binary.BigEndian.Uint32(varHeaderContent[12:16]) //Correct value but unknown purpose
		}

		h.VarLenHeaders = append(h.VarLenHeaders, varLenHeader)
		if id == 0 {
			break
		}
	}
	h.HeaderLength = counter
	if (counter % int(h.RecordSize)) == 0 {
		return
	}

	var headerPadding int
	for i := 1; i <= 88; i++ {
		headerPadding = int(h.RecordSize) * i
		if headerPadding >= counter {
			break
		} else if i == 88 {
			err = fmt.Errorf("Failed to find end of silk header")
			return
		}
	}

	h.HeaderLength = counter
	var headerPaddingLength = int64(headerPadding - counter)
	if headerPaddingLength != 0 {
		var readPaddingBytes = make([]byte, headerPaddingLength)
		// fmt.Printf("headerPaddingLength:%d readPaddingBytes:%d\n", headerPaddingLength, len(readPaddingBytes))
		if _, err = f.Read(readPaddingBytes); err != nil {
			return
		}
		// if _, err = f.Seek(headerPaddingLength, 1); err != nil {
		// 	return
		// }
	}

	return
}
