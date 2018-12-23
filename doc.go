/*
Package silk is written without cgo to read common silk file formats.

What is silk?

https://tools.netsa.cert.org/silk/faq.html#what-silk

	"SiLK is a suite of network traffic collection and analysis tools developed and
	maintained by the CERT Network Situational Awareness Team (CERT NetSA) at
	Carnegie Mellon University to facilitate security analysis of large networks.
	The SiLK tool suite supports the efficient collection, storage, and analysis
	of network flow data, enabling network security analysts to rapidly query
	large historical traffic data sets."


Example:

	import (
		"fmt"
		"log"
		"silk"
	)

	func main() {

		var testFile = "testdata/FT_RWIPV6-v2-c0-L.dat"
		var err error
		var sf silk.File

		if sf, err = silk.OpenFile(testFile); err != nil {
			log.Fatalf("OpenFile() error:%s", err)
		}

		log.Printf("Compression:%d", sf.Header.Compression)
		log.Printf("FileFlags:%d", sf.Header.FileFlags)
		log.Printf("FileVersion:%d", sf.Header.FileVersion)
		log.Printf("HeaderLength:%d", sf.Header.HeaderLength)
		log.Printf("MagicNumber:%x", sf.Header.MagicNumber)
		log.Printf("RecordFormat:%d", sf.Header.RecordFormat)
		log.Printf("RecordSize:%d", sf.Header.RecordSize)
		log.Printf("RecordVersion:%d", sf.Header.RecordVersion)
		log.Printf("SilkVersion:%d", sf.Header.SilkVersion)

		log.Printf("File record count:%d\n", len(sf.Flows))

		fmt.Printf("start_time_ms,src_ip,dst_ip,src_port,dst_port\n")
		for _, flow := range sf.Flows {
			fmt.Printf("%d,%s,%s,%d,%d\n",
				flow.StartTimeMS,
				flow.SrcIP.String(),
				flow.DstIP.String(),
				flow.SrcPort,
				flow.DstPort,
			)
			//Etc... for other silk.Flow values
		}
	}

*/
package silk
